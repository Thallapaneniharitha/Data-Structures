// 1. permutation in go
package main

import "fmt"

func main() {
    sample := "Haritha@"
    sampleRune := []rune(sample)
    generatePermutation(sampleRune, 0, len(sampleRune)-1)
}

func generatePermutation(sampleRune []rune, left, right int) {
    if left == right {
        fmt.Println(string(sampleRune))
    } else {
        for i := left; i <= right; i++ {
            sampleRune[left], sampleRune[i] = sampleRune[i], sampleRune[left]
            generatePermutation(sampleRune, left+1, right)
            sampleRune[left], sampleRune[i] = sampleRune[i], sampleRune[left]
        }
    }
}

$go run main.go
Haritha@
Harith@a
Haritah@
Harita@h
Harit@ah
Harit@ha
Harihta@
Hariht@a
Harihat@
Hariha@t
Harih@at
Harih@ta
Hariaht@
Hariah@t
Hariath@
Hariat@h
Haria@th
Haria@ht
