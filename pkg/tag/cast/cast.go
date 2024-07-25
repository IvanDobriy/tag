package cast

import (
	"time"

	"github.com/spf13/cast"
)

func ToBoolE(el any) (bool, error) {
	return cast.ToBoolE(el)
}

func ToFloat64E(el any) (float64, error) {
	return cast.ToFloat64E(el)
}

func ToFloat32E(el any) (float32, error) {
	return cast.ToFloat32E(el)
}

func ToInt64E(el any) (int64, error) {
	return cast.ToInt64E(el)
}

func ToInt32E(el any) (int32, error) {
	return cast.ToInt32E(el)
}

func ToInt16E(el any) (int16, error) {
	return cast.ToInt16E(el)
}

func ToInt8E(el any) (int8, error) {
	return cast.ToInt8E(el)
}

func ToIntE(el any) (int, error) {
	return cast.ToIntE(el)
}

func ToUintE(el any) (uint, error) {
	return cast.ToUintE(el)
}

func ToUint64E(el any) (uint64, error) {
	return cast.ToUint64E(el)
}

func ToUint32E(el any) (uint32, error) {
	return cast.ToUint32E(el)
}

func ToUint16E(el any) (uint16, error) {
	return cast.ToUint16E(el)
}

func ToUint8E(el any) (uint8, error) {
	return cast.ToUint8E(el)
}

func ToStringE(el any) (string, error) {
	return cast.ToStringE(el)
}

func ToStringMapE(el any) (map[string]any, error) {
	return cast.ToStringMapE(el)
}

func ToSliceE(el any) ([]any, error) {
	return cast.ToSliceE(el)
}

func ToTimeE(el any) (time.Time, error) {
	return cast.ToTimeE(el)
}

func ToTimeInDefaultLocationE(el any, locatioin *time.Location) (time.Time, error) {
	return cast.ToTimeInDefaultLocationE(el, locatioin)
}

func ToDurationE(el any) (time.Duration, error) {
	return cast.ToDurationE(el)
}
