package tasks

// Дан массив высот длинной n. Это n вертикальных линий, идущих друг за другом.
// Необходимо найти 2 линии, образующие контейнер в котором больше всего воды
func MaxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}

	var (
		l = 0
		r = len(height) - 1

		maxV = 0
	)

	for l < r {
		var smallestHeight int

		if height[l] > height[r] {
			smallestHeight = height[r]
			r--
		} else {
			smallestHeight = height[l]
			l++
		}

		curr := (r - l + 1) * smallestHeight

		if curr > maxV {
			maxV = curr
		}
	}

	return maxV
}
