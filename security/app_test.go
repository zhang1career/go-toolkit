package security_test

import (
	"fmt"
	"github.com/zhang1career/lib/security"
	"testing"
)

func TestFileMd5(t *testing.T) {
	checksum := security.FileChecksum("/Users/zhang/Downloads/50_business_model_examples.pdf")
	fmt.Printf("Md5 checksum: %x\n", checksum)
}

func TestPasswordHash(t *testing.T) {
	salt := security.GenerateSalt()
	hashed := security.PasswordHash("qwertyui", salt)
	fmt.Println(hashed)
}