package tools

import (
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	Dir  = "dir"
	File = "file"
)

type FileTree struct {
	Name      string      `json:"name"`
	TotalPath string      `json:"totalPath"`
	Type      string      `json:"type"`
	Content   string      `json:"content"`
	Children  []*FileTree `json:"children,omitempty"`
}

func OpenDir(dirName string) (*FileTree, error) {
	f, err := fs.Stat(os.DirFS(filepath.Dir(dirName)), filepath.Base(dirName))
	if err != nil {
		Logf("open dir failed,err=%v", err)
	}
	return generateFileTree(filepath.Dir(dirName), f)
}

func generateFileTree(dirName string, info os.FileInfo) (*FileTree, error) {
	ft := &FileTree{
		Name:      info.Name(),
		TotalPath: dirName + string(os.PathSeparator) + info.Name(),
		Type:      Dir,
		Content:   "",
		Children:  nil,
	}
	if !info.IsDir() {
		ft.Type = File
		return ft, nil
	}
	files, err := ioutil.ReadDir(ft.TotalPath)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(files); i++ {
		if !files[i].IsDir() {
			// 判断是否是md后缀
			if !(filepath.Ext(files[i].Name()) == ".md") {
				continue
			}
		}
		child, err := generateFileTree(ft.TotalPath, files[i])
		if err != nil {
			return nil, err
		}
		ft.Children = append(ft.Children, child)
	}
	return ft, nil
}

func ReadFileContent(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func SaveFileContent(filename string, content string) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	n, err := f.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}
	_, err = f.WriteAt([]byte(content), n)
	return err
}

type SaveFileArg struct {
	TotalPath string `json:"totalPath"`
	Content   string `json:"content"`
}

type UploadFileArg struct {
	Content []byte `json:"content"`
}

func BatchSaveFileContent(files []SaveFileArg) error {
	for i := 0; i < len(files); i++ {
		err := SaveFileContent(files[i].TotalPath, files[i].Content)
		if err != nil {
			return err
		}
	}
	return nil
}
