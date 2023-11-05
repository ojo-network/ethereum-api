package client

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	abigen "github.com/ojo-network/ethereum-api/abi"
)

func RunOldClient() {

	userInterrupt := make(chan os.Signal, 1)
	signal.Notify(userInterrupt, os.Interrupt, syscall.SIGTERM)

	// Connect to the Ethereum node
	ethClient, err := ethclient.Dial(nodeUrl)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(poolAddress)

	// Read the ABI from a JSON file
	abiFile, err := os.Open("abi/uniswap.json")
	if err != nil {
		log.Fatal(err)
	}
	abiReader := io.Reader(abiFile)
	parsedABI, err := abi.JSON(abiReader)

	uniswapContract := bind.NewBoundContract(address, parsedABI, ethClient, ethClient, ethClient)
	if err != nil {
		log.Fatal(err)
	}

	logs, subscription, err := uniswapContract.WatchLogs(nil, "Swap")
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-userInterrupt:
			log.Println("Exiting...")
			subscription.Unsubscribe()
			return
		case err := <-subscription.Err():
			log.Fatalf("error while listening for events: %v", err)
		case log := <-logs:
			event := new(abigen.PoolSwap)
			if err := uniswapContract.UnpackLog(event, "Swap", log); err != nil {
				fmt.Println(err)
			}
			// amount0 := event.Amount0 // use this number for volume
			// amount1 := event.Amount1

			// amount0Abs := math.Abs(float64(amount0.Int64()))

			fmt.Printf("%+v\n", event)
		}
	}
}
