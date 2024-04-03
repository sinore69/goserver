package cryptography

func Decrypt(text string, shift int) string {
    return Encrypt(text, 26-shift)
}