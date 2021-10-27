package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"milcos-chain/common"
	"milcos-chain/core/utils"
)

type ProofOfWork struct {
	Block *BlockChain
	Target *big.Int
}

func NewProofOfWork(b *BlockChain) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - common.TargetBits))
	return &ProofOfWork{b, target}
}

func (pow *ProofOfWork) CalculateHashForBlock() (int, []byte) {
	nonce := 0
	var hashInt big.Int
	var hash [32]byte

	for nonce <math.MaxInt64 {
		data := pow.prepareData(nonce)
		hash := sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.Target) == -1 {
			fmt.Printf("\r%x\n", hash)
			break
		} else {
			nonce++
		}

	}
	return  nonce, hash[:]
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	b := pow.Block
	data := bytes.Join([][]byte{
		utils.IntToHex(int64(b.Index)),
		b.PrevBlockChainHash,
		utils.IntToHex(b.Timestamp),
		b.Data,
		utils.IntToHex(int64(nonce)),
	},
	[]byte{})
	return data
}

func (pow *ProofOfWork) Validate() bool {
	data := pow.prepareData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	var hashInt big.Int
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.Target) == -1
}

