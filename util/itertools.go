package util

func NextProduct(a []int, r int) func() []int {
	p := make([]int, r)
	x := make([]int, len(p))
	return func() []int {
		p := p[:len(x)]
		for i, xi := range x {
			p[i] = a[xi]
		}
		for i := len(x) - 1; i >= 0; i-- {
			x[i]++
			if x[i] < len(a) {
				break
			}
			x[i] = 0
			if i <= 0 {
				x = x[0:0]
				break
			}
		}
		return p
	}
}

func ProductGen(a []int, repeat int, skip bool) [][]int {
	npt := NextProduct(a, repeat)
	vector := make([][]int, 0)

	for {
		np := npt()
		if len(np) == 0 {
			break
		}
		// skip the central cube [0,0,0]
		/*
			if np[0] == 0 && np[1] == 0 && np[2] == 0 {
				continue
			}
		*/
		flag := true
		for _, v := range np {
			if v != 0 {
				flag = false
				break
			}
		}
		if flag && skip {
			continue
		}

		nn := make([]int, len(np))
		copy(nn, np)
		vector = append(vector, nn)
	}

	return vector
}
