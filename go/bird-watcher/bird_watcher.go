package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var sum int = 0
    for _,value := range birdsPerDay{
        sum += value
    }
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    var sum int = 0
    for _, v := range birdsPerDay[7*(week-1):(7*week)]{
        sum += v
    }
	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i < len(birdsPerDay); i+=2{
        birdsPerDay[i]++ 
    }
	return birdsPerDay
}
