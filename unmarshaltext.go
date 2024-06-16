package ethaddr

import (
	"bytes"

	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-hexadeca"
)

// "0x"
var hexlitprefix [2]byte = [2]byte{'0', 'x'}

//unmarshalText unmarshals the ("0x" prefixed) hexadecimal-literal in 'text' into 'dst'.
//
// It, for example, turns this:
//
//	[]byte("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed")
//
// Into:
//
//	[40]byte{0x5a,0xAe,0xb6,0x05,0x3F,0x3E,0x94,0xC9,0xb9,0xA0,0x9f,0x33,0x66,0x94,0x35,0xE7,0xEf,0x1B,0xeA,0xed}
func unmarshalText(dst *[AddressLength]byte, text []byte) error {

	if nil == dst {
		return errNilDestination
	}

	var hex []byte
	{
		if !bytes.HasPrefix(text, hexlitprefix[:]) {
			return errMissingHexadecimalLiteralPrefix
		}

		hex = text[len(hexlitprefix):]
	}

	var value0 byte
	var rest []byte
	var numHandled int
	{
		switch len(hex) {
		case AddressLength*2:
			numHandled = 2

			var hex0 byte = hex[0]
			var hex1 byte = hex[1]

			var decoded0 byte
			{
				var ok bool

				decoded0, ok = hexadeca.DecodeByte(hex0)
				if !ok {
					return erorr.Errorf("ethaddr: byte number-0 (after \"0x\" prefix) of hexadecimal literal (%d) (%q) is not a valid hexadecimal symbol" , hex0, hex0)
				}
			}

			var decoded1 byte
			{
				var ok bool

				decoded1, ok = hexadeca.DecodeByte(hex1)
				if !ok {
					return erorr.Errorf("ethaddr: byte number-1 (after \"0x\" prefix) of hexadecimal literal (%d) (%q) is not a valid hexadecimal symbol" , hex1, hex1)
				}
			}


			value0 = (decoded0 << 4) | decoded1
			rest = hex[2:]

		case AddressLength*2 - 1:
			numHandled = 1

			var hex0 byte = hex[0]

			var decoded0 byte
			{
				var ok bool

				decoded0, ok = hexadeca.DecodeByte(hex0)
				if !ok {
					return erorr.Errorf("ethaddr: byte number-0 (after \"0x\" prefix) of hexadecimal literal (%d) (%q) is not a valid hexadecimal symbol" , hex0, hex0)
				}
			}

			value0 = (decoded0 << 4)
			rest = hex[1:]

		default:
			var expected1 int = AddressLength*2 + len(hexlitprefix)
			var expected2 int = expected1 - 1

			return erorr.Errorf("ethaddr: the eth-address is expected to be %d or %d bytes long, but was actually %d bytes long", expected1, expected2, len(text))
		}
	}

	dst[0] = value0

	{
		var limit int = len(rest) / 2

		for i:=0; i<limit; i++ {
			var indexToMostSignificant  int = i * 2
			var indexToLeastSignificant int = indexToMostSignificant + 1

			var mostSignificantHex  byte = rest[indexToMostSignificant]
			var leastSignificantHex byte = rest[indexToLeastSignificant]

			mostSignificant, ok := hexadeca.DecodeByte(mostSignificantHex)
			if !ok {
				return erorr.Errorf("ethaddr: byte number-%d (after \"0x\" prefix) of hexadecimal literal (%d) (%q) is not a valid hexadecimal symbol" , numHandled+indexToMostSignificant, mostSignificantHex, mostSignificantHex)
			}

			leastSignificant, ok := hexadeca.DecodeByte(leastSignificantHex)
			if !ok {
				return erorr.Errorf("ethaddr: byte number-%d (after \"0x\" prefix) of hexadecimal literal (%d) (%q) is not a valid hexadecimal symbol" , numHandled+indexToLeastSignificant, leastSignificantHex, leastSignificantHex)
			}

			var value byte = (mostSignificant << 4) | leastSignificant

			dst[1+i] = value
		}
	}

	return nil
}
