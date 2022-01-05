/*
 * @lc app=leetcode.cn id=1400 lang=javascript
 *
 * [1400] 构造 K 个回文字符串
 */

// @lc code=start
/**
 * @param {string} s
 * @param {number} k
 * @return {boolean}
 */
var canConstruct = function (s, k) {
    const max = s.length;
    let min = 0;
    for (const c of s) {
        m[c] = (m[c] || 0) + 1;
    }
    for (const v of Object.values(m)) {
        if (v & 1) min++;
    }
    return k >= min && k <= max;
};
// @lc code=end

