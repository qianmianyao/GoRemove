package mainFunc

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type conf struct {
	Size int `yaml:"size"`
}

// 读取配置
func (c *conf) readConf() int {
	confPath := homePath() + "/.config/goremove/conf.yml" // 获取到用户的家目录
	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Printf(Red, "Error: Configuration file conf.yml could not be found")
	}
	if err := yaml.Unmarshal(yamlFile, c); err != nil { // 读取配置
		log.Fatalf("Unmarshal: %v", err)
	}
	// fmt.Println(c.Size)
	return c.Size // 返回配置参数
}

// 获取用户家目录
func homePath() string {
	info, _ := user.Current()
	return info.HomeDir
}

// 不允许操作根目录
func isRoot(file []string) *[]string {
	for _, value := range file {
		if path, _ := filepath.Split(value); path == "/" {
			log.Printf(Red, "Note: Try to use grm -r [file] and grm -m [file]")
			os.Exit(1)
		}
	}
	return &file
}

// 列出回收站的文件
func listRecycle() {
	recycleBin, _ := ioutil.ReadDir(Tmp)
	// 列出回收站文件
	for _, value := range recycleBin {
		fmt.Printf(Red, value.Name())
	}
	fmt.Println("↑ These are the recycle bin files ↑")
}

// 判断文件大小
func fileSize(FileOrPath *[]string, fileSize int64) bool {
	for _, value := range *FileOrPath {
		if size, _ := os.Stat(value); size.Size() >= fileSize {
			return true
		}
	}
	return false
}
