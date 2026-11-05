package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
        return 0, errors.New("String length not equal")
    }
	hamming := 0
    for i := range a {
        if a[i] != b[i] {
            hamming += 1
        }
    }
    return hamming, nil
}
