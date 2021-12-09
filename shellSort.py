def shellSort(arr):
    length = len(arr)
    gap = length // 2       # 二分之一为步长
    while gap >= 1:
        for i in range(gap, length):
            j = i
            while j >= gap and arr[j-gap] > arr[j]:     # 执行插入排序  为什么 j >= gap
                arr[j], arr[j-gap] = arr[j-gap], arr[j]
                j -= gap
        gap = gap // 2      # 更新步长直到为 1
    return arr


arr = [1, 5, 8, 9, 3, 5, 7, 2, 1]
arr = shellSort(arr)
print("排序之后的数组===>")
print(arr)
