package cast

import (
	"github.com/spf13/cast"
)

func ToBoolE(el any) (bool, error) {
	return cast.ToBoolE(el)
}

func ToFloat64E(el any) (float64, error) {
	return cast.ToFloat64E(el)
}

func ToInt64E(el any) (int64, error) {
	return cast.ToInt64E(el)
}

