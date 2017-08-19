# easyCSV
GolangでCSVを簡単に吐き出せるライブラリ

`[][]String` な2次元配列を投げることで、その中身をCSVとして書き出す。

# How to use

``` go
str := [][]string{
		{"a", "b", "c", "d", "e"},
		{"aa", "bb", "cc", "dd", "ee"},
		{"aaa", "bbb", "ccc", "ddd", "eee"},
		{"aaaa", "bbbb", "cccc", "dddd", "eeee"},
		{"aaaaa", "bbbbb", "ccccc", "ddddd", "eeeee"},
	}
err := easyCSV.ExportCSV(str, "sample")
```
