package main

import (
	"changeme/config"
	"changeme/global"
	"changeme/tools"
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	goruntime "runtime"
)

//go:embed frontend/dist
var assets embed.FS

func init() {
	var err error
	config.UserHomeDir, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	workspace := filepath.Join(config.UserHomeDir, config.AppDir)
	err = os.MkdirAll(workspace, 0755)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(filepath.Join(workspace, fmt.Sprintf("%s.log", config.AppName)), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	global.Logger = log.New(f, "", log.Lshortfile|log.LstdFlags)

	err = global.Config.Load()
	if err != nil {
		panic(err)
	}
	err = global.Config.Save()
	if err != nil {
		panic(err)
	}
	fmt.Println(global.Config)
}

func main() {
	// Create an instance of the app structure
	app := NewApp()
	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	FileMenu.AddText("Open Dir", keys.CmdOrCtrl("o"), func(data *menu.CallbackData) {
		dir, err := runtime.OpenDirectoryDialog(app.ctx, runtime.OpenDialogOptions{})
		if err != nil {
			tools.Logf("fail to open dir,err=%v", err)
			return
		}
		if len(dir) == 0 {
			return
		}
		ft, err := tools.OpenDir(dir)
		if err != nil {
			tools.Logf("fail to get file tree,err=%v", err)
			return
		}
		runtime.EventsEmit(app.ctx, "opendir", ft)
	})
	FileMenu.AddSeparator()
	//FileMenu.AddText("Config", keys.CmdOrCtrl("c"), func(data *menu.CallbackData) {
	//	fmt.Println(data)
	//	str, err := runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
	//		Type:          runtime.InfoDialog,
	//		Title:         "test",
	//		Message:       "message",
	//		Buttons:       nil,
	//		DefaultButton: "",
	//		CancelButton:  "",
	//		Icon:          nil,
	//	})
	//	fmt.Println(str, err)
	//})

	FileMenu.AddText("Open File", nil, func(data *menu.CallbackData) {
		file, err := runtime.OpenFileDialog(app.ctx, runtime.OpenDialogOptions{
			Filters: []runtime.FileFilter{{DisplayName: "Markdown Files (*.md)", Pattern: "*.md"}},
		})
		if err != nil {
			tools.Logf("open file failed,err=%v", err)
		}
		if len(file) == 0 {
			return
		}
		tools.Log(file)
		ft, err := tools.OpenDir(file)
		if err != nil {
			tools.Logf("fail to get file tree,err=%v", err)
			return
		}
		fmt.Println("ft", ft)
		runtime.EventsEmit(app.ctx, "opendir", ft)
	})
	//FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
	//	runtime.Quit()
	//})

	if goruntime.GOOS == "darwin" {
		AppMenu.Append(menu.EditMenu()) // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut
	}

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "wails-demo",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		AssetsHandler:    &F{},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Menu:             AppMenu,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		tools.Logf("Error:%v", err)
	}
}

type F struct {
}

func (f F) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := os.Stat(r.URL.Path)
	if os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, err := ioutil.ReadFile(r.URL.Path)
	_, err = w.Write(data)
	if err != nil {
		tools.Logf("load file failed,err=%v", err)
	}
	return
}
