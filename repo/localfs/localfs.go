package localfs

import (
	"fmt"
	"github.com/croz-ltd/dpcmder/config"
	"github.com/croz-ltd/dpcmder/model"
	"github.com/croz-ltd/dpcmder/utils/logging"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strconv"
)

type LocalRepo struct {
	name string
}

// Repo is localfs implementation of repo/Repo interface.
var Repo = LocalRepo{name: "LocalRepo"}

// GetInitialParent returns initialy opened local directory info.
func (r LocalRepo) GetInitialParent() model.CurrentView {
	logging.LogDebug("repo/localfs/GetInitialParent()")
	currPath, err := filepath.Abs(*config.LocalFolderPath)
	if err != nil {
		logging.LogFatal("repo/localfs/GetInitialParent(): ", err)
	}

	initialView := model.CurrentView{Path: currPath}
	return initialView
}

// GetList returns list of items for current directory.
func (r LocalRepo) GetList(currentView model.CurrentView) model.ItemList {
	currPath := currentView.Path
	logging.LogDebug(fmt.Sprintf("repo/localfs/ListCurrent('%s')", currPath))

	parentDir := model.Item{Type: model.ItemDirectory, Name: "..", Size: "", Modified: "", Selected: false}
	items := listFiles(currPath)

	itemsWithParentDir := make([]model.Item, 0)
	itemsWithParentDir = append(itemsWithParentDir, parentDir)
	itemsWithParentDir = append(itemsWithParentDir, items...)

	return itemsWithParentDir
}

func (r LocalRepo) GetTitle(view model.CurrentView) string {
	return view.Path
}

func listFiles(dirPath string) []model.Item {
	logging.LogDebug(fmt.Sprintf("repo/localfs/listFiles('%s')", dirPath))
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		logging.LogFatal("repo/localfs/listFiles(): ", err)
	}

	items := make(model.ItemList, len(files))

	for idx, file := range files {
		var dirType model.ItemType
		if file.IsDir() {
			dirType = model.ItemDirectory
		} else {
			dirType = model.ItemFile
		}
		items[idx] = model.Item{Type: dirType, Name: file.Name(), Size: strconv.FormatInt(file.Size(), 10), Modified: file.ModTime().Format("2006-01-02 15:04:05"), Selected: false}
	}

	sort.Sort(items)

	return items
}