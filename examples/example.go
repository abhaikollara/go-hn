package main

import (
	"fmt"
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

	user, err := c.GetUser("TheThirdTuring")
	if err != nil {
		log.Fatal("error: ", err)
	}
	log.Printf("%#v", user)

	items, err := c.GetTopStoryIDs()
	if err != nil {
		log.Fatal("error: ", err)
	}
	stories, _ := c.WithConcurrency(100).GetItems(items)
	for _, s := range stories {
		fmt.Println(s)
	}
}
