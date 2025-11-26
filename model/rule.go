package model

import "fmt"

// Checking for an empty file list
func CheckingFiles(files []string) {
	if len(files) == 0 {
		fmt.Println("List void")
		return
	}

}
