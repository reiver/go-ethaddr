package ethaddr

import (
	"encoding"
	"fmt"
	"math/big"

	"github.com/reiver/go-eip55"
	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-opt"
)

// AddressLength is the length of the eth-address measured in number of bytes.
const AddressLength = 20

// Address holds the eth-address.
//
// Note that type Address is an optional-type.
//
// (The 'optional-type' is also known as an 'option-type' and 'maybe-type' in other programming languages and libraries.)
type Address struct {
	optional opt.Optional[[AddressLength]byte]
}

var _ encoding.BinaryMarshaler = Address{}
var _ encoding.BinaryUnmarshaler = &Address{}
var _ encoding.TextMarshaler = Address{}
var _ encoding.TextUnmarshaler = &Address{}

// Nothing returns an empty address.
//
// For example:
//
//	var address ethaddr.Address = ethaddr.Nothing()
//
// Nothing is is NOT the same thing as 0x0000000000000000000000000000000000000000
func Nothing() Address {
	return Address{}
}

// Something returns an address whose value is equal to the contents of the 'value' variable.
//
// For example:
//
//	// 0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed
//	var address ethaddr.Address = ethaddr.Something( [20]byte{0x5a,0xAe,0xb6,0x05,0x3F,0x3E,0x94,0xC9,0xb9,0xA0,0x9f,0x33,0x66,0x94,0x35,0xE7,0xEf,0x1B,0xeA,0xed} )
//
// And also, for example:
//
//	//0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359
//	var address ethaddr.Address = ethaddr.Something( [20]byte{0xfB,0x69,0x16,0x09,0x5c,0xa1,0xdf,0x60,0xbB,0x79,0xCe,0x92,0xcE,0x3E,0xa7,0x4c,0x37,0xc5,0xd3,0x59} )
//
// And again, for example:
//
//	// 0x0000000000000000000000000000000000000000
//	var address ethaddr.Address = ethaddr.Something( [20]byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00} )
func Something(value [AddressLength]byte) Address {
	return Address{
		optional: opt.Something(value),
	}
}

// Parse returns the eth-address represented by the hexadecimal-literal.
func Parse(text []byte) (Address, error) {
	var address Address

	err := address.UnmarshalText(text)
	if nil != err {
		return Nothing(), err
	}

	return address, nil
}

// ParseString returns the eth-address represented by the hexadecimal-literal.
func ParseString(text string) (Address, error) {
	return Parse([]byte(text))
}

// ParseElsePanic is similar to the Parse func, except that it panic()s if there is an error.
func ParseElsePanic(text []byte) Address {
	address, err := Parse(text)
	if nil != err {
		panic(err)
	}

	return address
}

// ParseStringElsePanic is similar to the ParseString func, except that it panic()s if there is an error.
func ParseStringElsePanic(text string) Address {
	return ParseElsePanic([]byte(text))
}

// BigInt returns the eth-address represented by the *big.Int.
func BigInt(bigint *big.Int) (Address, error) {
	if nil == bigint {
		return Nothing(), errNilBigInt
	}

	if bigint.Cmp(minAddress) < 0 {
		return Nothing(), erorr.Errorf("ethaddr: address-underflow — expected numerical value for address to be between %s and %s but actually was %s", minAddress, maxAddress, bigint)
	}

	if bigint.Cmp(maxAddress) > 0 {
		return Nothing(), erorr.Errorf("ethaddr: address-overflow — expected numerical value for address to be between %s and %s but actually was %s", minAddress, maxAddress, bigint)
	}

	var address [AddressLength]byte
	bigint.FillBytes(address[:])

	return Something(address), nil
}

// BigIntElsePanic is similar to the BigInt func, except that it panic()s if there is an error.
func BigIntElsePanic(bigint *big.Int) Address {

	address, err := BigInt(bigint)
	if nil != err {
		panic(err)
	}

	return address

}

// Bytes returns the (decoded) bytes of the eth-address.
//
// For example, if the hexadecimal-literal of an eth-addres is:
//
//	"0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed"
//
// Then this Bytes method would return:
//
//	[]byte{0x5a, 0xae, 0xb6, 0x05, 0x3f, 0x3e, 0x94, 0xc9, 0xb9, 0xa0, 0x9f, 0x33, 0x66, 0x94, 0x35, 0xe7, 0xef, 0x1b, 0xea, 0xed}
//
// It would NOT return []byte("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed")
//
// I.e., it returns the "binary" representation (NOT the ASCII or UTF-8 "text" representation).
func (receiver Address) Bytes() []byte {
	value, something := receiver.optional.Get()
	if !something {
		return nil
	}

	return value[:]
}

// If the receiver contains nothing, then BigInt returns nil.
// If the receiver contains something, then BitInt returns that something as a *big.Int.
//
// For example:
//
//	// 0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed
//	var address ethaddr.Something( [20]byte{0x5a, 0xae, 0xb6, 0x05, 0x3f, 0x3e, 0x94, 0xc9, 0xb9, 0xa0, 0x9f, 0x33, 0x66, 0x94, 0x35, 0xe7, 0xef, 0x1b, 0xea, 0xed )
//	
//	// 517705355260207604495801938720638392742277016301517705355260207604495801938720638392742277016301
//	bigint := address.BigInt()
func (receiver Address) BigInt() *big.Int {
	value, something := receiver.optional.Get()
	if !something {
		return nil
	}

	return new(big.Int).SetBytes(value[:])
}

// EIP55 returns the EIP-55 / ERC-55 encoded hexadecimal-literal of the eth-address.
//
// EIP-55 / ERC-55 is an error-detection encoding of a ("0x" prefixed) hexadecimal-literal.
// A EIP-55 / ERC-55 encoded hexadecimal-literal makes it so you can detect certain types of errors.
//
// The EIP-55 / ERC-55 encoding potentially changes the letter-case (i.e., lower-case versus upper-case) of any "a", "b", "c", "d", "e", and "f" character in the hexadecimal-literal.
//
// For example, the EIP-55 / ERC-55 encoding for this hexadecimal-literal eth-address —
//
//	0x5aaeb6053f3e94c9b9a09f33669435e7ef1beaed
//
// — would be —
//
//	0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed
//	    ^      ^ ^  ^   ^           ^ ^  ^ ^
func (receiver Address) EIP55() string {
	value, something := receiver.optional.Get()
	if !something {
		return ""
	}

	return eip55.Encode(value)
}

// If Address contains nothing, then method Get returns false.
// If Address contains something, then method Get return true and then [20]byte that represents the address.
//
// For example:
//
//	var address ethaddr.Address = ethaddr.Nothing()
//	
//	// something == false
//	array, something := address.Get()
//
// And, for example
//
//	// 0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed
//	var address ethaddr.Address = ethaddr.Something( [20]byte{0x5a,0xAe,0xb6,0x05,0x3F,0x3E,0x94,0xC9,0xb9,0xA0,0x9f,0x33,0x66,0x94,0x35,0xE7,0xEf,0x1B,0xeA,0xed} )
//	
//	// something == true
//	// array == [20]byte{0x5a,0xAe,0xb6,0x05,0x3F,0x3E,0x94,0xC9,0xb9,0xA0,0x9f,0x33,0x66,0x94,0x35,0xE7,0xEf,0x1B,0xeA,0xed}
//	array, something := address.Get()
func (receiver Address) Get() ([AddressLength]byte, bool) {
	return receiver.optional.Get()
}

// If Address contains nothing, then method GetElse returns the value of the 'alternative' parameter.
// If Address contains something, then method GetElse returns the [20]byte that represents the address.
//
// For example:
//
//	var alternative [20]byte{0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE}
//	
//	var address ethaddr.Address = ethaddr.Nothing()
//	
//	// array == alternative == [20]byte{0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE}
//	array := address.GetElse()
//
// And, for example
//
//	var alternative [20]byte{0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE,0xEE}
//	
//	// 0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed
//	var address ethaddr.Address = ethaddr.Something( [20]byte{0x5a,0xAe,0xb6,0x05,0x3F,0x3E,0x94,0xC9,0xb9,0xA0,0x9f,0x33,0x66,0x94,0x35,0xE7,0xEf,0x1B,0xeA,0xed} )
//	
//	// array == [20]byte{0x5a,0xAe,0xb6,0x05,0x3F,0x3E,0x94,0xC9,0xb9,0xA0,0x9f,0x33,0x66,0x94,0x35,0xE7,0xEf,0x1B,0xeA,0xed}
//	array := address.GetElse()
func (receiver Address) GetElse(alternative [AddressLength]byte) [AddressLength]byte {
	return receiver.optional.GetElse(alternative)
}

// GoString returns the value of the receiver in Go syntax.
//
// This function is called when the "%#v" verb is used with the printing-functions from the Go built-in "fmt" package.
//
// For example:
//
//	var address ethaddr = ethaddr.Something([20]byte{0x5a,0xAe,0xb6,0x05,0x3F,0x3E,0x94,0xC9,0xb9,0xA0,0x9f,0x33,0x66,0x94,0x35,0xE7,0xEf,0x1B,0xeA,0xed})
//	
//	fmt.Printf("address = %#v \n", address)
func (receiver Address) GoString() string {
	value, something := receiver.optional.Get()
	if !something {
		return "ethaddr.Nothing()"
	}

	return fmt.Sprintf("ethaddr.Something(%#v)", value)
}

// IsNothing return true if the receiver address contains nothing.
// IsNothing return false if the receiver address contains something.
func (receiver Address) IsNothing() bool {
	return receiver.optional.IsNothing()
}

// IsSomething return true if the receiver address contains something.
// IsSomething return false if the receiver address contains nothing.
func (receiver Address) IsSomething() bool {
	return receiver.optional.IsSomething()
}

// MarshalBinary returns the eth-address in its binary form as a []byte.
func (receiver Address) MarshalBinary() ([]byte, error) {
	value, something := receiver.optional.Get()
	if !something {
		return nil, errNothing
	}

	return value[:], nil
}

// MarshalText returns the eth-address in its textual form as a []byte.
//
// (Note that this is different than the "binary" form of the eth-address as a []byte.)
//
// More specifically, the MarshalText method return the EIP-55 / ERC-55 encoding of the hexadecimal-literal representation of an eth-address.
//
// For example:
//
//	[]byte("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed")
//
// Or:
//
//	[]byte("0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359")
//
// Or:
//
//	[]byte("0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB")
//
// Or:
//
//	[]byte("0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb")
func (receiver Address) MarshalText() ([]byte, error) {
	if receiver.IsNothing() {
		return nil, errNothing
	}

	return []byte(receiver.EIP55()), nil
}

// MarshalText returns the eth-address in its textual form as a string.
//
// More specifically, the MarshalText method return the EIP-55 / ERC-55 encoding of the hexadecimal-literal representation of an eth-address.
//
// For example:
//
//	[]byte("0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed")
//
// Or:
//
//	[]byte("0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359")
//
// Or:
//
//	[]byte("0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB")
//
// Or:
//
//	[]byte("0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb")
//
func (receiver Address) String() string {
	return receiver.EIP55()
}

// UnmarshalBinary sets the receiver to the eth-address in its binary form as a []byte.
func (receiver *Address) UnmarshalBinary(data []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	{
		const expected int = AddressLength
		var   actual   int = len(data)

		if expected != actual {
			return erorr.Errorf("ethaddr: the actual length of the data parameter (%d) is not what was expected (%d)", actual, expected)
		}
	}

	var address [AddressLength]byte
	{
		copy(address[:], data)
	}

	receiver.optional = opt.Something(address)
	return nil
}

// UnmarshalText sets the receiver the eth-address represented by the hexadecimal-literal.
func (receiver *Address) UnmarshalText(text []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	var address [AddressLength]byte

	err := unmarshalText(&address, text)
	if nil != err {
		return err
	}

	receiver.optional = opt.Something(address)
	return nil
}
