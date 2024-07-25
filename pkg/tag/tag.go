package tag

import (
	"fmt"
)

type TagError struct {
	Msg string
}

func (t *TagError) Error() string {
	return t.Msg
}

func NewTagError(msg string) error {
	return &TagError{Msg: msg}
}

func Get[T any](name string, raw map[string]any, callback func(value any) (T, error)) (T, error) {
	var result T
	data := raw[name]
	if data == nil {
		return result, nil
	}
	r, err := callback(data)
	if err != nil {
		tagError := NewTagError(fmt.Sprintf("tag: \"%s\" conversion error", name))
		return result, fmt.Errorf("%w %w", tagError, err)
	}
	result = r
	return result, nil
}

func GetRequired[T any](name string, raw map[string]any, callback func(value any) (T, error)) (T, error) {
	var result T
	data := raw[name]
	if data == nil {
		e := NewTagError(fmt.Sprintf("required tag: \"%s\" not found", name))
		return result, fmt.Errorf("%w", e)
	}
	r, err := callback(data)
	if err != nil {
		tagError := NewTagError(fmt.Sprintf("tag: \"%s\" conversion error", name))
		return result, fmt.Errorf("%w %w", tagError, err)
	}
	result = r
	return result, nil
}
