package grind75

// An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.

// You are also given three integers sr, sc, and color. You should perform a flood fill on the image starting from the pixel image[sr][sc].

// To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color), and so on. Replace the color of all of the aforementioned pixels with color.

// Return the modified image after performing the flood fill.

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	if color != newColor {
		changeColor(image, sr, sc, &color, &newColor)
	}
	return image
}

func changeColor(image [][]int, x, y int, color, newColor *int) {
	if image[x][y] == *color {
		image[x][y] = *newColor
		if x >= 1 {
			changeColor(image, x-1, y, color, newColor)
		}
		if y >= 1 {
			changeColor(image, x, y-1, color, newColor)
		}
		if x+1 < len(image) {
			changeColor(image, x+1, y, color, newColor)
		}
		if y+1 < len(image[0]) {
			changeColor(image, x, y+1, color, newColor)
		}
	}
}
