package main

import (
	"fmt"
	"log"
	"os/exec"
	"flag"
	"strings"
	"os"
)

func printStart() {
	fmt.Printf("-------------- apk安装 ----\n")
	fmt.Printf("\n")
}

func selectDevice(adb_path string) string {
	cmd := exec.Command(adb_path, "devices")
    out, err := cmd.CombinedOutput()

    if err != nil {
        fmt.Printf("exec error:\n%s\n", string(out))
        log.Fatalf("failed with %s\n", err)
    }
	
	deviceStr := string(out)
	strArr := strings.Split(deviceStr, "\n")

	for index, value := range strArr{
		if (0 != index) {
			fmt.Println(index, " =", value)
		}
	}

	// 读取输入
	fmt.Println("输入要安装到哪个设备")
	var num int
	fmt.Scan(&num)

	dstDevice := strArr[num]
	// fmt.Printf("dst device: %s \n", dstDevice)

	id := strings.Fields(dstDevice)[0]
	return id
}

func installApk(adb_path string, apk_path string, id string) {
	if "" == apk_path {
		return
	}

	cmd := exec.Command(adb_path, "-s", id, "install", apk_path)
    out, err := cmd.CombinedOutput()

    if err != nil {
        fmt.Printf("失败 :\n%s\n", string(out))
        // log.Fatalf("failed with %s\n", err)
    } else {

		fmt.Printf("执行成功:\n%s\n", string(out))
	}

}

func main() {
	
	var adb_path string
	var apk_path string
	var apk_path2 string
	var apk_path3 string
	var apk_path4 string
	var apk_path5 string
	flag.StringVar(&adb_path, "adb", "", "adb 位置")
	flag.StringVar(&apk_path, "apk", "", "apk 位置")
	flag.StringVar(&apk_path2, "apk2", "", "apk2 位置")
	flag.StringVar(&apk_path3, "apk3", "", "apk3 位置")
	flag.StringVar(&apk_path4, "apk4", "", "apk4 位置")
	flag.StringVar(&apk_path5, "apk5", "", "apk5 位置")

	flag.Parse()

	if "" == adb_path {
		log.Fatalf("adb_path 必填")
	}

	if "" == apk_path {
		log.Fatalf("apk_path 必填")
	}

	printStart()

	fmt.Printf("---- 以下是入参 ----\n")
	fmt.Printf("adb path: %s\n", adb_path)
	fmt.Printf("apk path: %s\n", apk_path)
	fmt.Printf("---- 入参 end ----\n\n")

	// 选择设备
	id := selectDevice(adb_path)
	fmt.Println("正在安装apk到设备 =", id)
	fmt.Println("请稍后...")

	// 安装
	installApk(adb_path, apk_path, id)
	installApk(adb_path, apk_path2, id)
	installApk(adb_path, apk_path3, id)
	installApk(adb_path, apk_path4, id)
	installApk(adb_path, apk_path5, id)

	fmt.Println("输入任意字符退出")
    b := make([]byte, 1)
    os.Stdin.Read(b)
}