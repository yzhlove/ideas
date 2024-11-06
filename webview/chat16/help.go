package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"math/rand"
	"sort"
	"strings"
)

const (
	httpTimestamp = "tm"
	httpNonce     = "nonce"
	httpSign      = "sign"
	httpSecret    = "*#06#*"
)

var (
	character    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	characterLen = len(character)
)

func MD5(input io.Reader) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, input); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func GenerateCharacter(bit int) string {
	runes := make([]rune, bit)
	for i := 0; i < bit; i++ {
		runes[i] = rune(character[int(rand.Int31())%characterLen])
	}
	return string(runes)
}

func GenerateSign(data map[string]string) (string, error) {

	strs := make([]string, 0, len(data))
	for k := range data {
		strs = append(strs, k)
	}

	// 自然排序
	sort.Strings(strs)
	var sb strings.Builder
	sb.WriteString(httpSecret)
	for k, v := range strs {
		sb.WriteString(v)
		sb.WriteString("=")
		sb.WriteString(data[v])
		if k != len(strs)-1 {
			sb.WriteString("&")
		}
	}
	sb.WriteString(httpSecret)
	return MD5(bytes.NewBufferString(sb.String()))
}
