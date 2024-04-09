package main

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"os"
)

func main() {

	bytes1, err := os.ReadFile("/Users/yurisa/Desktop/MockSave/DataServer/WorldMapList.csv")
	if err != nil {
		panic(err)
	}

	bytes2, err := os.ReadFile("/Users/yurisa/Desktop/MockSave/TempServers/WorldMapList.csv")
	if err != nil {
		panic(err)
	}

	df := diffmatchpatch.New()
	//r := df.DiffMain(string(bytes1), string(bytes2), false)
	r := df.DiffMainRunes([]rune(string(bytes1)), []rune(string(bytes2)), false)
	fmt.Println(df.DiffPrettyText(r))
}
