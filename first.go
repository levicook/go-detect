package detect

import "time"

const zeroString = ""

func String(strings ...string) string {
	for _, s := range strings {
		if s != zeroString {
			return s
		}
	}
	return zeroString
}

func Time(times ...time.Time) time.Time {
	for _, t := range times {
		if !t.IsZero() {
			return t
		}
	}
	return time.Time{}
}
