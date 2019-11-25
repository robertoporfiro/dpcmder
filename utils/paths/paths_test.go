package paths

import (
	"reflect"
	"testing"
)

func TestGetDpPath(t *testing.T) {
	testDataMatrix := [][]string{
		{"local:/dir1/dir2", "myfile", "local:/dir1/dir2/myfile"},
		{"local:", "myfile", "local:/myfile"},
		{"local:/dir1/dir2", "..", "local:/dir1"},
		{"local:/dir1/dir2", ".", "local:/dir1/dir2"},
		{"local:/dir1", "..", "local:"},
		{"local:", "..", "local:"},
		{"local:", ".", "local:"},
	}
	for _, testCase := range testDataMatrix {
		newPath := GetDpPath(testCase[0], testCase[1])
		if newPath != testCase[2] {
			t.Errorf("for GetFilePath('%s', '%s'): got '%s', want '%s'", testCase[0], testCase[1], newPath, testCase[2])
		}
	}
}

func TestGetFilePathUsingSeparator(t *testing.T) {
	testDataMatrix := [][]string{
		{"/usr/bin/share", "myfile", "/", "/usr/bin/share/myfile"},
		{"", "myfile", "/", "myfile"},
		{"/usr/bin/share", "..", "/", "/usr/bin"},
		{"/testdir", "..", "/", "/"},
		{"", "..", "/", ""},
		{"/", "..", "/", "/"},
		{"/", "testfile", "/", "/testfile"},
	}
	for _, testCase := range testDataMatrix {
		newPath := GetFilePathUsingSeparator(testCase[0], testCase[1], testCase[2])
		if newPath != testCase[3] {
			t.Errorf("for GetFilePathUsingSeparator('%s', '%s', '%s'): got '%s', want '%s'", testCase[0], testCase[1], testCase[2], newPath, testCase[3])
		}
	}
}

func TestSplitDpPath(t *testing.T) {
	testDataMatrix := []struct {
		path       string
		components []string
	}{
		{"", []string{}},
		{"local:", []string{"local:"}},
		{"local:/dir1", []string{"local:", "dir1"}},
		{"local:/dir1/dir2", []string{"local:", "dir1", "dir2"}},
	}
	for _, testCase := range testDataMatrix {
		got := SplitDpPath(testCase.path)
		want := testCase.components
		if !reflect.DeepEqual(got, want) {
			t.Errorf("for SplitDpPath('%s'): got %#v, want %#v", testCase.path, got, want)
		}
	}
}
