package lasagna

// PreparationTime function
func PreparationTime(layers []string, timePerLayer int) int{
   if timePerLayer <= 0 {
       timePerLayer = 2
   } 
	return timePerLayer * len(layers)
}

// Quantities function
func Quantities(layers []string) (noodles int, sauce float64){
    for _, v := range layers{
        if (v == "sauce"){
           sauce += 0.2
            }
        if (v == "noodles"){
            noodles += 50
        }
	}
	return 
}

// AddSecretIngredient function
func AddSecretIngredient(friendsList []string, myList []string){
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}
// ScaleRecipe function
func ScaleRecipe(quantities []float64, portions int) []float64{
    var factor float64 = float64(portions)/2.0
    var scaledQuantities []float64 = make([]float64, len(quantities))
    for i := 0; i < len(quantities); i++{
        scaledQuantities[i] = quantities[i]*factor
    }
	return scaledQuantities
}
