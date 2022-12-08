package qiniu

import (
	"testing"
)

func TestUploadBytes(t *testing.T) {
	InitOSS(
		"bucket",
		"ak",
		"sk",
		7200,
	)
	data := []byte("hello world")
	UploadBytes("test", data)
}
