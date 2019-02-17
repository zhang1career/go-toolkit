package security

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"github.com/zhang1career/lib/log"
	"io"
	"os"
)

func FileChecksum(filename string) []byte {
	// Open file for reading
	file, err := os.Open(filename)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	defer file.Close()
	// Create new hasher, which is a writer interface
	hasher := md5.New()
	// Default buffer size for copying is 32*1024 or 32kb per copy
	// Use io.CopyBuffer() if you want to specify the buffer to use
	// It will write 32kb at a time to the digest/hash until EOF
	// The hasher implements a Write() function making it satisfy
	// the writer interface. The Write() function performs the digest
	// at the time the data is copied/written to it. It digests
	// and processes the hash one chunk at a time as it is received.
	_, err = io.Copy(hasher, file)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	// Now get the final sum or checksum.
	// We pass nil to the Sum() function because
	// we already copied the bytes via the Copy to the
	// writer interface and don't need to pass any new bytes
	return hasher.Sum(nil)
}


// secretKey should be unique, protected, private,
// and not hard-coded like this. Store in environment var
// or in a secure configuration file.
// This is an arbitrary key that should only be used
// for example purposes.
var secretKey = "neictr98y85klfgneghre"

// Create a salt string with 32 bytes of crypto/rand data
func GenerateSalt() string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(randomBytes)
}

func PasswordHash(plainText string, salt string) string {
	hash := hmac.New(sha256.New, []byte(secretKey))
	io.WriteString(hash, plainText + salt)
	hashedValue := hash.Sum(nil)
	return hex.EncodeToString(hashedValue)
}