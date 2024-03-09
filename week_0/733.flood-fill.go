package week_0

/*
 * @lc app=leetcode id=733 lang=golang
 *
 * [733] Flood Fill
 */

// @lc code=start
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] != newColor {
		dfs(image, sr, sc, image[sr][sc], newColor)
	}
	return image
}

func dfs(image [][]int, x, y, sourceColor, newColor int) {
	if x < 0 || x >= len(image) || y < 0 || y >= len(image[0]) || image[x][y] != sourceColor {
		return
	}
	image[x][y] = newColor
	dfs(image, x+1, y, sourceColor, newColor)
	dfs(image, x-1, y, sourceColor, newColor)
	dfs(image, x, y+1, sourceColor, newColor)
	dfs(image, x, y-1, sourceColor, newColor)
}
// slice, map, channel alway pass by reference
// @lc code=end
