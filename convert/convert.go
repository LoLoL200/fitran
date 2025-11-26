package convert

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// Convert file Images
func ConvertImage(src string, dstExt string, quality int, deleteOriginal bool) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	dst := strings.TrimSuffix(src, filepath.Ext(src)) + dstExt

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// Selecting a format
	switch dstExt {
	case ".jpg", ".jpeg":
		if err := jpeg.Encode(out, img, &jpeg.Options{Quality: quality}); err != nil {
			return err
		}
	case ".png":
		if err := png.Encode(out, img); err != nil {
			return err
		}
	default:
		return fmt.Errorf("неподдерживаемый формат: %s", dstExt)
	}

	if deleteOriginal {
		return os.Remove(src)
	}

	return nil
}

// Convert JPG in PNG
func ConvertJPGtoPNG(src string, deleteOriginal bool) error {
	return ConvertImage(src, ".png", 0, deleteOriginal)
}

// Convert PNG in JPG
func ConvertPNGtoJPG(src string, quality int, deleteOriginal bool) error {
	return ConvertImage(src, ".jpg", quality, deleteOriginal)
}
