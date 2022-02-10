package service

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type IBracketsService interface {
	DetectPalindromeOfNumber(filePath string) error
}

type PalindromeService struct {
}

func (b *PalindromeService) DetectPalindromeOfNumber(filePath string) error {

	//ornek olsun diye filePath-a bir deger veriyorum
	filePath = "template/palindrome.xlsx"

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	rows := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, numberString := range row {

			number, err := strconv.ParseInt(strings.Split(numberString, ".")[0], 10, 64)
			if err != nil {
				panic(err)
			}
			findPalindrome(number)
			break
		}
	}

	return nil
}

func findPalindrome(number int64) {
	var count int64
	for !isPalindrome(number) {
		number += reverseNumber(number)
		count++
	}

	fmt.Println(number, " ", count)
}

func reverseNumber(number int64) int64 {
	var reverse int64
	var temp int64

	for number != 0 {
		temp = number % 10
		reverse = (reverse * 10) + temp
		number = number / 10
	}

	return reverse
}

func isPalindrome(number int64) bool {
	return reverseNumber(number) == number
}
