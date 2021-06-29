package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/hauntarl/phone-number/db"
)

func main() {
	check(db.Open())
	defer db.Close()

	check(db.CreateTable())

	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	for _, num := range numbers {
		insert(num)
	}

	nums, err := db.SelectAll()
	check(err)
	display(nums)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func insert(number string) {
	_, err := db.InsertNumber(number)
	check(err)
}

func display(nums []db.PhoneNumber) {
	fmt.Println("-------------------")
	for _, phn := range nums {
		fmt.Printf("uid: %3d, number: '%v'\n", phn.Uid, phn.Format())
	}
	fmt.Printf("total rows: %v\n", len(nums))
	fmt.Println("-------------------")
}

var numbers = [...]string{
	"123 456 7891",
	"(123) 456 7892",
	"(123) 456-7893",
	"123-456-7894",
	"123-456-7890",
	"1234567892",
	"(123)456-7892",
	"(223) 456-7890",
	"223.456.7890",
	"223 456   7890   ",
	"123456789",
	"22234567890",
	"12234567890",
	"+1 (223) 456-7890",
	"321234567890",
	"123-abc-7890",
	"123-@:!-7890",
	"(023) 456-7890",
	"(123) 456-7890",
	"(223) 056-7890",
	"(223) 156-7890",
	"1 (023) 456-7890",
	"1 (123) 456-7890",
	"1 (223) 056-7890",
	"1 (223) 156-7890",
}
