package block

import (
	DataTrans "../utils"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
 )

const targetBits = 24
const maxNonce = math.MaxInt64

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-target.BitLen()))
	pow := &ProofOfWork{b, target}
	return pow
}

func (proofOfWork *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			proofOfWork.block.PreBlockHash,
			proofOfWork.block.Data,
			DataTrans.Int2HexBytes(proofOfWork.block.Time),
			DataTrans.Int2HexBytes(int64(targetBits)),
			DataTrans.Int2HexBytes(int64(nonce)),
		}, []byte{})

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	fmt.Println("Mining the block contaning \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("\r %x", hash)
			break
		} else {
			nonce++
		}
	}
	fmt.Println("\n")
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data :=pow.prepareData(pow.block.Nonce)
	hash :=sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValidate :=hashInt.Cmp(pow.target)==-1
	return isValidate
}