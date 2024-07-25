package ethaddr

import (
	"github.com/reiver/go-erorr"
)

const (
	errMissingHexadecimalLiteralPrefix = erorr.Error("ethaddr: missing prefix for hexadecimal-literal (i.e., \"0x\")")
	errNilBigInt                       = erorr.Error("ethaddr: nil big-int")
	errNilDestination                  = erorr.Error("ethaddr: nil destination")
	errNilReceiver                     = erorr.Error("ethaddr: nil receiver")
	errNothing                         = erorr.Error("ethaddr: nothing")
)
