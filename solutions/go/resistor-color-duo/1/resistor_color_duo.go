package resistorcolorduo
import "strconv"
// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	encodedColors := map[string]int{
        "black" : 0,
        "brown" : 1,
        "red" : 2,
        "orange" : 3,
        "yellow" : 4,
        "green" : 5,
        "blue" : 6,
        "violet" : 7,
        "grey" : 8,
        "white" : 9,
    }
    result := ""
	result = strconv.Itoa(encodedColors[colors[0]]) + strconv.Itoa(encodedColors[colors[1]])
    final, _ := strconv.Atoi(result)
    return final
}
