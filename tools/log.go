package tools

import (
	"fmt"
	"github.com/cwww3/wmd/global"
)

func Logf(str string, args ...any) {
	global.Logger.Printf(fmt.Sprintf("%s\n", fmt.Sprintf(str, args)))
}

func Log(args ...any) {
	global.Logger.Print(args...)
}
