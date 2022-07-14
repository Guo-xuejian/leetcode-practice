/*
 * @lc app=leetcode.cn id=735 lang=golang
 *
 * [735] 行星碰撞
 */

// @lc code=start
func asteroidCollision(asteroids []int) (st []int) {
	for _, aster := range asteroids {
		alive := true
		for alive && aster < 0 && len(st) > 0 && st[len(st)-1] > 0 {
			alive = st[len(st)-1] < -aster // aster 是否存在
			if st[len(st)-1] <= -aster {   // 栈顶行星爆炸
				st = st[:len(st)-1]
			}
		}
		if alive {
			st = append(st, aster)
		}
	}
	return
}

// @lc code=end

