package main

import (
	"context"
	"log"
	"main/PairReader"
	"math/big"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

//lint:ignore U1000 noCopy
type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) UnLock() {}

// --- --- ---
const (
	RPC_HTTPS = "https://api.avax.network/ext/bc/C/rpc"
	RPC_WSS   = "wss://api.avax.network/ext/bc/C/ws"
)

// --- --- ---
var (
	BlockHeight uint64

	exp  int64
	expC chan struct{}
)

func sync() {
	exp = time.Now().Unix() + 15
}

func blockHeightDaemon() {
	runtime.LockOSThread()
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Println("blockHeightDaemon:", err)
					sync()
				}
			}()

			client, err := ethclient.Dial(RPC_WSS)
			if err != nil {
				panic(err)
			}
			defer client.Close()

			headers := make(chan *types.Header, 8)
			defer close(headers)

			sub, err := client.SubscribeNewHead(context.Background(), headers)
			if err != nil {
				panic(err)
			}
			defer func() {
				go sub.Unsubscribe()
			}()

			for {
				select {
				case <-expC:
					panic("header timeout")
				case err := <-sub.Err():
					panic(err)
				case header := <-headers:
					BlockHeight = header.Number.Uint64()
					sync()
				}
			}
		}()
		log.Println("blockHeightDaemon:", "cooldown")
		time.Sleep(3 * time.Second)
	}
}

func blockHeightWatchdog() {
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Println("blockHeightWatchdog:", err)
				}
			}()

			for range time.Tick(time.Second) {
				if time.Now().Unix() > exp {
					expC <- struct{}{}
				}
			}
		}()
	}
}

// --- --- ---
type Pair struct {
	//lint:ignore U1000 noCopy
	noCopy noCopy

	ID     int            `json:"id"`
	Pair   common.Address `json:"pair"`
	Token0 common.Address `json:"token0"`
	Token1 common.Address `json:"token1"`
	C      chan struct{}
}

func (p *Pair) Daemon() {
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Println("pair daemon:", p.ID, err)
				}
			}()

			client, err := ethclient.Dial(RPC_HTTPS)
			if err != nil {
				panic(err)
			}
			defer client.Close()

			instance, err := PairReader.NewPairReaderCaller(p.Pair, client)
			if err != nil {
				panic(err)
			}

			c := p.C
			blockHeight := uint64(0)
			reserve0 := big.NewInt(0)
			reserve1 := big.NewInt(0)
			for {
				<-c
				if blockHeight == BlockHeight {
					log.Println("pair daemon:", p.ID, "fast forward")
					continue
				}
				blockHeight = BlockHeight
				reserves, err := instance.GetReserves(nil)
				if err != nil {
					panic(err)
				}
				if reserve0.CmpAbs(reserves.Reserve0) != 0 ||
					reserve1.CmpAbs(reserves.Reserve1) != 0 {
					reserve0 = reserves.Reserve0
					reserve1 = reserves.Reserve1
				}
			}
		}()
	}
}

// --- --- ---

// --- --- ---
func main() {
	sync()
	expC = make(chan struct{})
	go blockHeightDaemon()
	go blockHeightWatchdog()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
