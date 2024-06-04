package myutil

import "fmt"

func PrintMessage(message string) {
	fmt.Println(message);
}

func Calc(a int, b int) (int, int) {
    return a + 10, b * 2;
}