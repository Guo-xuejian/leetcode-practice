def bubbleSort(arr):
    length = len(arr)

    # 遍历所有的数组元素
    for i in range(length):
        # 和后续元素作比较
        # for j in range(length-1):
        #     # for 循环内不断交换元素，将较大值向列表尾处移动
        #     if arr[j] > arr[j+1]:
        #         arr[j], arr[j+1] = arr[j+1], arr[j]
        for j in range(0, length-i-1):
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]


arr = [64, 34, 25, 12, 12,22, 11, 90]

bubbleSort(arr)

print("排序后的数组:")
for i in range(len(arr)):
    print("%d" % arr[i])


