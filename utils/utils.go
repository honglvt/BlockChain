package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

func main() {

}

func Int2HexBytes(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
