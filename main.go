package main

import (
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	abigen "github.com/ojo-network/uniswap-indexer/abi"
)

const (
	nodeUrl = "wss://ethereum.publicnode.com"
	//uniswapAddress = "0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45"
	// usdcAddress = "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
	// wethAddress = "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"

	poolAddress = "0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640"
)

var (
// usdcToken = common.HexToAddress(usdcAddress)
// wethToken = common.HexToAddress(wethAddress)
)

type SwapLogData struct {
	Amount0 big.Int
	Amount1 big.Int
}

func main() {

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
			// data, err := abi.ABI.Unpack(parsedABI, "Swap", log.Data)
			// if err != nil {
			// 	fmt.Println(err)
			// }
			fmt.Printf("%+v\n", event)
		}
	}
}
