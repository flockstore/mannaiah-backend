package util

// SafeDeref returns the value of the pointer or the fallback if it's nil.
func SafeDeref[T any](ptr *T, fallback T) T {
	if ptr != nil {
		return *ptr
	}
	return fallback
}
