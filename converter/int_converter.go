package converter

import "strconv"

func UintToString(num uint) string {
	return strconv.FormatUint(uint64(num), 10)
}

func IntToString(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}
