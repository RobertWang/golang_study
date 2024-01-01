package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"syscall"
)

func main() {
	// 打开 ELF 二进制文件
	file, err := os.Open("app.bin")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	hook_file, err := os.Open("funcs.so")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer hook_file.Close()

	// 获取 ELF 二进制文件的程序入口地址
	entry, err := getEntry(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 在程序入口处写入跳转指令，跳转到 hook 函数中
	hookAddress, err := getHookAddress(hook_file)
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := syscall.Write(int(entry), []byte("\xc3")); err != nil {
		fmt.Println(err)
		return
	}
	if _, err := syscall.Write(int(entry), []byte("\x8b")); err != nil {
		fmt.Println(err)
		return
	}
	if _, err := syscall.Write(int(entry), []byte("movl $0x"+string(hookAddress)+"\n")); err != nil {
		fmt.Println(err)
		return
	}
	if _, err := syscall.Write(int(entry), []byte("ret")); err != nil {
		fmt.Println(err)
		return
	}
}

// 获取 ELF 二进制文件的程序入口地址
func getEntry(file *os.File) (uintptr, error) {
	var entry uintptr
	if err := binary.Read(file, binary.LittleEndian, &entry); err != nil {
		return 0, err
	}
	return entry, nil
}

// 获取 hook 函数的地址
func getHookAddress(file *os.File) (uintptr, error) {
	var hookAddr uintptr
	if err := binary.Read(file, binary.LittleEndian, &hookAddr); err != nil {
		return 0, err
	}
	return hookAddr, nil
}
