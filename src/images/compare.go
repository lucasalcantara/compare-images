package images

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"strings"
)

var (
	diffFolderPath = "../diff/"
)

func CompareImages(actualPath, expectedPath string) bool {
	equals := true
	rgba := color.RGBA64{}
	var bounds image.Rectangle
	actual, actualCfg := loadImage(actualPath)
	expImg, _ := loadImage(expectedPath)
	bounds = actual.Bounds()
	width := actualCfg.Width
	height := actualCfg.Height
	diffRGBA := image.NewRGBA64(image.Rectangle{image.Point{0, 0}, image.Point{bounds.Max.X, bounds.Max.Y}})

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			actR, actG, actB, actA := actual.At(x, y).RGBA()
			expR, expG, expB, expA := expImg.At(x, y).RGBA()
			if actR != expR || actG != expG || actB != expB || actA != expA {
				equals = false
				rgba = color.RGBA64{uint16(155), uint16(actG), uint16(actB), uint16(actA)}
			} else {
				rgba = color.RGBA64{uint16(actR), uint16(actG), uint16(actB), uint16(actA)}
			}

			diffRGBA.Set(x, y, rgba)
		}
	}

	if !equals {
		createDiffImage(actualPath, diffRGBA)
	}

	return equals
}

func loadImage(path string) (image.Image, image.Config) {
	imgFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer imgFile.Close()

	imgCfg, _, err := image.DecodeConfig(imgFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	imgFile.Seek(0, 0)
	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return img, imgCfg
}

func createDiffImage(actualPath string, diffRGBA image.Image) {
	values := strings.Split(actualPath, "/")
	name := "diff-" + values[len(values)-1]
	outfile, err := os.Create(diffFolderPath + name)
	if err != nil {
		// replace this with real error handling
		panic(err.Error())
	}
	defer outfile.Close()
	jpeg.Encode(outfile, diffRGBA, nil)
}
