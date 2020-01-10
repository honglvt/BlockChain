package block

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	//时间戳
	Time int64
	//区块存储的有效信息
	Data []byte
	//上一个区块的Hash值
	PreBlockHash []byte
	//当前区块的Hash值
	Hash  []byte
	Nonce int
}

func NewBlock(data string, preBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), preBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func DeSerializeBlock(b []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}
