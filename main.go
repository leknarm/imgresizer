package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"gopkg.in/gographics/imagick.v2/imagick"
	"log"
	"strings"
	"flag"
)

func main() {
	resizerDir := flag.String("dir", "", "directory")
	ratio := flag.Float64("ratio", 0.5, "resize ratio")

	flag.Parse()
	if len(*resizerDir) == 0 {
		log.Fatal("Not found directory to resize")
	}

	fmt.Printf("Resize images in folder %v\n", *resizerDir)

	imagick.Initialize()
    defer imagick.Terminate()

    files, err := ioutil.ReadDir(*resizerDir)
    if err != nil {
    	log.Fatal(err)
    }

    for _, file := range files {
    	if file.IsDir() {
    		continue
    	}

    	if !strings.HasSuffix(strings.ToLower(file.Name()), "jpg") && 
    		!strings.HasSuffix(strings.ToLower(file.Name()), "jpeg") &&
    		!strings.HasSuffix(strings.ToLower(file.Name()), "png") {
    			continue
    	}

    	fmt.Printf("Found file %v\n", file.Name())

    	f, err := os.Open(*resizerDir + file.Name())
		if err != nil {
    		log.Fatal(err)
    	}

    	mw := imagick.NewMagickWand()

    	err = mw.ReadImageFile(f)
		if err != nil {
			panic(err)
		}

		// Get original logo size
		width := float64(mw.GetImageWidth()) * (*ratio)
		height := float64(mw.GetImageHeight()) * (*ratio)

		// Calculate half the size
		hWidth := uint(width)
		hHeight := uint(height)

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

		mw.WriteImage(*resizerDir + "rz_" + file.Name())
		fmt.Printf("%v: resized to %vx%v\n", file.Name(), mw.GetImageWidth(), mw.GetImageHeight())

    }
}