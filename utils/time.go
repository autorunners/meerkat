package utils

import (
	"time"
)

func GetTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Second)
}

func GetTimestampMilli() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetTimestampMicro() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

func GetTimestampNano() int64 {
	return time.Now().UnixNano()
}
