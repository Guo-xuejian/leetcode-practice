/*
 * @lc app=leetcode.cn id=649 lang=golang
 *
 * [649] Dota2 参议院
 */

// @lc code=start
func predictPartyVictory(senate string) string {
	var radiant, dire []int
	senateLength := len(senate)
	for index, v := range senate {
		if v == 'R' { // radiant 入队
			radiant = append(radiant, index)
		} else { // dire 入队
			dire = append(dire, index)
		}
	}
	// 当双方都仍然有议员在场则就继续新一轮投票
	for len(radiant) > 0 && len(dire) > 0 {
		if radiant[0] < dire[0] { // radiant 议员在前投票，禁止对方队首议员投票权后该议员进入队尾等待下一次投票权
			radiant = append(radiant, radiant[0]+senateLength)
		} else { // dire 议员在前投票，后续同上
			dire = append(dire, dire[0]+senateLength)
		}
		// 双方队首均出列，只不过被禁止的议员永久出队
		radiant = radiant[1:]
		dire = dire[1:]
	}
	// 最后判定双方哪一队不为空
	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}

// @lc code=end

