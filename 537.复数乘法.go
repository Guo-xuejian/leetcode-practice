/*
 * @lc app=leetcode.cn id=537 lang=golang
 *
 * [537] 复数乘法
 */

// @lc code=start
func complexNumberMultiply(num1 string, num2 string) string {
	split1, split2 := strings.Split(num1, "+"), strings.Split(num2, "+")
	pre1, _ := strconv.Atoi(split1[0])
	late1, _ := strconv.Atoi(split1[1][:len(split1[1])-1])
	pre2, _ := strconv.Atoi(split2[0])
	late2, _ := strconv.Atoi(split2[1][:len(split2[1])-1])
	return fmt.Sprintf("%d+%di", pre1*pre2-(late1*late2), pre1*late2+pre2*late1)
}

// @lc code=end

