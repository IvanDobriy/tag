package tag

import (
	"errors"
	"testing"

	"github.com/IvanDobriy/tag/pkg/tag/cast"
	"github.com/stretchr/testify/assert"
)


func TestGetRequiredTagPositive(t * testing.T) {
	assertions := assert.New(t)
	raw := map[string]any {
		"age": 12,
	}
	age, err := GetRequired("age", raw, func(value any) (*int64, error) {
		r, e := cast.ToInt64E(value)
		return &r, e
	})
	assertions.Equal(int64(12), *age)
	assertions.Nil(err)
}

func TestGetRequiredTageEntityNotFound(t *testing.T){
	assertions := assert.New(t)
	raw := map[string]any {}
	age, err := GetRequired("age", raw, func(value any) (*int64, error) {
		r, e := cast.ToInt64E(value)
		return &r, e
	})
	assertions.Nil(age)
	assertions.NotNil(err)

	var tagError *TagError
	r := errors.As(tagError, err)
	assertions.True(r, "must be true ---- ")
	assertions.ErrorAs(tagError, err)
}