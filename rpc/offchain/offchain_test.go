package offchain

import (
	"os"
	"testing"

	"github.com/ComposableFi/go-substrate-rpc-client/v4/client"
	"github.com/ComposableFi/go-substrate-rpc-client/v4/config"
)

var testOffchain Offchain

func TestMain(m *testing.M) {
	cl, err := client.Connect(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}
	testOffchain = NewOffchain(cl)
	os.Exit(m.Run())
}
