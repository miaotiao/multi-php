package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var phpMap map[string]string

func main() {

	Init()
	fmt.Println(SysPath)
	fmt.Println(PhpPath)
	return

	initConf()

	args := os.Args
	if len(args) < 2 {
		fmt.Println("xxx set|get [参数]")
		return
	}

	switch args[1] {
	case "set":
		setPhp(args[2])
	case "add":
		addPhp(args[2], args[3])
	}
	getAllPhp()
	execPhp()

}

func initConf() {
	//	获取文件内容
	raw, err := ioutil.ReadFile("./php.txt")
	if err != nil {
		fmt.Println("读取 php.json 文件失败")
		return
	}
	phpMap = make(map[string]string)
	rows := strings.Split(strings.ReplaceAll(string(raw), "\r\n", "\n"), "\n")
	for _, item := range rows {
		if item != "" {
			col := strings.Fields(item)
			phpMap[col[0]] = col[1]
		}

	}
}

func getAllPhp() {
	//	获取当前 php 环境
	currentPhpPath := os.Getenv("php")
	//	todo 如果为空，就提醒
	for key, item := range phpMap {
		//	todo 判断文件是否存在
		console := key + " " + item
		if item == currentPhpPath {
			fmt.Println("--> " + console)
		} else {
			fmt.Println("    " + console)
		}
	}
}

func addPhp(pKey string, pPath string) {
	if phpMap[pKey] != "" {
		fmt.Println(pKey + " 已经存在，请使用其他名称")
		return
	}
	phpMap[pKey] = pPath
	var txt string
	for key, val := range phpMap {
		txt += key + " " + val + "\r\n"
	}
	ioutil.WriteFile("./php.txt", []byte(txt), 0666)
}

func setPhp(key string) {
	path := phpMap[key]
	fmt.Println()
	if path == "" {
		fmt.Println("路径为空: " + key + " " + path)
	}
	err := os.Setenv("php", path)
	if err != nil {
		fmt.Println(path + " 设置失败: " + err.Error())
	}
}

func execPhp() {
	cmd := exec.Command("php", "-v")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(out))
}
