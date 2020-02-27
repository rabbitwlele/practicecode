package  main

func reverse(x int) int {
	fu := false
	if x < 0 {
		x = -x
		fu = true
	}
	ans := 0
	for x > 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	if fu {
		ans = -ans
	}
	if ans < -(1<<31) || ans > (1<<31)-1 {
		return 0
	}
	return ans
}
