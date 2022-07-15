package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)
	whatIsIt = string(sd)
	rune_arr := []rune(whatIsIt)
	var rev []rune
	for i := len(rune_arr) - 1; i >= 0; i-- {
		rev = append(rev, rune_arr[i])
	}
	fmt.Println(string(rev))
}
