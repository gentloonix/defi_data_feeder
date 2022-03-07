package main

import (
	"context"
	"log"
	"runtime"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/sys/cpu"
)

// ### WEB3

const (
	CHAIN_ID = 0xa86a

	RPC_HTTPS = "https://api.avax.network/ext/bc/C/rpc"
	RPC_WSS   = "wss://api.avax.network/ext/bc/C/ws"
)

// ### Block Height

type BlockHeightWrapped struct {
	N uint64
	_ cpu.CacheLinePad
}

var (
	BlockHeight BlockHeightWrapped
)

func blockHeightDaemon() {
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					log.Println(err)
					return
				}
			}()

			client, _ := ethclient.Dial(RPC_WSS)
			if client == nil {
				log.Println("client is nil")
				return
			}
			defer client.Close()

			headers := make(chan *types.Header, 128)
			defer close(headers)

			sub, err := client.SubscribeNewHead(context.Background(), headers)
			if err != nil {
				log.Println(err)
				return
			}
			defer sub.Unsubscribe()

			for {
				select {
				case err := <-sub.Err():
					log.Println(err)
					return
				case header := <-headers:
					atomic.StoreUint64(&BlockHeight.N, header.Number.Uint64())
					// TODO Debug
					log.Println(BlockHeight.N)
				}
			}
		}()

		time.Sleep(15 * time.Second)
	}
}

// ### MAIN

func main() {
	runtime.LockOSThread()

	blockHeightDaemon()
}
