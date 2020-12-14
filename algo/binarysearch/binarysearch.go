package binarysearch

// Returns if the value is present or not and the index where value is present. If needle is not present then it returns false, -ve of index where it should be present.
func BinarySearch(needle int, haystack []int) (bool, int) {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false, -low
	}

	return true, low
}
