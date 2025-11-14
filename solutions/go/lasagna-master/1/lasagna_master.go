package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTime int)(noLayers int){
    if prepTime == 0 {
         prepTime = 2
        return len(layers) * prepTime
    }
    return len(layers) * prepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
    for _, v := range layers {
        if v == "noodles"{
           noodles += 1 
        } else if v == "sauce" {
            sauce += 1
        }
        }
    return noodles * 50, sauce * 0.2
    }
    

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
    myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(slice []float64, numPortions int)[]float64{
    quant := []float64{}
    for _, v := range slice {
        newV := v * float64(numPortions)
        newV = newV / 2
        quant = append(quant, float64(newV))
    }
    return quant
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
