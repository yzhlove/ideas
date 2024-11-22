package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"math/rand"
	"sync"
	"time"
)

var c *cache.Cache

func init() {
	c = cache.New(5*time.Minute, 10*time.Minute)
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			token := getGenerateCharacter(32)
			handle(token)
			time.Sleep(time.Second * 1)
			handle(token)
		}()
	}
	wg.Wait()
}

func handle(token string) {

	var tk *Token

	if ret, _, ok := c.GetWithExpiration(token); ok {
		tk = ret.(*Token)
		fmt.Println("---> ", tk.A)
		if tk.A != token {
			panic("token is error!")
		}
	}

	if tk == nil {
		tmp := new(Token)
		tmp.A = token
		tk = tmp
	}
	c.SetDefault(token, tk)
}

type Token struct {
	A string
}

var (
	character    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	characterLen = len(character)
)

func getGenerateCharacter(bit int) string {
	runes := make([]rune, bit)
	for i := 0; i < bit; i++ {
		runes[i] = rune(character[int(rand.Int31())%characterLen])
	}
	return string(runes)
}
