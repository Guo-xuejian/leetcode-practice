// 剑指 Offer 40. 最小的k个数
// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

 

// 示例 1：

// 输入：arr = [3,2,1], k = 2
// 输出：[1,2] 或者 [2,1]
// 示例 2：

// 输入：arr = [0,1,2,1], k = 1
// 输出：[0]
 

// 限制：

// 0 <= k <= arr.length <= 10000
// 0 <= arr[i] <= 10000
// 通过次数322,305提交次数562,46

func getLeastNumbers(arr []int, k int) (res []int) {
    h := make(hp, len(arr))
    for idx, num := range arr {
        h[idx] = num
    }
    heap.Init(&h)
    for i := 0; i < k; i++ {
        res = append(res, (heap.Pop(&h)).(int))
    }
    return
}

type hp []int

func (h hp) Len() int { return len(h) }

func (h hp) Less(i, j int) bool { return h[i] < h[j] }

func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(num interface{}) { *h = append(*h, num.(int)) }

func (h *hp) Pop() (res interface{}) {
    *h, res = (*h)[:h.Len()-1], (*h)[h.Len()-1]
    return
}