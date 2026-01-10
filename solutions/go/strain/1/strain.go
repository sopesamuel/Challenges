package strain

// Implement the "Keep" and "Discard" function in this file.

func Keep[T any](t []T, predicate func(T) bool) []T{
	even := []T{}
    for _, v := range t {
        if predicate(v) == true {
            even = append(even, v)
        }
    }
    return even
}

func Discard[T any](t []T, predicate func(T) bool) []T{
	odd := []T{}
    for _, v := range t {
        if predicate(v) == false {
            odd = append(odd, v)
        }
    }
    return odd
}
// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
