/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	index1, index2 := len(a)-1, len(b)-1
	byteSlice := []byte{}
	addStatus := false
	for index1 >= 0 || index2 >= 0 {
		num1, num2, sumNow := 0, 0, 0
		if index1 > -1 {
			num1 = int(a[index1] - '0')
			index1--
		}
		if index2 > -1 {
			num2 = int(b[index2] - '0')
			index2--
		}
		sumNow = num1 + num2
		if addStatus {
			sumNow++
		}
		if sumNow >= 2 {
			addStatus = true
			byteSlice = append(byteSlice, []byte(strconv.Itoa(sumNow%2))...)
		} else {
			addStatus = false
			byteSlice = append(byteSlice, []byte(strconv.Itoa(sumNow))...)
		}
	}
	if addStatus {
		byteSlice = append(byteSlice, '1')
	}
	// 翻转 byteSlice
	length := len(byteSlice)
	for i := length - 1; i >= (length+1)/2; i-- {
		byteSlice[i], byteSlice[length-1-i] = byteSlice[length-1-i], byteSlice[i]
	}
	return string(byteSlice)
}

// @lc code=end

