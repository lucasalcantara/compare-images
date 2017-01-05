package main

import (
	"fmt"
	"images"
)

func main() {
	imgPath1 := "../images/normal.jpg"
	imgPath2 := "../images/equal-normal.jpg"
	fmt.Println(images.CompareImages(imgPath2, imgPath1))
}
