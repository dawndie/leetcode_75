package week_0

/*
 * @lc app=leetcode id=735 lang=golang
 *
 * [735] Asteroid Collision
 */

// @lc code=start

func asteroidCollision(asteroids []int) []int {
	var result []int // Initialize the result stack to store asteroids that haven't collided

	for i := 0; i < len(asteroids); i++ {
		asteroid := asteroids[i]
		if len(result) == 0 || result[len(result)-1] < 0 || asteroid > 0 {
			result = append(result, asteroid)
		} else {
			top := result[len(result)-1]
			if -asteroid == top {
				result = result[:len(result)-1]
			} else if -asteroid > top {
				result = result[:len(result)-1]
				i--
			}
		}
	}

	return result
}

// @lc code=end
