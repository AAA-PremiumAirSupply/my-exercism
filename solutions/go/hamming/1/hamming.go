package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("len not eq")
	}
	ilen := len(a)
	count := 0
	for i := 0; i < ilen; i++ {
		if a[i] != b[i] {
			count += 1
		}
	}
	return count, nil
}
