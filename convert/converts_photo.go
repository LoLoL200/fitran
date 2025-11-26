package convert

import (
	"fitran/config"
	"fitran/model"
	"fmt"
)

// Convert in .jpg
func FileInJPG() {
	files := config.SearchPngFile()
	model.CheckingFiles(files)

	for _, filePath := range files {
		if files == nil {
			fmt.Println("List void")
		}
		if err := ConvertPNGtoJPG(filePath, 90, true); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Converted:", filePath)
		}
	}
}

// Convert in .png
func FileInPNG() {
	files := config.SearchJpgFile()
	model.CheckingFiles(files)

	for _, filePath := range files {
		if err := ConvertJPGtoPNG(filePath, true); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Converted:", filePath)
		}
	}
}
