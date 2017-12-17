package kmp

import (
	// "fmt"
)

func prebuilt(x string) []int {
	i, j := 1,0
	length := len(x)
	// fmt.Println(i,j,length)
	nxt_table  := make([]int ,length)

	for i < length {
		if x[i] != x[j] {
			if j > 0 {
				j = nxt_table[j-1]
			} else {
				i++
			}
		} else {
			// fmt.Println(i)
			nxt_table[i] = j+1
			j++
			i++
		}
	}
	return nxt_table	
} 

func kmp(s string, t string) []int {
	nxt,rst := prebuilt(t), []int{}
	size := len(nxt)
	var i,j int
	for i < len(s) {
		if j < size {
			if s[i] == t[j] {
				i++
				j++
			} else if j > 0 {
				j = nxt[j-1]				
			} else {
				i++
			}
		}
		if j == size {
			rst = append(rst,i-size)
			i++
			j = 0	
		}
	}
	return rst
}