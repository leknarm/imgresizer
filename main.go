package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"gopkg.in/gographics/imagick.v2/imagick"
	"log"
)

func main() {
	var resizerDir = "/Users/tantai/Desktop/imgresizer_test/"

	imagick.Initialize()
    defer imagick.Terminate()

    files, err := ioutil.ReadDir(resizerDir)
    if err != nil {
    	log.Fatal(err)
    }

    for _, file := range files {
    	f, err := os.Open(resizerDir + file.Name())
		if err != nil {
    		log.Fatal(err)
    	}

    	mw := imagick.NewMagickWand()

    	err = mw.ReadImageFile(f)
		if err != nil {
			panic(err)
		}

		// Get original logo size
		width := mw.GetImageWidth()
		height := mw.GetImageHeight()

		// Calculate half the size
		hWidth := uint(width / 2)
		hHeight := uint(height / 2)

		// Resize the image using the Lanczos filter
		// The blur factor is a float, where > 1 is blurry, < 1 is sharp
		err = mw.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS, 1)
		if err != nil {
			panic(err)
		}

		// Set the compression quality to 95 (high quality = low compression)
		err = mw.SetImageCompressionQuality(95)
		if err != nil {
			panic(err)
		}

		mw.WriteImage(resizerDir + "rz_" + file.Name())
		fmt.Printf("%v: resized to %vx%v\n", file.Name(), mw.GetImageWidth(), mw.GetImageHeight())

    }
}