package util

func ForAllPairs[T any](slice []T, do func(a, b T)) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			do(slice[i], slice[j])
		}
	}
}
