package crypto

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

func TestGetCryptoer(t *testing.T) {
	c := getCryptoer()
	src := []byte("Hello")
	encodedStr := hex.EncodeToString(src)
	fmt.Println(src)
	fmt.Printf("%s\n", encodedStr)
	prv, pub, err := c.genKeyPair()
	if err != nil {
		return
	}
	prvStr := hex.EncodeToString(prv)
	pubStr := hex.EncodeToString(pub)
	fmt.Println("privateKey is:", prv, "publicKey", pub)
	fmt.Println("privateKeyString is:", prvStr, "publicKeyString is:", pubStr)
	addr := Address(pub)
	// 	log.Fatal(err)
	// }
	// pubString := "7c66ce7703e6e4c4e31ba36c6eee29de345a8e9d36611f6bd2c809d3d0d47788fe3a66ab1970a8ea7d8b1f46e67956a481d638a0ab92a9aaaf0fbd2151af702e"
	// pub, err := hex.DecodeString(pubString)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("signedDataByPriv is:", signedDataByte)
	ok, err := c.verify(pub, src, signedDataByte)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ok)
}
