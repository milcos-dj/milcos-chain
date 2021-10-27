package block

import (
	"time"
)

var (
	blocks   = []*BlockChain{genesisBlock}
)

var genesisBlock = &BlockChain{
	Index:        0,
	PrevBlockChainHash: []byte(""),
	Timestamp:    1465154705,
	Data:         []byte(""),
	Hash:         []byte("0000001ecea26a0894fbd46de4a2217a18e1c7ab965ca6b8b2b57cb62cbceeec"),
	Nonce: 15542467,
}

type BlockChain struct {
	Index int
	Timestamp int64
	Data []byte
	PrevBlockChainHash []byte
	Hash []byte
	Nonce int
}

func NewBlockChain(data []byte) *BlockChain {
	prevBlockChain := GetLastBlock()
	b := &BlockChain{prevBlockChain.Index+1,time.Now().Unix(), data,prevBlockChain.Hash, []byte{}, 0}
	pow := NewProofOfWork(b)
	nonce, hash := pow.CalculateHashForBlock()
	b.Hash = hash
	b.Nonce = nonce
	return b
}

func GetLastBlock() *BlockChain {
	return blocks[len(blocks) - 1]
}



