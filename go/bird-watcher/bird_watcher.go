package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	res := 0
	for _, bird := range birdsPerDay {
		res += bird
	}

	return res
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekIndex := week * 7
	res := 0
	for i:= weekIndex - 7; i < weekIndex; i++ {
		res += birdsPerDay[i]
	}

	return res
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	birdsPerDay[0] += 1
	for i:=1; i < len(birdsPerDay); i++ {
		if i % 2 == 0{
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay
}
