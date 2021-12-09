/*
 * @lc app=leetcode.cn id=677 lang=golang
 *
 * [677] 键值映射
 */

// @lc code=start
type MapSum map[string]int

func Constructor() MapSum {
	return MapSum{} // 初始化
}

func (this *MapSum) Insert(key string, val int) {
	(*this)[key] = val // 创建或者覆盖
}

func (this *MapSum) Sum(prefix string) (sum int) {
	for key, val := range *this { // 遍历 map，符合条件则累加
		if strings.HasPrefix(key, prefix) {
			sum += val
		}
	}
	return
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
// @lc code=end

