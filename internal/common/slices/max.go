package slices

func FindMax[S ~[]E, E any](x S, cmp func(a, b E) bool) E {
	if len(x) < 1 {
		panic("slices.MaxFunc: empty list")
	}
	m := x[0]
	for i := 1; i < len(x); i++ {
		if cmp(x[i], m) {
			m = x[i]
		}
	}
	return m
}
