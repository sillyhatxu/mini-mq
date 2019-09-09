package uuid

import (
	"log"
	"testing"
)

func TestGeneratorUUID(t *testing.T) {
	log.Println(GeneratorUUID())
}
