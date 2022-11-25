package main

import (
	"fmt"
	"github.com/tendermint/tendermint/enclave/packages/gate/api"
)

// This is just a demo to ensure we can compile a static go binary
func main() {
	_, err := api.GetRandom()
	if err != nil {
		return
	}

	fmt.Println("finished")
}
