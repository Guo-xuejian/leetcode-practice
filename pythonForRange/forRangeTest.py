test_dict = {
    23: 12,
    45: 21,
    1:3,
    5:56,
    2:34,
    "guo":12,
    "fa":23,
}

for key, val in test_dict.items():
    print(key, val)


arr = []
arr.sort()
right = len(arr) - 1
curr = 65
while right >= 0:
    if arr[right] < curr:
        curr -= arr[right]