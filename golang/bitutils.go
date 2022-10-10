package main

import "unsafe"

// TODO: maybe change these to methods of som new struct

func setBit(n uint, pos uint) uint {
	n |= (1 << pos)
	return n
}

// there are other ways of doing this, but I like xor
func checkMask(n uint, mask uint) bool {
	return (n ^ mask) == 0
}

// set all bits where num is the number of bits to set;
// num is probably numPoints
func determineMask(num uint) uint {
	return (1 << num) - 1
}

// a function to return a string of the bit representation of a number
func bitString(n uint) string {
	var s string
	for i := uint(0); i < uint(8*unsafe.Sizeof(n)); i++ {
		if (n & (1 << uint(i))) != 0 {
			s = "1" + s
		} else {
			s = "0" + s
		}
	}
	return s
}
