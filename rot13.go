package rot13

import "strings"

func Rot13(toEncrypt string) (crypted string) {
	alphab := "abcdefghijklmnopqrstuvwxyz"
	alphabCap := strings.ToUpper(alphab)
	alphabet := strings.Split(alphab, "")
	alphabetCap := strings.Split(alphabCap, "")

	lettersToCrypt := strings.Split(toEncrypt, "")

	for _, letterToCrypt := range lettersToCrypt {
		if strings.Index(alphabCap, letterToCrypt) >= 0 {
			crypted += alphabetCap[(strings.Index(alphabCap, letterToCrypt)+13)%26]
		} else if strings.Index(alphab, letterToCrypt) >= 0 {
			crypted += alphabet[(strings.Index(alphab, letterToCrypt)+13)%26]
		} else {
			crypted += letterToCrypt
		}
	}
	return crypted

}
