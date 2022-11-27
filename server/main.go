package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	fmt.Println(verifySig(
		"0xda79856828a136c81ca7a5e17a09fa4ba78752c2",
		"0x1fdd5bf5fda9184176d31f0de4b01e25365659b43390ab3d8a092c0bfe66098611925ea5c225d7bca4aac93c511f5104d934dbc5a9d9d726207aec553727b7901c",
		[]byte("Example `personal_sign` message"),
	))
}

func verifySig(from, sigHex string, msg []byte) bool {
	sig := hexutil.MustDecode(sigHex)

	msg = accounts.TextHash(msg)
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	fmt.Println("Finded Address:", recoveredAddr)
	fmt.Println("Real Address", from)
	fmt.Println("Msg Hex:", hexutil.Encode(msg))
	return from == recoveredAddr.Hex()
}
