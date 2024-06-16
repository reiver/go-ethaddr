package ethaddr

// hexdecode decodes the ASCII / UTF-8 hexadecimal symbol into its numberic value.
//
// For example —
//
//	'0' ->  0
//	'1' ->  1
//	'2' ->  2
//	'3' ->  3
//	'4' ->  4
//	'5' ->  5
//	'6' ->  6
//	'7' ->  7
//	'8' ->  8
//	'9' ->  9
//	'A' -> 10
//	'B' -> 11
//	'C' -> 12
//	'D' -> 13
//	'E' -> 14
//	'F' -> 15
//	'a' -> 10
//	'b' -> 11
//	'c' -> 12
//	'd' -> 13
//	'e' -> 14
//	'f' -> 15
func hexdecode(b byte) (byte, bool) {
	switch {
	case '0' <= b && b <= '9':
		return (b - '0'), true
	case 'A' <= b && b <= 'F':
		return (b - 'A' + 10), true
	case 'a' <= b && b <= 'f':
		return (b - 'a' + 10), true
	default:
		return 0, false
	}
}
