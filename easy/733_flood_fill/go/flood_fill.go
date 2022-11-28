package leetcode

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	visited := make([][]bool, len(image))
	for i := 0; i < len(image); i++ {
		visited[i] = make([]bool, len(image[i]))
	}
	return fill(image, visited, sr, sc, color, image[sr][sc])
}

func fill(image [][]int, visited [][]bool, sr, sc, color, orig int) [][]int {
	if image[sr][sc] == orig && !visited[sr][sc] {
		image[sr][sc] = color
		visited[sr][sc] = true
	} else {
		return image
	}
	if sr > 0 {
		image = fill(image, visited, sr-1, sc, color, orig)
	}
	if sc > 0 {
		image = fill(image, visited, sr, sc-1, color, orig)
	}
	if sr < len(image)-1 {
		image = fill(image, visited, sr+1, sc, color, orig)
	}
	if sc < len(image[sr])-1 {
		image = fill(image, visited, sr, sc+1, color, orig)
	}
	return image
}
