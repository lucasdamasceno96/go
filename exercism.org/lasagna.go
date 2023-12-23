package lasagna

// TODO: define the 'OvenTime' constant

const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	var remaining int
	remaining = OvenTime - actualMinutesInOven
	return remaining
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	var layer int
	layer = 2 * numberOfLayers
	return layer
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	var elapsed int
	var layer = numberOfLayers * 2
	elapsed = layer + actualMinutesInOven
	return elapsed
}
