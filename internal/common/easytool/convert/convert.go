package convert

import (
	"fmt"
	"math"
	"strconv"
)

// ToStr convert any to string
func ToStr(a any) string {
	return fmt.Sprintf("%v", a)
}

// ToInt64 convert any to int64
func ToInt64(a any) int64 {
	switch v := a.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case string:
		num, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			return num
		}
	default:
		return 0
	}
	return 0
}

// ToUint64 convert any to uint64
func ToUint64(a any) uint64 {
	switch v := a.(type) {
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return v
	case int:
		if v >= 0 {
			return uint64(v)
		}
	case int8:
		if v >= 0 {
			return uint64(v)
		}
	case int16:
		if v >= 0 {
			return uint64(v)
		}
	case int32:
		if v >= 0 {
			return uint64(v)
		}
	case int64:
		if v >= 0 {
			return uint64(v)
		}
	case float32:
		if v >= 0 {
			return uint64(v)
		}
	case float64:
		if v >= 0 {
			return uint64(v)
		}
	case string:
		num, err := strconv.ParseUint(v, 10, 64)
		if err == nil {
			return num
		}
	default:
		return 0
	}
	return 0
}

// ToInt32 convert any to int32
func ToInt32(a any) int32 {
	switch v := a.(type) {
	case int:
		return int32(v)
	case int8:
		return int32(v)
	case int16:
		return int32(v)
	case int32:
		return v
	case int64:
		if v >= math.MinInt32 && v <= math.MaxInt32 {
			return int32(v)
		}
	case float32:
		return int32(v)
	case float64:
		if float64(int32(v)) == v {
			return int32(v)
		}
	case string:
		num, err := strconv.ParseInt(v, 10, 32)
		if err == nil {
			return int32(num)
		}
	default:
		return 0
	}
	return 0
}

// ToUInt32 convert any to int32
func ToUInt32(a any) uint32 {
	switch v := a.(type) {
	case uint:
		return uint32(v)
	case uint8:
		return uint32(v)
	case uint16:
		return uint32(v)
	case uint32:
		return v
	case int:
		if v >= 0 {
			return uint32(v)
		}
	case int8:
		if v >= 0 {
			return uint32(v)
		}
	case int16:
		if v >= 0 {
			return uint32(v)
		}
	case int32:
		if v >= 0 {
			return uint32(v)
		}
	case int64:
		if v >= 0 {
			return uint32(v)
		}
	case float32:
		if v >= 0 {
			return uint32(v)
		}
	case float64:
		if v >= 0 {
			return uint32(v)
		}
	case string:
		num, err := strconv.ParseUint(v, 10, 32)
		if err == nil {
			return uint32(num)
		}
	default:
		return 0
	}
	return 0
}
