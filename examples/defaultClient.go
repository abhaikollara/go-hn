package main

import (
	"fmt"
	"github.com/abhaikollara/hn"
)

func main() {
	item, _ := hn.GetItem(1)
	fmt.Println(item.Type, item.Text, len(item.Kids))
	kids, _ := hn.GetItems(item.Kids)
	for _, i := range kids {
		fmt.Println(i.Type, i.Text)
	}
}
