package main

import (
	"fmt"
	"strings"
)

func Decode(cipherText string, numRails int) string {
	length := len(cipherText)
	plainText := make([]string, length)
	for from, to := range MakeCipher(len(cipherText), numRails) {
		plainText[to] = string(cipherText[from])
	}
	return strings.Join(plainText, "")
}

func MakeCipher(items int, numRails int) []int {
	var rails [][]int
	rail, dRail := 0, 1
	for i := 0; i < items; i++ {
		if rail <= numRails {
			rails = append(rails, []int{})
		}
		rails[rail] = append(rails[rail], i)
		if rail+dRail < 0 || numRails <= rail+dRail {
			dRail = -dRail
		}
		rail += dRail
	}
	result := []int{}
	for _, slice := range rails {
		result = append(result, slice...)
	}
	return result
}

func CutFreqently(txt string) string {
	var remember [256]int
	for i := 0; i < 256; i++ {
		remember[i] = 0
	}
	for i := 0; i < len(txt); i++ {
		remember[int(txt[i])] = remember[int(txt[i])] + 1
	}
	xx := txt
	for i := 0; i < 256; i++ {
		if remember[i] >= 10 {
			xx = strings.Replace(xx, string(i), "", -1)
		}
	}
	return xx
}

func main() {
	var whatIsIt string
	var temp string
	secret := "CYtZBsWZaZliYZocWLZlXuZZYWYeYXZsXeZXtXWpXeRYYYd!ZnYeWXoYXasnX,WXWrWPoAdWesnciGenWr"
	temp = CutFreqently(secret)
	whatIsIt = Decode(temp, 4)
	fmt.Println(string(whatIsIt))
}
