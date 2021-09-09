package blockrand_test

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"testing"

	"crypto/rand"

	blockrand "github.com/tyeolrik/Go-BlockRand"
)

func TestHello(t *testing.T) {
	var blockSize uint16 = 6
	useridIndex := 0

	test_6block := blockrand.NewSR(blockSize)

	var isAllMined bool = false

	// Input Real Random Golang CryptoRand = /dev/urandom
	for block := 0; block < 6; block++ {
		for i := 0; i < 3; i++ {
			tempRandom, err := rand.Int(rand.Reader, big.NewInt(int64(^uint64(0)>>1)))
			if err != nil {
				log.Fatalln(err)
			}
			test_6block.Participate(strconv.Itoa(useridIndex), tempRandom.Uint64())
			useridIndex++
		}
		isAllMined = test_6block.Mining()
	}
	for isAllMined == false {
		isAllMined = test_6block.Mining()
	}
	fmt.Println(test_6block.GetFirst3Returns())
}

func TestMax(t *testing.T) {
	str := fmt.Sprintf("%b", ^uint64(0))
	fmt.Println(str)
	fmt.Println(len(str))
}
