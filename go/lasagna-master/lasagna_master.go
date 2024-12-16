package lasagna

func PreparationTime(layers []string, minutes int) int {
	if minutes == 0 {
		minutes = 2
	}
	return minutes * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauces := 0.0
	for _, item := range layers {
		if item == "noodles" {
			noodles++
		} else if item == "sauce" {
			sauces++
		}
	}
	return (noodles * 50), (sauces * 0.2)

}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	x := make([]float64, len(quantities))
	copy(x, quantities)
	for i, item := range x {
		x[i] = item * float64(portions) / 2
	}
	return x

}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
