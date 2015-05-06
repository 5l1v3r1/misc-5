package bitcoin

import (
	"encoding/hex"
	"testing"
)

func TestHash256(t *testing.T) {
	input, _ := hex.DecodeString("0100000081cd02ab7e569e8bcd9317e2fe99f2de44d49ab2b8851ba4a308000000000000" +
		"e320b6c2fffc8d750423db8b1eb942ae710e951ed797f7affc8892b0f1fc122bc7f5d74df2b9441a42a14695")
	expected := "1dbd981fe6985776b644b173a4d0385ddc1aa2a829688d1e0000000000000000"

	hash256 := Hash256(input)

	if expected != hex.EncodeToString(hash256[:]) {
		t.Errorf("Encoded hash is %02x instead of %02x", hash256, expected)
	}
}

func TestHash160(t *testing.T) {
	input, _ := hex.DecodeString("0279be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798")
	expected := "751e76e8199196d454941c45d1b3a323f1433bd6"
	hash160 := Hash160(input)

	if expected != hex.EncodeToString(hash160[:]) {
		t.Errorf("Encoded hash is %02x instead of %02x", hash160, expected)

	}
}
