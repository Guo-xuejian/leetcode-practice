# 编辑距离
# m*n的grid中，grid[i][j]为1代表一个疫情小区，它的上下左右需要有一支疫情防控队伍进行管控，为0则是正常小区，求需要多少个队伍，ac了

# https://leetcode-cn.com/problems/edit-distance/
# 72. 编辑距离
# 给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

# 你可以对一个单词进行如下三种操作：

# 插入一个字符
# 删除一个字符
# 替换一个字符
 

# 示例 1：

# 输入：word1 = "horse", word2 = "ros"
# 输出：3
# 解释：
# horse -> rorse (将 'h' 替换为 'r')
# rorse -> rose (删除 'r')
# rose -> ros (删除 'e')
# 示例 2：

# 输入：word1 = "intention", word2 = "execution"
# 输出：5
# 解释：
# intention -> inention (删除 't')
# inention -> enention (将 'i' 替换为 'e')
# enention -> exention (将 'n' 替换为 'x')
# exention -> exection (将 'n' 替换为 'c')
# exection -> execution (插入 'u')
 

# 提示：

# 0 <= word1.length, word2.length <= 500
# word1 和 word2 由小写英文字母组成