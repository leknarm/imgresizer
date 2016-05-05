# imgresizer
Batch resize images with ratio

# usages
go run main.go -dir=/path/to/your/imgs/ -ratio=0.5

# options
```
go run main.go --help  
  -dir string  
    	Directory of images to resize
  -file string
    	Path file to resize
  -quality int
    	Image quality (default 95) (default 95)
  -ratio float
    	Resize ratio (0 > ratio < 1) (default 0.5)
```
