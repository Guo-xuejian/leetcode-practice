def swap(a,b):
	print("中")
	print(id(a))
	a, b = b, a
	print("后")
	print(id(a))

def swapArray(arr):
	arr[0], arr[1] = arr[1], arr[0]
a = 1
b = 2
alist = [1,2]
print(id(a))
swap(a, b)
print("a====>", a)
print("b====>", b)

swapArray(alist)
print("first====>", alist[0])
print("second====>", alist[1])
