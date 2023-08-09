package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
    cows int
    message string
}

func (ice *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", ice.cows, ice.message)
}
// TODO: define the 'DivideFood' function
func DivideFood(foc FodderCalculator, cows int) (float64, error){
    totalAmount, err := foc.FodderAmount(cows)
    if err != nil {
       return 0, err 
    }
    fatteningFactor, errFf := foc.FatteningFactor()
    if errFf != nil {
        return 0, errFf
    }

    return (totalAmount * fatteningFactor) / float64(cows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(foc FodderCalculator, cows int) (float64, error){
    if cows > 0 {
        return DivideFood(foc, cows)
    }

    return 0, errors.New("invalid number of cows")
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
    if cows == 0 {
        return &InvalidCowsError{cows: cows, message: "no cows don't need food"}
    }
    if cows < 0 {
        return &InvalidCowsError{cows: cows, message: "there are no negative cows"}
    }
    return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
