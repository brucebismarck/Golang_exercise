
package countbits

import (
	"crypto/sha256"
)

// pc[i] is the population count of i
var pc [256]byte

// CompareBitDiff compares bit diff between s1 and s2
func CompareBitDiff(s1, s2 string) int{
	sha1 := sha256.Sum256([]byte(s1))
	sha2 := sha256.Sum256([]byte(s2))
	return bitDiff(sha1[:], sha2[:])
}

func bitDiff(sha1, sha2 []byte) int{
	count := 0
	for i :=0; i<len(sha1) || i<len(sha2); i++{
		switch {
		case i >= len(sha1):
			count += PopCount(sha2[i])
		case i >= len(sha2):
			count += PopCount(sha1[i])
		default:
			count += PopCount(sha1[i] ^ sha2[i])  //bitwise XOR
		}
	}
	return count
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x byte) int { // 如果叫 PopCountLoop 因为太长，会被叫做stutter 口吃。。。
	count := 0
	for ; x != 0; count++ {
		x &= x - 1
	}
	return count
}