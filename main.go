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

func resize(path string, ratio float64, quality int) {
	var err error
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	
	mw := imagick.NewMagickWand()

	err = mw.ReadImageFile(f)
	if err != nil {
		panic(err)
	}

	// Get original logo size
	width := float64(mw.GetImageWidth())
	height := float64(mw.GetImageHeight())

	// Calculate half the size
	hWidth := uint(width * ratio)
	hHeight := uint(height * ratio)

	// Resize the image using the Lanczos filter
	// The blur factor is a float, where > 1 is blurry, < 1 is sharp
	err = mw.ResizeImage(hWidth, hHeight, imagick.FILTER_LANCZOS, 1)
	if err != nil {
		panic(err)
	}

	// Set the compression quality to 95 (high quality = low compression)
	err = mw.SetImageCompressionQuality(uint(quality))
	if err != nil {
		panic(err)
	}

	mw.WriteImage(f.Name())
	fmt.Printf("Resized %vx%v -> %vx%v\n", width, height, hWidth, hHeight)
}

func resizeFile(path string, ratio float64, quality int) {
	resize(path, ratio, quality)
}

func resizeDirectory(dir string, ratio float64, quality int) {
	files, err := ioutil.ReadDir(dir)
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
    	resize(dir + file.Name(), ratio, quality)
    }
}

func main() {
	resizerDir := flag.String("dir", "", "Directory of images to resize")
	filePath := flag.String("file", "", "Path file to resize")
	ratio := flag.Float64("ratio", 0.5, "Resize ratio (0 > ratio < 1)")
	quality := flag.Int("quality", 95, "Image quality (default 95)")

	flag.Parse()
	if len(*resizerDir) == 0 && len(*filePath) == 0 {
		log.Fatal("Not found image file or directory to resize")
	}

	imagick.Initialize()
    defer imagick.Terminate()

	if len(*filePath) > 0 {
		fmt.Printf("Resize image: %v\n", *filePath)
		resizeFile(*filePath, *ratio, *quality)
	} else {
		fmt.Printf("Resize images in folder %v\n", *resizerDir)
		resizeDirectory(*resizerDir, *ratio, *quality)
	}

}