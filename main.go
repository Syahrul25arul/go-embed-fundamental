package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed db-normal.jpg
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	// dengan menggunakan go build, file yang di embed akan di load dan disimpan kedalam binary golang. mejadikan file yang di embed tidak akan bisa diubah ketika di embed
	// keuntungannya adalah ketika build dipindahkan, file yang di embed tetap akan ada meskipun file originalnya tidak di ikut sertakan, karna sudah disimpan di file binary golang

	fmt.Println(version)
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(content))
		}
	}
}
