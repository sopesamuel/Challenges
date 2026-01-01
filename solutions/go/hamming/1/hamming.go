package hamming
import "errors"
func Distance(a, b string) (int, error) {
    if len(a) != len(b){
        return 0, errors.New("Not equal length")
    }
    count := 0
	for i := range len(a) {
            if a[i] != b[i] {
                count += 1
            }
    }
    return count, nil
}
