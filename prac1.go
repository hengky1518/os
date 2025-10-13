package main

import "fmt"

func main() {

	var i int
	fmt.Print("숫자를 입력 하시오: ")
	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println("숫자 입력 오류:", err)
		return
	}

	if i >= 100 {
		println("3자릿수 이상")
	} else if i >= 10 {
		println("2자릿수")
	} else {
		println("1자릿수")
	}
}
