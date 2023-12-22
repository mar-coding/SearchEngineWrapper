package domain

type Bridger interface {
	None()
}

// Bridge you can get specific usecase or repository
// Example:
//
//	usecase := Bridge[domain.ExampleUseCase](config.EXAMPLE_COLLECTION, u.bridges)
func Bridge[T Bridger](collectionName string, bridges map[string]Bridger) T {
	return bridges[collectionName].(T)
}
