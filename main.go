package main

import "fmt"

func runsh(command string, args []string) (s string, err error) {
	fmt.Println(command)
	return command, nil
}
func main() {
	var input string
	var args []string
	for {
		fmt.Scanln(&input)
		//fmt.Println(input)
		runsh(input, args)
	}
}
