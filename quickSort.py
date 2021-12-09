def quickSort(arr):
    if len(arr) < 2:
        return arr
    pivot = arr[0]
    left = right = []

    for one in arr[1:]:
        if one <= pivot:
            print("left in", one)
            left.append(one)
        else:
            print("right in", one)
            right.append(one)
    print(left)
    print(right)
    return quickSort(left) + [pivot] + quickSort(right)


arr = [2, 1, 8]
arr = quickSort(arr)
print(arr)
