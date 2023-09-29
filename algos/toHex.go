package algos

func ToHex(num int) string {
	if num == 0 {
		return "0"
	}

	n, ref := uint32(num), "0123456789abcdef"

	res := ""

	for n != 0 {
		res = string(ref[n&0xF]) + res
		n = n >> 4
	}

	return res
}
