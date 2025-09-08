package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct {
	width int
	height int
	data [][]byte
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
	pixelValue := i.data[y][x]
	return color.RGBA{pixelValue, pixelValue, 255, 255}
}

func (i *Image) compute() {
	for y := 0; y < i.height; y++ {
		for x := 0; x < i.width; x++ {
			i.data[y][x] = byte(x * y)
		}
	}
}

func NewImage(width, height int) Image {
	data := make([][]byte, height)
	for row := 0; row < height; row++ {
		data[row] = make([]byte, width)
	}
	return Image{width: width, height: height, data: data}
}

func main() {
	m := NewImage(100, 100)
	//m.compute()
	pic.ShowImage(m)
}
