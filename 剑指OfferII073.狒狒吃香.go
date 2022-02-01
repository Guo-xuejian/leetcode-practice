// 剑指 Offer II 073. 狒狒吃香蕉
// 狒狒喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。

// 狒狒可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉，下一个小时才会开始吃另一堆的香蕉。

// 狒狒喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。

// 返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。

// 示例 1：

// 输入: piles = [3,6,7,11], H = 8
// 输出: 4
// 示例 2：

// 输入: piles = [30,11,23,4,20], H = 5
// 输出: 30
// 示例 3：

// 输入: piles = [30,11,23,4,20], H = 6
// 输出: 23

// 提示：

// 1 <= piles.length <= 10^4
// piles.length <= H <= 10^9
// 1 <= piles[i] <= 10^9

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 0 // 速度的范围
	for _, num := range piles {
		if num > right {
			right = num
		}
	}
	// 二分查找适合的速度
	for left <= right {
		mid := (left + right) / 2
		timeCost := 0

		for _, num := range piles {
			if num%mid != 0 {
				timeCost += num/mid + 1 // 有余数向上取整 +1
			} else {
				timeCost += num / mid // 没有余数计算需要几个小时即可
			}
			if timeCost > h { // 已经超时不必继续,不能等于,大于和等于走的边界计算是不同的
				break
			}
		}

		if timeCost <= h {
			// 当前速度下可以吃完,试试减慢速度
			right = mid - 1
		} else {
			// 速度太慢,加快速度
			left = mid + 1
		}
	}
	return left
}