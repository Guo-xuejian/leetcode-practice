type RecentCounter struct {
	query []int // 切片模拟
}

// Constructor 初始化
func Constructor() RecentCounter {
	return RecentCounter{
		query: []int{},
	}
}

// Ping 函数
func (this *RecentCounter) Ping(t int) int {
	this.query = append(this.query, t)
	for this.query[0] < t-3000 { // 去除那些不再范围内的请求
		this.query = this.query[1:len(this.query)]
	}
	return len(this.query) // 长度即为请求数量
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */