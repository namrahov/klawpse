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

			if areBracketsShaped(colCell) {
				fmt.Println("Dogru")
			} else {
				fmt.Println("Yalnis")
			}
			break
		}
	}

	return nil
}

func areBracketsShaped(brackets string) bool {

	var list []string

	for i := 0; i < len(brackets); i++ {
		var bracket = string(brackets[i])

		if bracket == "(" || bracket == "[" || bracket == "{" {
			list = append(list, bracket)
			continue
		}

		if len(list) == 0 {
			return false
		}

		var check string
		switch bracket {
		case ")":
			check = list[len(list)-1]
			list = list[:len(list)-1]
			if check == "{" || check == "[" {
				return false
			}
			break
		case "}":
			check = list[len(list)-1]
			list = list[:len(list)-1]
			if check == "(" || check == "[" {
				return false
			}
			break
		case "]":
			check = list[len(list)-1]
			list = list[:len(list)-1]
			if check == "(" || check == "{" {
				return false
			}
			break
		}
	}

	if len(list) == 0 {
		return true
	}

	return false
}
