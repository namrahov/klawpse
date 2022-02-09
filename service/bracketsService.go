package service

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	log "github.com/sirupsen/logrus"
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

	rows := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {

			if isValid(colCell) {
				fmt.Println("Dogru")
			} else {
				fmt.Println("Yalnis")
			}
			break
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
