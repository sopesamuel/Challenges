package etl
import "strings"
func Transform(in map[int][]string) map[string]int {
	final := map[string]int{}
    for point, sliceString := range in {
        for _, v := range sliceString {
            final[strings.ToLower(v)] = point
        }
    }
    return final
}
