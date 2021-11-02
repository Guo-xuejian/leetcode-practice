/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 */

// @lc code=start
func convertToTitle(columnNumber int) string {
	tempList := []byte{}
	for columnNumber > 0 {
		temp := columnNumber % 26 // 获取余数
		if temp == 0 {            // 余数为0时需要特殊处理
			tempList = append(tempList, 'A'+byte(26-1)) // 或者 64 + byte[temp]
			columnNumber = (columnNumber - 26) / 26
		} else {
			tempList = append(tempList, 'A'+byte(temp-1))
			columnNumber = (columnNumber - temp) / 26
		}
	}
	// 由于反向写入，最终结果也需要反向
	for i, n := 0, len(tempList); i < n/2; i++ {
		tempList[i], tempList[n-1-i] = tempList[n-1-i], tempList[i]
	}
	return string(tempList)

}

// @lc code=end

