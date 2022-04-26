// 剑指 Offer 45. 把数组排成最小的数
// 输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

// 示例 1:

// 输入: [10,2]
// 输出: "102"
// 示例 2:

// 输入: [3,30,34,5,9]
// 输出: "3033459"

// 提示:

// 0 < nums.length <= 100
// 说明:

// 输出结果可能非常大，所以你需要返回一个字符串而不是整数
// 拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0

func minNumber(nums []int) string {
	// 快排实现排序，排序后转成string
	res := make([]string, len(nums))
	for i, v := range nums {
		res[i] = strconv.Itoa(v)
	}
	compare := func(str1, str2 string) bool {
		num1, _ := strconv.Atoi(str1 + str2)
		num2, _ := strconv.Atoi(str2 + str1)
		if num1 < num2 {
			return true
		}
		return false
	}
	var quickSort func(strArr []string, l, r int)
	quickSort = func(strArr []string, l, r int) {
		if l >= r {
			return
		}
		i, j, pivot := l, r, strArr[l]
		for i < j {
			for i < j && compare(pivot, strArr[j]) {
				j--
			}
			for i < j && !compare(pivot, strArr[i]) {
				i++
			}
			strArr[i], strArr[j] = strArr[j], strArr[i]
		}
		strArr[i], strArr[l] = strArr[l], strArr[i]
		quickSort(strArr, l, i-1)
		quickSort(strArr, i+1, r)
	}
	quickSort(res, 0, len(nums)-1)
	ans := ""
	for _, s := range res {
		ans += s
	}
	return ans
}
