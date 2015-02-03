package rot13

func Rot13(s string) string {
	p := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		b := s[i]
		if 'a' <= b && b <= 'm' || 'A' <= b && b <= 'M' {
			b += 13
		} else if 'n' <= b && b <= 'z' || 'N' <= b && b <= 'Z' {
			b -= 13
		}
		p[i] = b
	}
	return string(p)
}
