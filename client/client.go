package client

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ojo-network/ethereum-api/abi"
)

const (
	nodeUrl     = "wss://ethereum.publicnode.com"
	poolAddress = "0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640"
)

func Test() {

	ethClient, err := ethclient.Dial(nodeUrl)
	if err != nil {
		log.Fatal(err)
	}

	poolFilterer, err := abi.NewPoolFilterer(common.HexToAddress(poolAddress), ethClient)
	if err != nil {
		log.Fatal(err)
	}

	eventSink := make(chan *abi.PoolSwap)

	opts := &bind.WatchOpts{Start: nil, Context: ctx}
	subscription, err := poolFilterer.WatchSwap(opts, eventSink, nil, nil)

	if err != nil {
		log.Fatal(err)
	}

	subscriptionErr := subscription.Err()

	subscription.Unsubscribe()

}
