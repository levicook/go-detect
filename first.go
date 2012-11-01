package detect

const zeroString = ""

func String(strings ...string) string {
	for _, s := range strings {
		if s != zeroString {
			return s
		}
	}
	return zeroString
}
