package ethaddr

import (
	"testing"
)

func TestHexDecode(t *testing.T) {

	tests := []struct{
		Byte byte
		ExpectedByte byte
		ExpectedBool bool
	}{
		{
			Byte: 0x00,
		},
		{
			Byte: 0x01,
		},
		{
			Byte: 0x02,
		},
		{
			Byte: 0x03,
		},
		{
			Byte: 0x04,
		},
		{
			Byte: 0x05,
		},
		{
			Byte: 0x06,
		},
		{
			Byte: 0x07,
		},
		{
			Byte: 0x08,
		},
		{
			Byte: 0x09,
		},
		{
			Byte: 0x0A,
		},
		{
			Byte: 0x0B,
		},
		{
			Byte: 0x0C,
		},
		{
			Byte: 0x0D,
		},
		{
			Byte: 0x0E,
		},
		{
			Byte: 0x0F,
		},
		{
			Byte: 0x10,
		},
		{
			Byte: 0x11,
		},
		{
			Byte: 0x12,
		},
		{
			Byte: 0x13,
		},
		{
			Byte: 0x14,
		},
		{
			Byte: 0x15,
		},
		{
			Byte: 0x16,
		},
		{
			Byte: 0x17,
		},
		{
			Byte: 0x18,
		},
		{
			Byte: 0x19,
		},
		{
			Byte: 0x1A,
		},
		{
			Byte: 0x1B,
		},
		{
			Byte: 0x1C,
		},
		{
			Byte: 0x1D,
		},
		{
			Byte: 0x1E,
		},
		{
			Byte: 0x1F,
		},
		{
			Byte: 0x20,
		},
		{
			Byte: 0x21,
		},
		{
			Byte: 0x22,
		},
		{
			Byte: 0x23,
		},
		{
			Byte: 0x24,
		},
		{
			Byte: 0x25,
		},
		{
			Byte: 0x26,
		},
		{
			Byte: 0x27,
		},
		{
			Byte: 0x28,
		},
		{
			Byte: 0x29,
		},
		{
			Byte: 0x2A,
		},
		{
			Byte: 0x2B,
		},
		{
			Byte: 0x2C,
		},
		{
			Byte: 0x2D,
		},
		{
			Byte: 0x2E,
		},
		{
			Byte: 0x2F,
		},
		{
			Byte: 0x30, // '0'
			ExpectedByte:   0,
			ExpectedBool: true,
		},
		{
			Byte: 0x31, // '1'
			ExpectedByte:   1,
			ExpectedBool: true,
		},
		{
			Byte: 0x32, // '2'
			ExpectedByte:   2,
			ExpectedBool: true,
		},
		{
			Byte: 0x33, // '3'
			ExpectedByte:   3,
			ExpectedBool: true,
		},
		{
			Byte: 0x34, // '4'
			ExpectedByte:   4,
			ExpectedBool: true,
		},
		{
			Byte: 0x35, // '5'
			ExpectedByte:   5,
			ExpectedBool: true,
		},
		{
			Byte: 0x36, // '6'
			ExpectedByte:   6,
			ExpectedBool: true,
		},
		{
			Byte: 0x37, // '7'
			ExpectedByte:   7,
			ExpectedBool: true,
		},
		{
			Byte: 0x38, // '8'
			ExpectedByte:   8,
			ExpectedBool: true,
		},
		{
			Byte: 0x39, // '9'
			ExpectedByte:   9,
			ExpectedBool: true,
		},
		{
			Byte: 0x3A,
		},
		{
			Byte: 0x3B,
		},
		{
			Byte: 0x3C,
		},
		{
			Byte: 0x3D,
		},
		{
			Byte: 0x3E,
		},
		{
			Byte: 0x3F,
		},
		{
			Byte: 0x40,
		},
		{
			Byte: 0x41, // 'A'
			ExpectedByte:  10,
			ExpectedBool: true,
		},
		{
			Byte: 0x42, // 'B'
			ExpectedByte:  11,
			ExpectedBool: true,
		},
		{
			Byte: 0x43, // 'C'
			ExpectedByte:  12,
			ExpectedBool: true,
		},
		{
			Byte: 0x44, // 'D'
			ExpectedByte:  13,
			ExpectedBool: true,
		},
		{
			Byte: 0x45, // 'E'
			ExpectedByte:  14,
			ExpectedBool: true,
		},
		{
			Byte: 0x46, // 'F'
			ExpectedByte:  15,
			ExpectedBool: true,
		},
		{
			Byte: 0x47, // 'G'
		},
		{
			Byte: 0x48, // 'H'
		},
		{
			Byte: 0x49, // 'I'
		},
		{
			Byte: 0x4A, // 'J'
		},
		{
			Byte: 0x4B, // 'K'
		},
		{
			Byte: 0x4C, // 'L'
		},
		{
			Byte: 0x4D, // 'M'
		},
		{
			Byte: 0x4E, // 'N'
		},
		{
			Byte: 0x4F, // 'O'
		},
		{
			Byte: 0x50, // 'P'
		},
		{
			Byte: 0x51, // 'Q'
		},
		{
			Byte: 0x52, // 'R'
		},
		{
			Byte: 0x53, // 'S'
		},
		{
			Byte: 0x54, // 'T'
		},
		{
			Byte: 0x55, // 'U'
		},
		{
			Byte: 0x56, // 'V'
		},
		{
			Byte: 0x57, // 'W'
		},
		{
			Byte: 0x58, // 'X'
		},
		{
			Byte: 0x59, // 'Y'
		},
		{
			Byte: 0x5A, // 'Z'
		},
		{
			Byte: 0x5B,
		},
		{
			Byte: 0x5C,
		},
		{
			Byte: 0x5D,
		},
		{
			Byte: 0x5E,
		},
		{
			Byte: 0x5F,
		},
		{
			Byte: 0x60,
		},
		{
			Byte: 0x61, // 'a'
			ExpectedByte:  10,
			ExpectedBool: true,
		},
		{
			Byte: 0x62, // 'b'
			ExpectedByte:  11,
			ExpectedBool: true,
		},
		{
			Byte: 0x63, // 'c'
			ExpectedByte:  12,
			ExpectedBool: true,
		},
		{
			Byte: 0x64, // 'd'
			ExpectedByte:  13,
			ExpectedBool: true,
		},
		{
			Byte: 0x65, // 'e'
			ExpectedByte:  14,
			ExpectedBool: true,
		},
		{
			Byte: 0x66, // 'f'
			ExpectedByte:  15,
			ExpectedBool: true,
		},
		{
			Byte: 0x67,
		},
		{
			Byte: 0x68,
		},
		{
			Byte: 0x69,
		},
		{
			Byte: 0x6A,
		},
		{
			Byte: 0x6B,
		},
		{
			Byte: 0x6C,
		},
		{
			Byte: 0x6D,
		},
		{
			Byte: 0x6E,
		},
		{
			Byte: 0x6F,
		},
		{
			Byte: 0x70,
		},
		{
			Byte: 0x71,
		},
		{
			Byte: 0x72,
		},
		{
			Byte: 0x73,
		},
		{
			Byte: 0x74,
		},
		{
			Byte: 0x75,
		},
		{
			Byte: 0x76,
		},
		{
			Byte: 0x77,
		},
		{
			Byte: 0x78,
		},
		{
			Byte: 0x79,
		},
		{
			Byte: 0x7A,
		},
		{
			Byte: 0x7B,
		},
		{
			Byte: 0x7C,
		},
		{
			Byte: 0x7D,
		},
		{
			Byte: 0x7E,
		},
		{
			Byte: 0x7F,
		},
		{
			Byte: 0x80,
		},
		{
			Byte: 0x81,
		},
		{
			Byte: 0x82,
		},
		{
			Byte: 0x83,
		},
		{
			Byte: 0x84,
		},
		{
			Byte: 0x85,
		},
		{
			Byte: 0x86,
		},
		{
			Byte: 0x87,
		},
		{
			Byte: 0x88,
		},
		{
			Byte: 0x89,
		},
		{
			Byte: 0x8A,
		},
		{
			Byte: 0x8B,
		},
		{
			Byte: 0x8C,
		},
		{
			Byte: 0x8D,
		},
		{
			Byte: 0x8E,
		},
		{
			Byte: 0x8F,
		},
		{
			Byte: 0x90,
		},
		{
			Byte: 0x91,
		},
		{
			Byte: 0x92,
		},
		{
			Byte: 0x93,
		},
		{
			Byte: 0x94,
		},
		{
			Byte: 0x95,
		},
		{
			Byte: 0x96,
		},
		{
			Byte: 0x97,
		},
		{
			Byte: 0x98,
		},
		{
			Byte: 0x99,
		},
		{
			Byte: 0x9A,
		},
		{
			Byte: 0x9B,
		},
		{
			Byte: 0x9C,
		},
		{
			Byte: 0x9D,
		},
		{
			Byte: 0x9E,
		},
		{
			Byte: 0x9F,
		},
		{
			Byte: 0xA0,
		},
		{
			Byte: 0xA1,
		},
		{
			Byte: 0xA2,
		},
		{
			Byte: 0xA3,
		},
		{
			Byte: 0xA4,
		},
		{
			Byte: 0xA5,
		},
		{
			Byte: 0xA6,
		},
		{
			Byte: 0xA7,
		},
		{
			Byte: 0xA8,
		},
		{
			Byte: 0xA9,
		},
		{
			Byte: 0xAA,
		},
		{
			Byte: 0xAB,
		},
		{
			Byte: 0xAC,
		},
		{
			Byte: 0xAD,
		},
		{
			Byte: 0xAE,
		},
		{
			Byte: 0xAF,
		},
		{
			Byte: 0xB0,
		},
		{
			Byte: 0xB1,
		},
		{
			Byte: 0xB2,
		},
		{
			Byte: 0xB3,
		},
		{
			Byte: 0xB4,
		},
		{
			Byte: 0xB5,
		},
		{
			Byte: 0xB6,
		},
		{
			Byte: 0xB7,
		},
		{
			Byte: 0xB8,
		},
		{
			Byte: 0xB9,
		},
		{
			Byte: 0xBA,
		},
		{
			Byte: 0xBB,
		},
		{
			Byte: 0xBC,
		},
		{
			Byte: 0xBD,
		},
		{
			Byte: 0xBE,
		},
		{
			Byte: 0xBF,
		},
		{
			Byte: 0xC0,
		},
		{
			Byte: 0xC1,
		},
		{
			Byte: 0xC2,
		},
		{
			Byte: 0xC3,
		},
		{
			Byte: 0xC4,
		},
		{
			Byte: 0xC5,
		},
		{
			Byte: 0xC6,
		},
		{
			Byte: 0xC7,
		},
		{
			Byte: 0xC8,
		},
		{
			Byte: 0xC9,
		},
		{
			Byte: 0xCA,
		},
		{
			Byte: 0xCB,
		},
		{
			Byte: 0xCC,
		},
		{
			Byte: 0xCD,
		},
		{
			Byte: 0xCE,
		},
		{
			Byte: 0xCF,
		},
		{
			Byte: 0xD0,
		},
		{
			Byte: 0xD1,
		},
		{
			Byte: 0xD2,
		},
		{
			Byte: 0xD3,
		},
		{
			Byte: 0xD4,
		},
		{
			Byte: 0xD5,
		},
		{
			Byte: 0xD6,
		},
		{
			Byte: 0xD7,
		},
		{
			Byte: 0xD8,
		},
		{
			Byte: 0xD9,
		},
		{
			Byte: 0xDA,
		},
		{
			Byte: 0xDB,
		},
		{
			Byte: 0xDC,
		},
		{
			Byte: 0xDD,
		},
		{
			Byte: 0xDE,
		},
		{
			Byte: 0xDF,
		},
		{
			Byte: 0xE0,
		},
		{
			Byte: 0xE1,
		},
		{
			Byte: 0xE2,
		},
		{
			Byte: 0xE3,
		},
		{
			Byte: 0xE4,
		},
		{
			Byte: 0xE5,
		},
		{
			Byte: 0xE6,
		},
		{
			Byte: 0xE7,
		},
		{
			Byte: 0xE8,
		},
		{
			Byte: 0xE9,
		},
		{
			Byte: 0xEA,
		},
		{
			Byte: 0xEB,
		},
		{
			Byte: 0xEC,
		},
		{
			Byte: 0xED,
		},
		{
			Byte: 0xEE,
		},
		{
			Byte: 0xEF,
		},
		{
			Byte: 0xF0,
		},
		{
			Byte: 0xF1,
		},
		{
			Byte: 0xF2,
		},
		{
			Byte: 0xF3,
		},
		{
			Byte: 0xF4,
		},
		{
			Byte: 0xF5,
		},
		{
			Byte: 0xF6,
		},
		{
			Byte: 0xF7,
		},
		{
			Byte: 0xF8,
		},
		{
			Byte: 0xF9,
		},
		{
			Byte: 0xFA,
		},
		{
			Byte: 0xFB,
		},
		{
			Byte: 0xFC,
		},
		{
			Byte: 0xFD,
		},
		{
			Byte: 0xFE,
		},
		{
			Byte: 0xFF,
		},
	}

	for testNumber, test := range tests {

		actualByte, actualBool := hexdecode(test.Byte)

		{
			expected := test.ExpectedBool
			actual   := actualBool

			if expected != actual {
				t.Errorf("For test #%d, the actual bool-value is not what was expected.", testNumber)
				t.Logf("EXPECTED: %t", expected)
				t.Logf("ACTUAL:   %t", actual)
				t.Logf("BYTE: 0x%02X (%d) (%q)", test.Byte, test.Byte, test.Byte)
				continue
			}
		}

		{
			expected := test.ExpectedByte
			actual   := actualByte

			if expected != actual {
				t.Errorf("For test #%d, the actual bool-value is not what was expected.", testNumber)
				t.Logf("EXPECTED: 0x%02X (%d)", expected, expected)
				t.Logf("ACTUAL:   0x%02X (%d)", actual, actual)
				t.Logf("BYTE: 0x%02X (%d) (%q)", test.Byte, test.Byte, test.Byte)
				continue
			}
		}
	}
}
