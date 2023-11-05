package client

import (
	"context"
	"fmt"

	sdkmath "cosmossdk.io/math"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ojo-network/ethereum-api/abi"
	"github.com/ojo-network/indexer/indexer"
	"github.com/ojo-network/indexer/utils"
)

const (
	nodeUrl     = "wss://ethereum.publicnode.com"
	poolAddress = "0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640"
)

func WatchSwapEvent(i *indexer.Indexer, ctx context.Context) error {
	ethClient, err := ethclient.Dial(nodeUrl)
	if err != nil {
		return err
	}

	poolFilterer, err := abi.NewPoolFilterer(common.HexToAddress(poolAddress), ethClient)
	if err != nil {
		return err
	}

	eventSink := make(chan *abi.PoolSwap)

	opts := &bind.WatchOpts{Start: nil, Context: ctx}
	subscription, err := poolFilterer.WatchSwap(opts, eventSink, nil, nil)
	if err != nil {
		return err
	}

	srvErrCh := make(chan error, 1)

	go func() {
		select {
		case <-ctx.Done():
			subscription.Unsubscribe()
			srvErrCh <- nil
			break
		case err := <-subscription.Err():
			srvErrCh <- err
			break
		case event := <-eventSink:
			swap := indexer.Swap{
				BlockNum:     indexer.BlockNum(event.Raw.BlockNumber),
				Timestamp:    utils.CurrentUnixTime(),
				ExchangePair: "WETH/ETH",
				AmountIn:     sdkmath.LegacyNewDecFromBigInt(event.Amount0),
				AmountOut:    sdkmath.LegacyNewDecFromBigInt(event.Amount1),
			}
			i.AddSwap(swap)
			fmt.Println(event)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		case err = <-srvErrCh:
			return err
		}
	}
}
