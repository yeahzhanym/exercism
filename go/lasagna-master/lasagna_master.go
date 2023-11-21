package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		return len(layers) * 2
	} else {
		return len(layers) * avgTime
	}
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		} else if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
//  Последний элемент в вашем списке ингредиентов всегда имеет значение "?". Функция должна заменить его последним элементом из вашего списка друзей.
func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portionCount int) []float64 {
	var newQuant []float64
	newQuant = append(newQuant, quantities...)

	for i := 0; i < len(quantities); i++ {
		newQuant[i] = newQuant[i] * float64(portionCount) / 2.0
	}

	return newQuant
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
