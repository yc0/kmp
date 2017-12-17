package kmp

import (
	"testing"
	"fmt"
	"strings"
)

func TestNextTable(t *testing.T) {
	rst := prebuilt("abcdabca")
	if strings.Compare(array_to_string(rst,""),"00001231") != 0 {
		t.Error("prebuilt fails on case 1")
	} else {
		t.Log("passed prebuilt case 1")
	}
	rst = prebuilt("abcaby")
	if strings.Compare(array_to_string(rst,""),"000120") != 0 {
		t.Error("prebuilt fails on case 2")
	} else {
		t.Log("passed prebuilt case 2")
	}
	
}

func TestKMP(t *testing.T) {
	rst := kmp("abxabcabcaby","abcaby")
	if strings.Compare(array_to_string(rst," "),"6") != 0 { 
		t.Error("kmp fails on case 1")
	} else {
		t.Log("passed kmp case 1")
	}
	rst = kmp("abxabcabcabyaaababcabyb","abcaby")
	if strings.Compare(array_to_string(rst," "),"6 16") != 0 { 
		t.Error("kmp fails on case 2", rst)
	} else {
		t.Log("passed kmp case 2")
	}

	rst = kmp("cocacola","co")
	if strings.Compare(array_to_string(rst," "),"0 4") != 0 { 
		t.Error("kmp fails on case 3", rst)
	} else {
		t.Log("passed kmp case 3")
	}
}

func array_to_string(i []int, delim string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(i),"[]")," ", delim,-1)
}