def selectionSort(arr):
    for i in range(len(arr) - 1):
        # 记录最小数的索引
        minIndex = i
        for j in range(i + 1, len(arr)):
            if arr[j] < arr[minIndex]:
                minIndex = j
        # i 不是最小数时，将 i 和最小数进行交换
        if i != minIndex:
            arr[i], arr[minIndex] = arr[minIndex], arr[i]
    return arr


arr = [1, 5, 8, 9, 3, 5, 7, 2, 1]
selectionSort(arr)
print("排序之后的数组===>")
print(arr)


def quickSort(arr):
    pivot = arr[0]
    left, right = [], []
    for i in range(1, len(arr)):
        if arr[i] > pivot:
            right.append(arr[i])
        else:
            left.append(arr[i])
    
    return quickSort(left) + [pivot] + quickSort(right)
