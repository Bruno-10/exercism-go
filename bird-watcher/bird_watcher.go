package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var result int;
    for i := 0; i < len(birdsPerDay); i++ {
        result += birdsPerDay[i]
    }
    return result
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    var result int;
    endDay := week * 7
    startDay := endDay - 7

    for i := startDay; i < endDay; i++ {
        result += birdsPerDay[i]
    }
    return result
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i < len(birdsPerDay); i = i + 2 {
        birdsPerDay[i] += 1
    }
    return birdsPerDay
}
