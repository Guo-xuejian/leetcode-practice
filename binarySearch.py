def binarySearch(arr, target):
	low = 0
	high = len(arr) - 1
	
	step = 0
	while True:
		step += 1
		if low <= high:
			mid = (low + high) / 2
			guess = arr[mid]
			if guess == target:
				print(f"共查找了{}词\n", step)
				return mid
			if guess > target:
				high = mid - 1
			else:
				low = mid + 1
		else:
			break
