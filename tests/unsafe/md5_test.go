package unsafe

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/url"
	"testing"
)

func TestMd5(t *testing.T) {
	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!")
	fmt.Printf("%x \n", h.Sum(nil))
}

func TestUrl(t *testing.T) {
	u, err := url.Parse("http://bing.com/search?q=dotnet")
	if err != nil {
		t.Fatal(err)
	}
	u.Scheme = "https"
	u.Host = "google.com"
	q := u.Query()
	q.Set("q", "golang")
	u.RawQuery = q.Encode()
	t.Log(u)
}
