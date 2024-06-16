package ethaddr

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errMissingHexadecimalLiteralPrefix = erorr.Error("ethaddr: missing prefix for hexadecimal-literal (i.e., \"0x\")")
	errNilDestination                  = erorr.Error("ethaddr: nil destination")
	errNilReceiver                     = erorr.Error("ethaddr: nil receiver")
)
