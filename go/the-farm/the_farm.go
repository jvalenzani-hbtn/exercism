package thefarm

import (
    "errors"
    "fmt"
)
// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amount, err := weightFodder.FodderAmount()
    if err != nil{
        if err == ErrScaleMalfunction{
            if amount >= 0{
            	amount *= 2.0
        	} else {
            	return 0.0, errors.New("negative fodder")
            } 
        } else {
        	return 0.0, err
        }
    }
	if amount < 0 {
        return 0.0, errors.New("negative fodder")
    }
	if cows == 0{
        return 0.0, errors.New("division by zero")
    }
	if cows < 0{
        return 0.0, errors.New(fmt.Sprintf("silly nephew, there cannot be %d cows", cows))
    }
	return float64(amount)/float64(cows), nil
}
