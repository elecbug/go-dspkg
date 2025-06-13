package compare_type

// CompareType is an enumeration used for comparison results.
// It defines whether two values are equal, or if one is smaller or bigger than the other.
type CompareType int

const (
	EQUAL CompareType = iota
	TARGET_SMALL
	TARGET_BIG
)
