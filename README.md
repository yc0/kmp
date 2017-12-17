# KMP

![Golang version](https://img.shields.io/badge/golang-1.9.2-green.svg) ![passed](https://img.shields.io/badge/test-4%20passed%2C%200%20skipped-green.svg)[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![PR](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)

In computer science, the **Knuth–Morris–Pratt** string searching algorithm (or KMP algorithm) searches for occurrences of a "word" W within a main "text string" S by employing the observation that when a mismatch occurs, the word itself embodies sufficient information to determine where the next match could begin, thus bypassing re-examination of previously matched characters. 

reference [wikipedia](https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm)

## Complexity 
Time complexity O(M+N),

Space complexity O(N) 

## Install

`go get github.com/yc0/kmp`

## Usage


```go
package main

import (
  "fmt"
  "github.com/yc0/kmp"
)

func main() {
  fmt.Println(kmp("cocacola", "co"))
  //0, 4
  fmt.Println(kmp("abxabcabcabyaaababcabyb", "abcaby"))
  //6, 16
}	
```

## Inspired
[https://www.youtube.com/watch?v=GTJr8OvyEVQ](https://www.youtube.com/watch?v=GTJr8OvyEVQ)

## License
This package is licensed under MIT license. See LICENSE for details.