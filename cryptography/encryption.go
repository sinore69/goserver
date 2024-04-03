package cryptography

func Encrypt(text string, shift int) string {
	var encrypted string
    for _, char := range text {
        if char >= 'a' && char <= 'z' {
            encrypted += string((char-'a'+rune(shift))%26 + 'a')
        } else if char >= 'A' && char <= 'Z' {
            encrypted += string((char-'A'+rune(shift))%26 + 'A')
        } else {
            encrypted += string(char)
        }
    }
    return encrypted
}
