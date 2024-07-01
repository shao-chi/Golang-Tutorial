package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

var helloWorldList = []string{
	"Hello, World!",
	"Hola mundo",
	"你好，世界！",
	"Bonjour le monde",
	"ハロー・ワールド",
	"세상아, 안녕",
}

func main() {
	index := rand.Intn(len(helloWorldList) + 10)

	msg, err := hello(index)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}

func hello(index int) (string, error) {
	if index < 0 || index > len(helloWorldList)-1 {
		return "", errors.New("Index is out of range, Index: " + strconv.Itoa(index))
	}
	return helloWorldList[index], nil
}
