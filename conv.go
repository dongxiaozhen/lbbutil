package lbbutil

import "strconv"

func ToUint32(data string) uint32 {
	a, err := strconv.ParseUint(data, 10, 32)
	if err != nil {
		return 0
	}
	return uint32(a)
}
func ToUint64(data string) uint64 {
	a, err := strconv.ParseUint(data, 10, 64)
	if err != nil {
		return 0
	}
	return a
}

func ToInt64(data string) int64 {
	a, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return -1
	}
	return a
}
func Int64ToString(value int64) string {
	a := strconv.FormatInt(value, 64)
	return a
}
