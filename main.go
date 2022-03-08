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
	CHAIN_ID  = 0xa86a
	RPC_HTTPS = "https://api.avax.network/ext/bc/C/rpc"
	RPC_WSS   = "wss://api.avax.network/ext/bc/C/ws"
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
					log.Println(err)
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
			defer sub.Unsubscribe()

			for {
				select {
				case err := <-sub.Err():
					panic(err)
				case header := <-headers:
					BlockHeight = header.Number.Uint64()
				}
			}
		}()

		time.Sleep(time.Minute)
	}
}

// --- --- ---
type Pair struct {
	//lint:ignore U1000 noCopy
	noCopy noCopy

	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
	ID                 int
	Static             *PairStatic
}
type PairStatic struct {
	//lint:ignore U1000 noCopy
	noCopy noCopy

	C      chan struct{}
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

			instance, err := PairReader.NewPairReaderCaller(p.Static.Pair, client)
			if err != nil {
				panic(err)
			}

			c := p.Static.C
			for {
				<-c
				reserves, err := instance.GetReserves(nil)
				if err != nil {
					panic(err)
				}
				p.Reserve0 = reserves.Reserve0
				p.Reserve1 = reserves.Reserve1
				p.BlockTimestampLast = reserves.BlockTimestampLast
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
