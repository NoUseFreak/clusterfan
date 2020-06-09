package clusterfan

import "os"

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func maxInSlice(slice []int) int {
	max := 0
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}
