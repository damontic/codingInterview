// Package rotateimage contains a function that allows to rotate an image 90 degrees
package rotateimage

import "image/color"

// Function Rotateimage rotates an image 90 degrees
func Rotateimage(matrix [][]color.RGBA) [][]color.RGBA {
	n := len(matrix)
	newM := make([][]color.RGBA, n)

	for i := 0; i < n; i++ {
		newM[i] = make([]color.RGBA, n)
	}

	for i := 0; i < n; i++ {
		for j:=0; j < n; j++ {
			newM[i][j] = matrix[n-1-j][i]
		}
	}
	return newM
}

// Function RotateimageInPlace rotates an image 90 degress in place
func RotateimageInPlace(matrix [][]color.RGBA) [][]color.RGBA {
        n := len(matrix)
        for i := 0; i < n*n; i++ {
	        matrix[i][j] = matrix[n-1-j][i]
        }
        return matrix
}
