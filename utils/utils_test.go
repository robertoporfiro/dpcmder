package utils

import (
	"testing"
)

func TestSplitOnFirst(t *testing.T) {
	testDataMatrix := [][]string{
		{"/usr/bin/share", "/", "", "usr/bin/share"},
		{"usr/bin/share", "/", "usr", "bin/share"},
		{"/share", "/", "", "share"},
		{"share", "/", "share", ""},
		{"my big testing task", " ", "my", "big testing task"},
	}
	for _, testCase := range testDataMatrix {
		gotPreffix, gotSuffix := SplitOnFirst(testCase[0], testCase[1])
		if gotPreffix != testCase[2] || gotSuffix != testCase[3] {
			t.Errorf("for SplitOnFirst('%s', '%s'): got ('%s', '%s'), want ('%s', '%s')", testCase[0], testCase[1], gotPreffix, gotSuffix, testCase[2], testCase[3])
		}
	}
}

func TestSplitOnLast(t *testing.T) {
	testDataMatrix := [][]string{
		{"/usr/bin/share", "/", "/usr/bin", "share"},
		{"usr/bin/share", "/", "usr/bin", "share"},
		{"/share", "/", "", "share"},
		{"share", "/", "share", ""},
		{"my big testing task", " ", "my big testing", "task"},
		{"local:/test1/test2", "/", "local:/test1", "test2"},
		{"local:/test1", "/", "local:", "test1"},
		{"local:", "/", "local:", ""},
	}
	for _, testCase := range testDataMatrix {
		gotPreffix, gotSuffix := SplitOnLast(testCase[0], testCase[1])
		if gotPreffix != testCase[2] || gotSuffix != testCase[3] {
			t.Errorf("for SplitOnLast('%s', '%s'): got ('%s', '%s'), want ('%s', '%s')", testCase[0], testCase[1], gotPreffix, gotSuffix, testCase[2], testCase[3])
		}
	}
}

func TestGetDpPath(t *testing.T) {
	testDataMatrix := [][]string{
		{"local:/dir1/dir2", "myfile", "local:/dir1/dir2/myfile"},
		{"local:", "myfile", "local:/myfile"},
		{"local:/dir1/dir2", "..", "local:/dir1"},
		{"local:/dir1", "..", "local:"},
		{"local:", "..", "local:"},
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
