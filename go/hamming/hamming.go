package hamming

import "errors"
func Distance(a, b string) (int, error) {
    if (len(a) != len(b)){
    	return 0, errors.New("Sequences lenght don't match")
    }
	var diff int = 0
	for i := 0; i < len(a); i++{
        if a[i] != b[i]{
            diff++
        }
    }
	return diff, nil
}
