package ethaddr

import (
	"math/big"
)

var (
	minAddress *big.Int = big.NewInt(0).SetBytes( []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} )
	maxAddress *big.Int = big.NewInt(0).SetBytes( []byte{0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF} )
)
