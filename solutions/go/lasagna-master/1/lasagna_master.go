package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	rTime := 0
	if time == 0 {
		rTime = 2
	} else {
		rTime = time
	}
	return len(layers) * rTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	n := 0
	s := 0
	for _, v := range layers {
		if v == "noodles" {
			n++
		}
		if v == "sauce" {
			s++
		}
	}
	return n * 50, float64(s) * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(frd []string, mine []string) {
	mine[len(mine)-1] = frd[len(frd)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(s []float64, t int) []float64 {
	newS := append([]float64{}, s...)
	for idx, v := range newS {
		newS[idx] = v * float64(t) / 2
	}
	return newS
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
