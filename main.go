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
	RPC_HTTPS       = "https://api.avax.network/ext/bc/C/rpc"
	RPC_WSS         = "wss://api.avax.network/ext/bc/C/ws"
	TIMEOUT_SECONDS = 10
)

// --- --- ---
var (
	BlockHeight uint64
)

func blockHeightDaemon() {
	runtime.LockOSThread()
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Println("blockHeightDaemon:", err)
				}
			}()

			client, err := ethclient.Dial(RPC_WSS)
			if err != nil {
				panic(err)
			}
			defer client.Close()

			headers := make(chan *types.Header, 16)
			defer close(headers)

			sub, err := client.SubscribeNewHead(context.Background(), headers)
			if err != nil {
				panic(err)
			}
			defer func() {
				go sub.Unsubscribe()
			}()

			exp := time.Now().Unix() + TIMEOUT_SECONDS
			for {
				select {
				case err := <-sub.Err():
					panic(err)
				case header := <-headers:
					BlockHeight = header.Number.Uint64()
					exp = time.Now().Unix() + TIMEOUT_SECONDS
				default:
					if time.Now().Unix() > exp {
						panic("header timeout")
					}
				}
			}
		}()
		log.Println("blockHeightDaemon:", "cooldown")
		time.Sleep(10 * time.Second)
	}
}

// --- --- ---
type Pair struct {
	//lint:ignore U1000 noCopy
	noCopy noCopy

	C      chan struct{}
	ID     int
	Pair   common.Address
	Token0 common.Address
	Token1 common.Address
}

func (p *Pair) Daemon() {
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Println(err)
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
			reserve0 := big.NewInt(0)
			reserve1 := big.NewInt(0)
			for {
				<-c
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
func main() {
	go blockHeightDaemon()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
