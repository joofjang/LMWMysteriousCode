package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	//var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		fmt.Printf("There is an err %s\n", err)
	}

	strSD := string(sd)
	strLength := len(strSD)
	charList := make([]string, strLength)

	for i, c := range strSD {
		charList[strLength-1-i] = string(c)
	}

	result := strings.ReplaceAll(strings.Join(charList, ""), ":", " ")

	fmt.Printf("%s.\n", result)
}
