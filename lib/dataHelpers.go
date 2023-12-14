package lib

func PartitionByN[T any](items []T, numberOfChunks int) (chunks [][]T) {
	for i := 0; i < numberOfChunks; i++ {
		min := (i * len(items) / numberOfChunks)
		max := ((i + 1) * len(items)) / numberOfChunks
		chunks = append(chunks, items[min:max])
	}
	return chunks
}
