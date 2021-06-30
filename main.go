package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/hauntarl/phone-number/db"
)

func main() {
	// 1. initialize database
	must(db.Open())
	defer db.Close()
	// 2. add phone_numbers table
	must(db.CreateTable())
	// 3. insert list of numbers in random order
	must(insert())
	// 4. display all the inserted numbers
	nums, err := db.SelectAll()
	must(err)
	fmt.Println("Phone Numbers Table...")
	for _, phn := range nums {
		fmt.Printf("uid: %3d, number: '%v'\n", phn.Uid, phn.Val)
	}
	fmt.Printf("total rows: %v\n", len(nums))
	fmt.Println("----------------------")
	// 5. normalize the phone numbers
	must(db.Normalize(nums))
	// 6. display all the numbers in a generic format after normalization
	nums, err = db.SelectAll()
	must(err)
	fmt.Println("Formatted Numbers...")
	for _, phn := range nums {
		fmt.Printf("uid: %3d, number: '%v'\n", phn.Uid, phn.Format())
	}
	fmt.Printf("total rows: %v\n", len(nums))
	fmt.Println("--------------------")
}

func insert() error {
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	for _, val := range numbers {
		_, err := db.InsertNumber(val)
		if err != nil {
			return err
		}
	}
	return nil
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
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
