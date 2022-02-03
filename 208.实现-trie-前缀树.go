/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	chidren [26]*Trie
	isEnd   bool // 是否是当前字符串结尾
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.chidren[ch] == nil {
			node.chidren[ch] = &Trie{}
		}
		node = node.chidren[ch]
	}
	// 当前这个单词暂时不会有其余字符，所以 node.isEnd == true
	node.isEnd = true
}

func (this *Trie) SearchPrefix(prefix string) *Trie {
	node := this
	for _, ch := range prefix {
		ch -= 'a'
		if node.chidren[ch] == nil { // 找不到 nil
			return nil
		}
		node = node.chidren[ch] // 找到了向下继续
	}
	return node
}

func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	// 这个字符串在前缀树中不是前缀，而是单独的单词，所以 node.isEnd == true 才可以
	return node != nil && node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	// 是否存在前缀，不需要判定是不是完整单词
	return this.SearchPrefix(prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

