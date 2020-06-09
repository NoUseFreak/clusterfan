package clusterfan

type store struct {
	data   map[string][]int
	maxLen int
}

func newStore() store {
	return store{
		data:   map[string][]int{},
		maxLen: 5,
	}
}

func (r *store) Add(origin string, temp int) {
	if _, ok := r.data[origin]; !ok {
		// New origin
		r.data[origin] = []int{}
	}

	if len(r.data[origin]) > r.maxLen {
		r.data[origin] = r.data[origin][len(r.data[origin])-r.maxLen:]
	}

	r.data[origin] = append(r.data[origin], temp)
}

func (r *store) Max() int {
	max := 0

	for _, o := range r.data {
		if m := maxInSlice(o); m > max {
			max = m
		}
	}

	return max
}

func (r *store) Stats() map[string]int {
	m := map[string]int{}

	for o, v := range r.data {
		m[o] = v[0]
	}

	return m
}
