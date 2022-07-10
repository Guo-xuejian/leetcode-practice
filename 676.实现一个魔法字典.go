/*
 * @lc app=leetcode.cn id=676 lang=golang
 *
 * [676] 实现一个魔法字典
 */

// @lc code=start
type MagicDictionary []string

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (d *MagicDictionary) BuildDict(dictionary []string) {
	*d = dictionary
}

func (d *MagicDictionary) Search(searchWord string) bool {
next:
	for _, word := range *d {
		if len(word) != len(searchWord) {
			continue
		}
		diff := false
		for i := range word {
			if word[i] != searchWord[i] {
				if diff {
					continue next
				}
				diff = true
			}
		}
		if diff {
			return true
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
// @lc code=end

