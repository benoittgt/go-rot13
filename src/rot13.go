package rot13

import "strings"

func Rot13(toCrypt string) (crypted string) {
	alphab := "abcdefghijklmnopqrstuvwxyz"
	alphabMaj := strings.ToUpper(alphab)
	alphabet := strings.Split(alphab, "")
	alphabetMaj := strings.Split(alphabMaj, "")

	lettersToCrypt := strings.Split(toCrypt, "")

	for _, letterToCrypt := range lettersToCrypt {
		if strings.Index(alphabMaj, letterToCrypt) >= 0 {
			crypted += alphabetMaj[(strings.Index(alphabMaj, letterToCrypt)+13)%26]
		} else if strings.Index(alphab, letterToCrypt) >= 0 {
			crypted += alphabet[(strings.Index(alphab, letterToCrypt)+13)%26]
		} else {
			crypted += letterToCrypt
		}
	}
	return crypted

}
