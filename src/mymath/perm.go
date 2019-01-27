package mymath

type PermuStategy interface {
	permutate() [][]interface{}
}

type OrdiPermutate struct {
	V []interface{}
}

func (p *OrdiPermutate) Permutate() [][]interface{} {
	ans := make([][]interface{}, 0)
	p.permutate(&ans, p.V, 0)

	return ans
}

func (p *OrdiPermutate) permutate(ans *[][]interface{}, v []interface{}, start int) {
	if start >= len(v) {
		tmp := make([]interface{}, len(v))
		copy(tmp, v)
		*ans = append(*ans, tmp)
		return
	}

	for i := start; i < len(v); i++ {
		v[start], v[i] = v[i], v[start]
		p.permutate(ans, v, start + 1)
		v[start], v[i] = v[i], v[start]
	}
}

