package slices

func GroupBy[T any](values []T, groupator func(T) (string, T)) [][]T {
	var res [][]T = make([][]T, 0)

	temp := make(map[string][]T)
	for _, v := range values {
		key, val := groupator(v)
		if _, hasVal := temp[key]; !hasVal {
			temp[key] = []T{val}
		} else {
			temp[key] = append(temp[key], val)
		}
	}
	for _, v := range temp {
		res = append(res, v)
	}

	return res
}
