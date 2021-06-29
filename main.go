package main

import (
	"fmt"
	"log"

	"github.com/hauntarl/phone-number/db"
)

func main() {
	check(db.Open())
	defer db.Close()

	check(db.CreateTable())

	_, err := db.InsertNumber("(223) 456-7890")
	check(err)

	nums, err := db.SelectAll()
	check(err)
	fmt.Println(nums)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
