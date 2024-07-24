package tag

import (
	"errors"
	"fmt"
)

type TagError struct {
	E error

}

func (t *TagError) Error() string {
	return t.E.Error()
}

func (t* TagError) Unwrap() error {
	return t.E
}

func NewTagError(msg string) error{
	return &TagError{E: errors.New(msg)}
}

func Get[T any](name string, raw map[string]any, callback func(value any) (T, error)) (T, error) {
	var result T
	data := raw[name]
	if data == nil {
		return result, nil
	}
	result, err := callback(data)
	if err != nil {
		tagError := NewTagError(fmt.Sprintf("tag: \"%s\" conversion error", name))
		return result, fmt.Errorf("%w %w", tagError, err)
	}
	return result, nil
}

func GetRequired[T any](name string, raw map[string]any, callback func(value any) (T, error)) (T, error) {
	var result T
	data := raw[name]
	if data == nil {
		e := NewTagError(fmt.Sprintf("required tag: \"%s\" not found", name))
		return result, fmt.Errorf("%w", e)
	}
	result, err := callback(data)
	if err != nil {
		tagError := NewTagError(fmt.Sprintf("tag: \"%s\" conversion error", name))
		return result, fmt.Errorf("%w %w", tagError, err)
	}
	return result, nil
}
