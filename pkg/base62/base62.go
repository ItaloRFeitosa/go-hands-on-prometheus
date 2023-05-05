package base62

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	base     = len(alphabet)
	slugLen  = 7
)

func Encode(numInt64 int64) string {
	num := int(numInt64)
	if num == 0 {
		return string(alphabet[0])
	}

	var encoded []byte
	for num > 0 {
		remainder := num % base
		num = num / base
		encoded = append([]byte{alphabet[remainder]}, encoded...)
	}

	paddingLen := slugLen - len(encoded)
	if paddingLen > 0 {
		zeroPadding := make([]byte, paddingLen)
		for i := range zeroPadding {
			zeroPadding[i] = alphabet[0]
		}
		encoded = append(zeroPadding, encoded...)
	}

	return string(encoded)
}

func Decode(encoded string) int64 {
	num := 0

	for i := 0; i < len(encoded); i++ {
		charIndex := -1
		for j := 0; j < base; j++ {
			if alphabet[j] == encoded[i] {
				charIndex = j
				break
			}
		}

		num = num*base + charIndex
	}

	return int64(num)
}
