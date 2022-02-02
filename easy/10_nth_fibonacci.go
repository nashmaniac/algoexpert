package easy

func GetNthFib(n int) int {
	arr := make([]int, n)
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	arr[0] = 0
	arr[1] = 1
	for i:= 2; i<n;i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n-1]
}