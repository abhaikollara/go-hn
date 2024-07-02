package main

import (
	"github.com/abhaikollara/hn"
	"log"
)

func main() {
	c := hn.New()
	item, err := c.GetItem(1)
	if err != nil {
		log.Fatal("error: ", err)
	}
	log.Printf("%#v", item)
}
