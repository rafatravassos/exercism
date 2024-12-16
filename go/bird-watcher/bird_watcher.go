package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, value := range birdsPerDay {
		total += value
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekStart := 0
	weekEnd := 6
	if week > 1 {
		weekEnd = weekEnd*week + 1
		weekStart = weekEnd - 6
	}
	total := 0
	for i := weekStart; i <= weekEnd; i++ {
		total += birdsPerDay[i]
	}
	return total

}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay
}
