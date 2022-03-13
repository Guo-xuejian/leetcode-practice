package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Score int
}

type StuScores []Student

// Len
func (s StuScores) Len() int {
	return len(s)
}

// Less
func (s StuScores) Less(i, j int) bool {
	return s[i].Score > s[j].Score
}

// Swap
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	stus := StuScores{
		{"guo", 95},
		{"zhang", 97},
		{"xu", 98},
		{"jiang", 99},
	}
	// 未排序
	fmt.Println("Default:\n\t", stus)
	// StuScores 实现了 sort.Interface 接口，所以可以调用 Sort 函数进行排序
	sort.Sort(stus)
	// 判断是否已经排序，排序好了会打印 true
	fmt.Println("IS Sorted?\n\t", sort.IsSorted(stus))
	// 排序好的 stus 数据
	fmt.Println("Sorter\n\t", stus)
}
