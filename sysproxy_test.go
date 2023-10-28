package sysproxy

import (
	"os"
	"path"
	"testing"
    "fmt"

	"github.com/stretchr/testify/assert"
)

func TestAllEquals(t *testing.T) {
	assert.True(t, allEquals("127.0.0.1:8888", "127.0.0.1:8888\n127.0.0.1:8888"))
	assert.True(t, allEquals("127.0.0.1:8888", "127.0.0.1:8888\n127.0.0.1:8888\n"))
	assert.False(t, allEquals("127.0.0.1:8888", "127.0.0.1:8888\n127.0.0.1:8887"))
	assert.True(t, allEquals("", "\n\n"))
	assert.True(t, allEquals("", "\r\n"))
	assert.False(t, allEquals("", "127.0.0.1:8888"))
	assert.False(t, allEquals("127.0.0.1:8888", ""))
}

func TestGetOutput(t *testing.T) {
	path := path.Join(os.TempDir(), "sysproxy")
	err := EnsureHelperToolPresent(path, "For test purpose", "")
	assert.NoError(t, err, "should install helper tool")
    show, err := Show()
    fmt.Println("before On:", show)
	off, err := On("localhost:8888")
	assert.NoError(t, err, "should set system proxy on")
    show, err = Show()
    fmt.Println("after On:", show)
	err = off()
	assert.NoError(t, err, "should set system proxy off")
    show, err = Show()
    fmt.Println("after Off:", show)
}
