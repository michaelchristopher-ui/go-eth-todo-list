package common

import (
	"context"
	"fmt"
	"go-eth-todo-list/internal/pkg/abi/todolist"
	"go-eth-todo-list/internal/pkg/constant"
	"go-eth-todo-list/internal/pkg/util"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	Conn   *todolist.Todolist
	Auth   *bind.TransactOpts
	Client *ethclient.Client
)

func getAccountAuth(client *ethclient.Client, accountAddress string) *bind.TransactOpts {
	//fetch the last use nonce of account
	nonce, privateKey := util.GetNonceAndKeys(client, accountAddress)
	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units
	auth.GasPrice = big.NewInt(20000000000)
	return auth
}

func init() {
	var err error
	//Create Client
	Client, err = ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}
	//Create Auth
	Auth = getAccountAuth(Client, constant.GetWalletKey())
	//Set up Smart Contract Deploymen
	//Create Connection
	address, tx, conn, err := todolist.DeployTodolist(Auth, Client)
	if err != nil {
		panic(err)
	}
	Conn = conn
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P
}
