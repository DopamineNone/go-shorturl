package sequence

// Sequence generate sequence
type Sequence interface {
	Next() (uint64, error)
}