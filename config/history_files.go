package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

type FileEvent struct {
	Time      time.Time // file modification time in the OS
	ImageTime time.Time // time from EXIF
	Event     string
	FilePath  string
}

var history []FileEvent

// Get time from EXIF ​​(JPEG only)
func GetImageTime(path string) (time.Time, error) {
	f, err := os.Open(path)
	if err != nil {
		return time.Time{}, err
	}
	defer f.Close()

	x, err := exif.Decode(f)
	if err != nil {
		return time.Time{}, err
	}

	t, err := x.DateTime()
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

// Рекурсивно ищем все файлы JPG и PNG
func SearchImages(rootDir string) ([]string, error) {
	var files []string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := strings.ToLower(filepath.Ext(path))
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
				files = append(files, path)
			}
		}
		return nil
	})
	return files, err
}

func HistoryFiles(rootDir string) {
	history = []FileEvent{} // очищаем историю перед добавлением

	files, err := SearchImages(rootDir)
	// Error files
	if err != nil {
		fmt.Println("File search error:", err)
		return
	}
	//if there are no files
	if len(files) == 0 {
		fmt.Println("Files not found")
		return
	}

	// Iterates through the list of files
	for _, file := range files {

		// Get file informations
		fileInfo, err := os.Stat(file)
		if err != nil {
			fmt.Println("Ошибка при получении информации о файле:", err)
			continue
		}

		// Determining the time of image creation
		var imgTime time.Time
		ext := strings.ToLower(filepath.Ext(file))
		if ext == ".jpg" || ext == ".jpeg" {
			imgTime, _ = GetImageTime(file)
		}

		//Adding an event to history
		history = append(history, FileEvent{
			Time:      fileInfo.ModTime(),
			ImageTime: imgTime,
			Event:     "checked",
			FilePath:  file,
		})
	}

	// Get history
	fmt.Println("=================HISTORY============================")
	for index, fileEvent := range history {
		fmt.Printf(
			"%d. FileTime: %s | ImageTime: %s | Path: %s\n",
			index+1,
			fileEvent.Time.Format("2006-01-02 15:04:05"),
			fileEvent.ImageTime.Format("2006-01-02 15:04:05"),
			fileEvent.FilePath,
		)
	}
}
