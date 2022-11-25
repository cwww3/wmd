package tools

import (
	"fmt"
	"testing"
)

func TestOpenDir(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		got, err := OpenDir("/Users/chenwei/project/wails-demo")
		if err != nil {
			t.Errorf("OpenDir() error = %v", err)
			return
		}
		fmt.Println(got)
	})
}
