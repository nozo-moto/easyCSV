package easyCSV

import (
	"io/ioutil"
	"os"
	"strings"
)

func Exportfile(str string, filename string)(error){
	content := []byte(str)
	ioutil.WriteFile("./" + filename + ".csv" , content, os.ModePerm)
	return nil
}

func JoinTexts(texts [][]string)(string){
	var joinedTexts string
	for i := range texts{
		joinedTexts += strings.Join(texts[i][:], ", ") + "\n"
	}
	return joinedTexts
}

func ExportCSV(texts [][]string, filename string) (error) {
	joinedTexts := JoinTexts(texts)

	err := Exportfile(joinedTexts, filename)
	if err != nil {
		return err
	}
	return nil
}
