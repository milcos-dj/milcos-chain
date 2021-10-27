package block

import (
	"fmt"
	"testing"
)

func TestPow(t *testing.T)  {
	b := NewBlockChain([]byte("test block"))
	fmt.Println(string(b.PrevBlockChainHash))
	fmt.Println(string(b.Data))
	fmt.Println(b.Nonce)

}
func TestPowValidate(t *testing.T)  {
	pow := NewProofOfWork(GetLastBlock())
	validate := pow.Validate()
	fmt.Println(validate)

}
