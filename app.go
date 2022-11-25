package main

import (
	"changeme/config"
	"changeme/global"
	"changeme/tools"
	"context"
	"github.com/google/uuid"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io/fs"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.EventsOn(a.ctx, "log", func(optionalData ...interface{}) {
		tools.Log(optionalData...)
	})
}

// ReadFile returns content
func (a *App) ReadFile(filename string) (string, error) {
	return tools.ReadFileContent(filename)
}

func (a *App) SaveFile(filename string, content string) error {
	return tools.SaveFileContent(filename, content)
}

func (a *App) SaveFiles(files []tools.SaveFileArg) error {
	return tools.BatchSaveFileContent(files)
}

// 获取图片URL
func (a *App) UploadFiles(file tools.UploadFileArg) (string, error) {
	if global.Config.Enable {
		picgo, err := exec.LookPath("picgo")
		if err != nil {
			tools.Logf("picgo not install")
			goto Local
		}

		// 查看类型是否符合
		if !tools.Contains(config.SupportUploader, global.Config.Picgo.PicBed.Uploader) {
			tools.Logf("uploader  %s not support", global.Config.Picgo.PicBed.Uploader)
			goto Local
		}

		// 上传
		cmd := exec.Command(picgo, "upload", "-c", filepath.Join(config.UserHomeDir, config.AppDir, config.DataName))
		data, err := cmd.Output()
		if err != nil {
			tools.Logf("upload failed err=%v", err)
			goto Local
		}
		if strings.Contains(string(data), "[PicGo ERROR]") || strings.Contains(string(data), "[PicGo WARN]") {
			tools.Logf("upload failed err=%v", string(data))
			goto Local
		}

		result := string(data)

		patten := "[PicGo SUCCESS]:"
		index := strings.LastIndex(result, patten)
		if index < 0 {
			tools.Logf("upload failed err=%v", result)
			goto Local
		}
		url := result[index+len(patten):]

		tools.Logf("upload success url=%s", url)

		return url, nil
	}

Local:
	name := filepath.Join(global.Config.GetImagePath(), uuid.New().String()+".png")
	err := ioutil.WriteFile(name, file.Content, fs.ModePerm)
	if err != nil {
		return "", err
	}
	return name, nil

}

// 更新配置文件
func (a *App) SaveConfig(cfg config.Config) error {
	// TODO 原子性
	global.Config = cfg
	return global.Config.Save()
}

// 读取配置文件
func (a *App) GetConfig() (config.Config, error) {
	return global.Config, nil
}

// 获取文件目录
func (a *App) GetDir() (string, error) {
	dir, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		tools.Logf("fail to open dir,err=%v", err)
		return "", err
	}
	return dir, nil
}

func (a *App) Log(args ...any) {
	tools.Log(args...)
}
