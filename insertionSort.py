# 插入排序（英语：Insertion Sort）是一种简单直观的排序算法。它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。

def insertionSort(arr):
    for i in range(1, len(arr)):    # 从 1 开始保留了至少上一位的存在
        temp = arr[i]
        print(arr[i])
        j = i - 1
        while j >= 0 and temp < arr[j]:
            arr[j+1] = arr[j]
            j -= 1
        arr[j+1] = temp


arr = [1,5,8,9,3,5,7,2,1]
insertionSort(arr)
print("排序之后的数组===>")
print(arr)