package main

import (
	"context"
	"log"
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

type Pair struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	KLast    *big.Int
	Extras   *PairExtras
}

type PairExtras struct {
	Token0 common.Address
	Token1 common.Address
}

const (
	CHAIN_ID  = 0xa86a
	RPC_HTTPS = "https://api.avax.network/ext/bc/C/rpc"
	RPC_WSS   = "wss://api.avax.network/ext/bc/C/ws"
)

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

			headers := make(chan *types.Header, 8)
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
					log.Println(BlockHeight)
				}
			}
		}()

		time.Sleep(10 * time.Second)
	}
}

func main() {
	go blockHeightDaemon()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
