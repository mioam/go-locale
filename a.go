package main

import (
	"log"
	"main/locale"
)

func main() {

	tags, err := locale.DetectAll()
	if err != nil {
		log.Fatal(err)
	}
	// Get all available tags
	for i, t := range tags {

		println(i, t.String())
	}
}
