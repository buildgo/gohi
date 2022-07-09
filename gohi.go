package main

import (
	"math/rand"
	"time"
)

var hiInManyLanguages []string

func init() {
	hiWords := []string{"Hi", "Hello", "안녕하세요", "こんにちわ", "你好", "Hola"}
	hiInManyLanguages = hiWords
	rand.Seed(time.Now().UnixNano())
}

func sayHiRandomly() string {
	i := rand.Intn(len(hiInManyLanguages))
	return hiInManyLanguages[i]
}

func SayHi() string {
	return sayHiRandomly()
}
