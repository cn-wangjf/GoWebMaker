package main

import (
	"fmt"
	"html/template"
	"io"
	"os"
)

const (
	EXPORT_PATH = "export/"
)

type ExportWriter struct {
	pageName  string
	clearFile bool
}

func NewExportWriter(pageName string) *ExportWriter {
	return &ExportWriter{pageName: pageName, clearFile: true}
}

func (this *ExportWriter) Write(p []byte) (n int, err error) {

	pageFile := EXPORT_PATH + this.pageName

	if this.clearFile {
		os.Remove(pageFile)
		this.clearFile = false
	}

	f, err := os.OpenFile(pageFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(fmt.Sprintf("ExportWriter open error:%v", err))
		return
	}
	defer f.Close()
	n, err = f.Write(p)
	if err == nil && n < len(p) {
		err = io.ErrShortWrite
		fmt.Println(fmt.Sprintf("ExportWriter write error:%v", err))
		return
	}
	return
}

func main() {
	tpl, err := template.ParseFiles("template/index.tpl")
	if err != nil {
		fmt.Println(fmt.Sprintf("template error:%v", err))
	}
	writer := NewExportWriter("index.html")
	data := make(map[string]string)
	data["content"] = "hello world! 你好！"
	err = tpl.Execute(writer, data)
	if err != nil {
		fmt.Println(fmt.Sprintf("write error:%v", err))
	}

}
