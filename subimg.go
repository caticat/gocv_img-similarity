package main

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	f, e := os.OpenFile("../go_screenshot/20201023-145649.jpg", os.O_RDONLY, 0755)
	if e != nil {
		log.Println("1e:", e)
		return
	}
	defer f.Close()
	i, n, e := image.Decode(f)
	if e != nil {
		log.Println("2e:", e)
		return
	}
	log.Println("n:", n)
	rgbImg := i.(*image.YCbCr)
	subImg := rgbImg.SubImage(image.Rect(0, 0, 50, 50)).(*image.YCbCr)
	f1, e := os.Create("tmp.jpg")
	if e != nil {
		log.Println("3e:", e)
		return
	}
	defer f1.Close()
	jpeg.Encode(f1, subImg, nil)
	log.Println("done")
}
