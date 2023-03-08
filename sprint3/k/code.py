def merge(arr: list, left: int, mid: int, right: int) -> list:
	res = []
	left = arr[left:mid]
	right = arr[mid:right]
	mod_left = i_right = 0
	while len(res) < len(left) + len(right):
		if left[mod_left] <= right[i_right]:
			res.append(left[mod_left])
			mod_left += 1
		else:
			res.append(right[i_right])
			i_right += 1
		if i_right == len(right):
			res += left[mod_left:]
			break
		if mod_left == len(left):
			res += right[i_right:]
			break
	return res


def merge_sort(arr: list, left: int, right: int) -> list:
	if right - left >= 2:
		mid = (left + right) // 2
		merge_sort(arr, left, mid)
		merge_sort(arr, mid, right)
		arr[left:right] = merge(arr, left, mid, right)


def test():
	a = [1, 4, 9, 2, 10, 11]
	b = merge(a, 0, 3, 6)
	expected = [1, 2, 4, 9, 10, 11]
	assert b == expected
	c = [1, 4, 2, 10, 1, 2]
	merge_sort(c, 0, 6)
	expected = [1, 1, 2, 2, 4, 10]
	assert c == expected


if __name__ == '__main__':
	test()