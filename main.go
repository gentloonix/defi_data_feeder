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

// noCopy may be embedded into structs which must not be copied
// after the first use.
//
// See https://golang.org/issues/8005#issuecomment-190753527
// for details.
//lint:ignore U1000 noCopy
type noCopy struct{}

// Lock is a no-op used by -copylocks checker from `go vet`.
func (*noCopy) Lock()   {}
func (*noCopy) UnLock() {}

const (
	CHAIN_ID  = 0xa86a
	RPC_HTTPS = "https://api.avax.network/ext/bc/C/rpc"
	RPC_WSS   = "wss://api.avax.network/ext/bc/C/ws"
)

// TypeID
const (
	JOE = iota
	PANGOLIN
	UNISWAP_V2
)

type Pair struct {
	//lint:ignore U1000 noCopy
	noCopy noCopy

	Reserve0 *big.Int
	Reserve1 *big.Int
	KLast    *big.Int
	C        chan struct{}
	Static   *PairStatic
}
type PairStatic struct {
	//lint:ignore U1000 noCopy
	noCopy noCopy

	ID     int
	TypeID int
	Pair   common.Address
	Token0 common.Address
	Token1 common.Address
}

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
					log.Println(BlockHeight, time.Now().Unix(), header.Time)
				}
			}
		}()

		time.Sleep(time.Minute)
	}
}

func main() {
	go blockHeightDaemon()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
