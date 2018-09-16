package rotateimage

import (
	"fmt"
	"image/color"
)

func ExampleRotateimage() {
	image := [][]color.RGBA{
		{color.RGBA{0,0,0,0},color.RGBA{0,1,0,0},color.RGBA{0,2,0,0},color.RGBA{0,3,0,0}},
		{color.RGBA{1,0,0,0},color.RGBA{1,1,0,0},color.RGBA{1,2,0,0},color.RGBA{1,3,0,0}},
		{color.RGBA{2,0,0,0},color.RGBA{2,1,0,0},color.RGBA{2,2,0,0},color.RGBA{2,3,0,0}},
		{color.RGBA{3,0,0,0},color.RGBA{3,1,0,0},color.RGBA{3,2,0,0},color.RGBA{3,3,0,0}}}

	newImage := Rotateimage(image)
	fmt.Println(newImage)
	// Output:
	// [[{3 0 0 0} {2 0 0 0} {1 0 0 0} {0 0 0 0}] [{3 1 0 0} {2 1 0 0} {1 1 0 0} {0 1 0 0}] [{3 2 0 0} {2 2 0 0} {1 2 0 0} {0 2 0 0}] [{3 3 0 0} {2 3 0 0} {1 3 0 0} {0 3 0 0}]]
}
