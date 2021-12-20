package main

//  插入排序（英语：Insertion Sort）是一种简单直观的排序算法。它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。

func InsertionSort(arr *[]int) {
	length := len(*arr)
	for i := 1; i < length; i++ {
		temp := (*arr)[i]
		j := i - 1
		for j >= 0 && temp < (*arr)[j] {
			(*arr)[j+1] = (*arr)[j]
			j--
		}
		(*arr)[j+1] = temp
	}
}
