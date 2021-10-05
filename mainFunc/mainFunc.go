package mainFunc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

const Red = "\u001B[1;31;40m%s\u001B[0m\n"

var Tmp = homePath() + "/.tmp/"

// MainFunc 主逻辑函数
func MainFunc(FileOrPath *[]string) {
	files := *FileOrPath
	var c conf
	size := c.readConf() // 多大的文件直接删除
	file := isRoot(files)
	if fileSize(file, int64(size)) == true {
		Remover(file)
	} else {
		Recycle(file)
	}
}

// Remover 删除
func Remover(FileOrPath *[]string) {
	for _, value := range *FileOrPath {
		if err := os.RemoveAll(value); err != nil {
			log.Printf(Red, "Error: Delete failure!")
		}
	}
}

// Recycle 回收
func Recycle(FileOrPath *[]string) {
	var tmp string
	for _, value := range *FileOrPath {
		filename := path.Base(value) // 获取文件名称
		tmp = Tmp + filename         // 拼接字符串
		if err := os.Rename(value, tmp); err != nil {
			log.Printf(Red, "Error: Recovery of failure")
			// log.Println(err)
		}
	}
}

// Empty 清空回收站
func Empty() {
	listRecycle()
	var yesAndNo string
	fmt.Print("Whether to empty the recycle bin [y/n]: ")
	if _, err := fmt.Scanln(&yesAndNo); err != nil {
		return
	}
	switch yesAndNo {
	case "y":
		if err := os.RemoveAll(Tmp); err != nil {
			log.Printf(Red, "Error: Failed to empty the recycle bin, Try using: sudo grm -e")
			// log.Println(err)
		}
	}

}

// Restore 文件恢复
func Restore() {
	listRecycle()
	// 接收多个输入，转换成切片
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Restore files to current directory, It can be multiple: ")
	file, _ := input.ReadString('\n')
	files := strings.Fields(file)

	for _, value := range files {
		oldFiles := Tmp + value
		newFile := "./" + value
		if err := os.Rename(oldFiles, newFile); err != nil {
			log.Printf(Red, "Error: File recovery failed")
			log.Println(err)
		}
	}

}

// Init 初始化
func Init() {
	// 建立回收站
	if err := os.MkdirAll(Tmp, 0764); err != nil {
		return
	}
	// 建立配置文件目录
	confPath := homePath() + "/.config/goremove/"
	if err := os.MkdirAll(confPath, 0764); err != nil {
		return
	}
}

func Help() {
	fmt.Print(`
Usage:  grm [OPTIONS] FileOrPath

Options:
	-h		See the help;
	-r		Delete files directly;
	-m		Move files to the recycle bin;
	-e		Empty the recycle bin;
	-b		Restore files
	init	Initial chemical tool;
Config:
	Path: Configuration file path /etc/grm/conf.yml;
	Size: Files larger than size will be deleted and will not be reclaimed;

`)
}
