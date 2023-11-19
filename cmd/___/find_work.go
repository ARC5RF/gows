package ___

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/ARC5RF/fsup"
)

type GoWork struct {
	Go        string
	Toolchain string
	Use       []Use
	Replace   []Replace
}

type Use struct {
	DiskPath   string
	ModulePath string
}

type Replace struct {
	Old Module
	New Module
}

type Module struct {
	Path    string
	Version string
}

func FindWorkFile() (*GoWork, string, error) {
	res, err := fsup.FromWD("go.work")
	if err != nil {
		return nil, "", err
	}
	return_point := filepath.Dir(res)
	os.Chdir(return_point)

	cmd := exec.Command("go", "work", "edit", "-json")
	var wb bytes.Buffer
	cmd.Stdout = &wb
	cmd.Run()

	wf := &GoWork{}
	json.Unmarshal(wb.Bytes(), wf)
	return wf, return_point, nil
}
