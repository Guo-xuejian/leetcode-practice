// 面试题 16.04. 井字游戏
// 设计一个算法，判断玩家是否赢了井字游戏。输入是一个 N x N 的数组棋盘，由字符" "，"X"和"O"组成，其中字符" "代表一个空位。

// 以下是井字游戏的规则：

// 玩家轮流将字符放入空位（" "）中。
// 第一个玩家总是放字符"O"，且第二个玩家总是放字符"X"。
// "X"和"O"只允许放置在空位中，不允许对已放有字符的位置进行填充。
// 当有N个相同（且非空）的字符填充任何行、列或对角线时，游戏结束，对应该字符的玩家获胜。
// 当所有位置非空时，也算为游戏结束。
// 如果游戏结束，玩家不允许再放置字符。
// 如果游戏存在获胜者，就返回该游戏的获胜者使用的字符（"X"或"O"）；如果游戏以平局结束，则返回 "Draw"；如果仍会有行动（游戏未结束），则返回 "Pending"

func tictactoe(board []string) string {
	length := len(board)
	hasEmpty := false

	// 处理横排
	for i := 0; i < length; i++ {
		firstSymbol := string(board[i][0])
		hasDiff := false
		if firstSymbol == " " {
			hasEmpty = true
			continue
		}
		for j := 1; j < length; j++ {
			if string(board[i][j]) != firstSymbol {
				hasDiff = true
				break
			}
		}

		if !hasDiff {
			return firstSymbol
		}
	}

	// 处理竖排
	for i := 0; i < length; i++ {
		firstSymbol := string(board[0][i])
		hasDiff := false
		if firstSymbol == " " {
			hasEmpty = true
			continue
		}
		for j := 1; j < length; j++ {
			if string(board[j][i]) != firstSymbol {
				hasDiff = true
				if string(board[j][i]) == "" {
					hasEmpty = true
					break
				}
			}
		}

		if !hasDiff {
			return firstSymbol
		}
	}

	// 处理左斜对角
	firstSymbol := string(board[0][0])
	hasDiff := false
	if firstSymbol != " " {
		for i := 1; i < length; i++ {
			if firstSymbol != string(board[i][i]) {
				hasDiff = true
				break
			}
		}

		if !hasDiff {
			return firstSymbol
		}
	}

	// 处理右斜对角
	firstSymbol = string(board[0][length-1])
	hasDiff = false
	if firstSymbol != " " {
		for i := 0; i < length; i++ {
			if firstSymbol != string(board[i][length-1-i]) {
				hasDiff = true
				break
			}
		}

		if !hasDiff {
			return firstSymbol
		}
	}

	// 两边都没获胜
	if hasEmpty {
		return "Pending"
	} else {
		return "Draw"
	}
}
