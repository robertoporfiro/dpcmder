package repo

import (
	"github.com/croz-ltd/dpcmder/model"
)

// Repo is a common repository methods implemented by local filesystem and DataPower
type Repo interface {
	// new methods
	GetInitialParent() model.CurrentView
	GetList(currentView model.CurrentView) model.ItemList
	GetTitle(view model.CurrentView) string

	// deprecated below
	InitialLoad(m *model.Model)
	LoadCurrent(m *model.Model)
	EnterCurrentDirectoryMissingPassword(m *model.Model) bool
	EnterCurrentDirectorySetPassword(m *model.Model, password string) bool
	EnterCurrentDirectory(m *model.Model)
	ListFiles(m *model.Model, dirPath string) []model.Item
	GetFileType(m *model.Model, parentPath, fileName string) model.ItemType
	GetFileTypeFromPath(m *model.Model, filePath string) model.ItemType
	GetFileName(filePath string) string
	GetFilePath(parentPath, fileName string) string
	GetFile(m *model.Model, parentPath, fileName string) []byte
	UpdateFile(m *model.Model, parentPath, fileName string, newFileContent []byte) bool
	Delete(m *model.Model, parentPath, fileName string) bool
	CreateDir(m *model.Model, parentPath, dirName string) bool
	IsEmptyDir(m *model.Model, parentPath, dirName string) bool
}