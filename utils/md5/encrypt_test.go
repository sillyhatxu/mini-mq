package md5

import (
	"log"
	"testing"
)

func TestEncrypt(t *testing.T) {
	log.Println(Encrypt("Administrator"))
}
