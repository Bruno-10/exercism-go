package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
    if timePerLayer == 0 {
        return 2 * len(layers)
    }

   return timePerLayer * len(layers) 
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodles += 50
        }
        if layers[i] == "sauce" {
            sauce += 0.2
        }
    }

    return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(list1[]string, list2[]string) {
   list2[len(list2) - 1] = list1[len(list1) - 1]
}


// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    returnValue := make([]float64, len(quantities), len(quantities))
    for i:= 0; i < len(quantities); i++ {
        returnValue[i] = (quantities[i] * float64(portions)) / 2
    }    
    return returnValue
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
