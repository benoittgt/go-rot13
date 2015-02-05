package rot13

func Rot13(message string) string {
	newMessage := make([]byte, len(message))
	for i := 0; i < len(message); i++ {
		letterBytes := message[i]
		letterBytes += uint8(13 * letterDirection(letterBytes))
		newMessage[i] = letterBytes
	}
	return string(newMessage)
}

func letterIn(letterBytes byte, a byte, b byte, c byte, d byte) bool {
	return a <= letterBytes && letterBytes <= b || c <= letterBytes && letterBytes <= d
}

func letterDirection(letterBytes byte) (direction int) {

	if letterIn(letterBytes, 'a', 'm', 'A', 'M') {
		return 1
	} else if letterIn(letterBytes, 'n', 'z', 'N', 'Z') {
		return -1
	}
	return 0
}
