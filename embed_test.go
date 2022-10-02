package golang_embed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

//go:embed version.txt
var version2 string

func TestString(t *testing.T) {
	// saat mengisi golang embed, komentar nya jangan dikasih spasi di go:embed nya
	fmt.Println(version)
	fmt.Println(version2)
}

//go:embed db-normal.jpg
var logo []byte

func TestEmbedByte(t *testing.T) {
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
