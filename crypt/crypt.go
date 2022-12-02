//go:build linux

package crypt

/*
#cgo LDFLAGS: -lcrypt
#include <unistd.h>

char* crypto(char *word, char *salt) {
   return crypt(word,salt);
}
*/
import "C"

func Crypt(word, salt string) string {
	// word: Encrypted plaintext
	// salt: Salt
	// It can support salt in the form of ${ident} ${salt} -> ${ident} ${salt}${passwd}.
	// ident represents the encryption method and supports the following encryption:
	// ident encryption:
	// 0     crypt
	// 1     md5
	// 5     sha256
	// 6     sha512

	wordc := C.CString(word)
	saltc := C.CString(salt)
	pwdc := C.crypto(wordc, saltc)
	pwd := C.GoString(pwdc)
	return pwd
}
