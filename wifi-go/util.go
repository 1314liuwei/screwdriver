package wifi

func max(m, n int) int {
	if m > n {
		return m
	} else {
		return n
	}
}

func set(data []interface{}) []interface{} {
	set := map[interface{}]struct{}{}
	for i := 0; i < len(data); i++ {
		set[data[i]] = struct{}{}
	}

	var result []interface{}
	for k := range set {
		result = append(result, k)
	}
	return result

}
