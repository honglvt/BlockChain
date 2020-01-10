package block

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockChain struct {
	Tip []byte
	DB  *bolt.DB
}

func (bc *BlockChain) AddBlock(data string) {
	var lashHash []byte
	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lashHash = b.Get([]byte("l"))
		return nil
	})

	newBlock := NewBlock(data, lashHash)

	err = bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		err = b.Put([]byte("l"), newBlock.Hash)
		bc.Tip = newBlock.Hash
		log.Fatal(err)
		return nil
	})

	log.Fatal(err)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

const dbFile = ""
const blocksBucket = ""

func NewBlockChain() *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
			log.Panic(err)
		} else {
		}
		return nil
	})
	bc := &BlockChain{tip, db}
	log.Fatal(err)
	return bc
}
