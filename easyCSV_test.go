package easyCSV

import (
	"testing"
)

func TestJoinTexts(t *testing.T) {
	str := [][]string{
		{"a", "b", "c", "d", "e"},
		{"aa", "bb", "cc", "dd", "ee"},
		{"aaa", "bbb", "ccc", "ddd", "eee"},
		{"aaaa", "bbbb", "cccc", "dddd", "eeee"},
		{"aaaaa", "bbbbb", "ccccc", "ddddd", "eeeee"},
	}
	if JoinTexts(str) != "a, b, c, d, e\naa, bb, cc, dd, ee\naaa, bbb, ccc, ddd, eee\naaaa, bbbb, cccc, dddd, eeee\naaaaa, bbbbb, ccccc, ddddd, eeeee\n" {
		t.Error("err")
	}
}

func TestExportCSV(t *testing.T) {
	str := [][]string{
		{"a", "b", "c", "d", "e"},
		{"aa", "bb", "cc", "dd", "ee"},
		{"aaa", "bbb", "ccc", "ddd", "eee"},
		{"aaaa", "bbbb", "cccc", "dddd", "eeee"},
		{"aaaaa", "bbbbb", "ccccc", "ddddd", "eeeee"},
	}
	err := ExportCSV(str, "test")
	if err != nil {
		t.Error("err")
	}
}