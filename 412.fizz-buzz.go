/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) (res []string) {
	for i := 1; i <= n; i++ {
		temp := ""
		if i%3 == 0 {
			temp += "Fizz"
		}
		if i%5 == 0 {
			temp += "Buzz"
		}
		if len(temp) == 0 {
			temp = strconv.Itoa(i)
		}
		res = append(res, temp)
	}
	return
}

// @lc code=end

