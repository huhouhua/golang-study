package gin

import (
	"crypto/sha1"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestRouter(t *testing.T) {
	register()
}

func Test_GenPassword(t *testing.T) {
	s := sha1.New()
	_, err := io.WriteString(s, "123456"+"1")
	assert.NoError(t, err)

	password := fmt.Sprintf("%x", s.Sum(nil))

	assert.Equal(t, password, "b5cf498b70a176efeacbc5b07d88e0da76a7f4cb")

}
