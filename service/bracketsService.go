package service

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type IBracketsService interface {
	DetectBracketsType(filePath string) error
}

type BracketsService struct {
}

func (b *BracketsService) DetectBracketsType(filePath string) error {

	//ornek olsun diye filePath-a bir deger veriyorum
	filePath = "template/brackets.xlsx"

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for i := 2; i < 10; i++ {
		brackets := f.GetCellValue("Sheet1", "A"+strconv.Itoa(i))
		if len(brackets) == 0 {
			break
		}

		if isValid(brackets) {
			fmt.Println("Dogru")
		} else {
			fmt.Println("Yalnis")
		}
	}

	return nil
}

func isValid(brackets string) bool {

	var stack []int32

	for _, bracket := range brackets {
		n := len(stack) - 1

		if bracket == '}' {
			if n < 0 {
				return false
			}
			current := stack[n]
			stack = stack[:n]
			if current != '{' {
				return false
			}
		} else if bracket == ']' {
			if n < 0 {
				return false
			}
			current := stack[n]
			stack = stack[:n]
			if current != '[' {
				return false
			}
		} else if bracket == ')' {
			if n < 0 {
				return false
			}
			current := stack[n]
			stack = stack[:n]
			if current != '(' {
				return false
			}
		} else {
			stack = append(stack, bracket)
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
