package main

import (
	"context"
	"crypto/ecdsa"
	"doodles/mycontract"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// eth的每个用户的交易都有nonce值递增的，如果发了一笔新的交易的nonce值大于上一笔
// 且新交易的gas费用比上一笔大且新一笔交易的优先被打包，上一笔交易会被取消
func getAuth(value int64) *bind.TransactOpts {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/42557ae4b4b1433eab1cdeeab805a934")
	// 私钥
	privateKey, err := crypto.HexToECDSA("40254bd534ed7bc2bdacdf7a6ea3e74188987af30a28988a94926bccbf1e2364")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nonce)
	// estimate gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(value) // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	return auth
}

const wei = 1000000000000000000

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/42557ae4b4b1433eab1cdeeab805a934")
	if err != nil {
		fmt.Println("eth client dial error:", err)
		panic(err)
	}

	// 推荐gas费多少，可以乘以一个系数.
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}

	//contractClient, err := mycontract.NewMycontractCaller(common.HexToAddress("0x3245fFC283D150CF6eAe94e74c37b0DeAB46Aea5"), client)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//name, err := contractClient.Name(nil)
	//getAuth()
	//contractClient.

	contractTransactor, err := mycontract.NewMycontractTransactor(common.HexToAddress("0x3245fFC283D150CF6eAe94e74c37b0DeAB46Aea5"), client)
	if err != nil {
		log.Fatal(err)
	}
	// 设置合约状态
	//saleAction, err := contractTransactor.SetSaleState(getAuth(0), true)
	//if err != nil {
	//	log.Fatal(err)
	//}

	fmt.Println(0.123 * wei * 2)
	mintAction, err := contractTransactor.Mint(getAuth(0.1233*wei*2), big.NewInt(2))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mintAction)
}
