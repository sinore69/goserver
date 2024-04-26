package cryptography

import "encoding/base64"

func Encode(input string) (string,error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	return encoded,nil
}
