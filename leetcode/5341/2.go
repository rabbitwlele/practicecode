package _341

type ProductOfNumbers struct {
	a []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{[]int{}}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.a = []int{}
	}
	if len(this.a) == 0 {
		this.a = append(this.a, num)
	} else {
		this.a = append(this.a, this.a[len(this.a)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if len(this.a) < k {
		return 0
	}
	if len(this.a) == k {
		return this.a[len(this.a)-1]
	}
	return this.a[len(this.a)-1] / this.a[len(this.a)-k-1]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
