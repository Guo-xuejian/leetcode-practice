/*
 * @lc app=leetcode.cn id=535 lang=golang
 *
 * [535] TinyURL 的加密与解密
 */

// @lc code=start
type Codec struct {
	dataBase map[int]string
	id       int
}

func Constructor() Codec {
	return Codec{map[int]string{}, 0}
}

// Encodes a URL to a shortened URL.
func (c *Codec) encode(longUrl string) string {
	c.id++
	c.dataBase[c.id] = longUrl
	return "http://tinyurl.com/" + strconv.Itoa(c.id)
}

// Decodes a shortened URL to its original URL.
func (c *Codec) decode(shortUrl string) string {
	i := strings.LastIndexByte(shortUrl, '/')
	id, _ := strconv.Atoi(shortUrl[i+1:])
	return c.dataBase[id]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
// @lc code=end

