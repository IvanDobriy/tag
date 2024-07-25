package tag

import (
	"testing"

	"github.com/IvanDobriy/tag/pkg/tag/cast"
	"github.com/stretchr/testify/assert"
)

var tagError *TagError

func TestGetRequiredTagPositive(t *testing.T) {
	assertions := assert.New(t)
	raw := map[string]any{
		"age": 12,
	}
	age, err := GetRequired("age", raw, func(value any) (*int64, error) {
		r, e := cast.ToInt64E(value)
		return &r, e
	})
	assertions.Equal(int64(12), *age)
	assertions.Nil(err)
}

func TestGetRequiredTagEntityNotFound(t *testing.T) {
	assertions := assert.New(t)
	raw := map[string]any{}
	age, err := GetRequired("age", raw, func(value any) (*int64, error) {
		r, e := cast.ToInt64E(value)
		return &r, e
	})
	assertions.Nil(age)
	assertions.NotNil(err)

	assertions.ErrorAs(err, &tagError)
}

func TestGetRequiredTagEntityConversionError(t *testing.T) {
	assertions := assert.New(t)
	raw := map[string]any{
		"age": "hello, world",
	}
	age, err := GetRequired("age", raw, func(value any) (*int64, error) {
		r, e := cast.ToInt64E(value)
		return &r, e
	})

	assertions.Nil(age)
	assertions.NotNil(err)
	assertions.ErrorAs(err, &tagError)
}

func TestGetTagPositve(t *testing.T) {
	assertions := assert.New(t)
	raw := map[string]any{
		"age": 12,
	}
	age, err := Get("age", raw, func(value any) (*int64, error) {
		r, e := cast.ToInt64E(value)
		return &r, e
	})
	assertions.Nil(err)
	assertions.Equal(int64(12), *age)
}

func TestGetTagEntityNotFound(t *testing.T) {
	assertions := assert.New(t)
	raw := map[string]any{}
	age, err := Get("age", raw, func(value any) (*int64, error){
		r, e := cast.ToInt64E(value)
		return &r, e
	})
	assertions.Nil(err)
	assertions.Nil(age)
}

func TestGetTagEntityConversionError(t *testing.T){
	assertions := assert.New(t)
	raw := map[string]any{
		"age": "hello, world",
	}
	age, err := Get("age", raw, func(value any) (*int64, error) {
		r, e := cast.ToInt64E(value)
		return &r, e
	})
	assertions.Nil(age)
	assertions.NotNil(err)
	assertions.ErrorAs(err, &tagError)
}
