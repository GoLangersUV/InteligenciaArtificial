package utils

import (
	"image"
	"image/png"
	"log"
	"os"
)

func LoadIcon(path string) image.Image {
	// Open the icon file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the image
	icon, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	return icon
}
