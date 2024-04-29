package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
//Solution function
func Solution(number string) string {
	//Replace All " " and "-"
	number = strings.ReplaceAll(number, " ", "")
	number = strings.ReplaceAll(number, "-", "")
	//Create slice 
	groups := make([]string, 0)
	for len(number) > 0 {
		//If len(number) > 4 will group 3 number
		if len(number) > 4 { 
			groups = append(groups, number[:3]) 
			number = number[3:]                  
		} else if len(number) == 4 { //If last string have 4 number will splitting and return number = "" or 0 
			groups = append(groups, number[:2]) 
			groups = append(groups, number[2:]) 
			number = ""                     
		} else { //Add 1 or 3 number as a group and return number = "" or 0
			groups = append(groups, number) 
			number = ""                  
		}
	}
	//Add "-" in every single string has slice
	formattedNumber := strings.Join(groups, "-")

	return formattedNumber
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Nhập số cần định dạng: ")
		scanner.Scan()
		number := scanner.Text()
		if len(number) < 2 || len(number) > 100 {
			fmt.Println("Nhập lại số có độ dài từ 2 đến 100 ký tự.")
			continue
		}

		formattedNumber := Solution(number)
		fmt.Println("Số đã được định dạng lại:", formattedNumber)
		break
	}
}
