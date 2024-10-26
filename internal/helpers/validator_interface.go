package helpers

type Validator[T any] interface {
	Validate(input *T) error
}
