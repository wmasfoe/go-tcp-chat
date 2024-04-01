package log

import "fmt"

type BaseLog interface {
	Warn(msg ...any)
	Error(msg ...any)
	Info(msg ...any)
}

type ConnectLog interface {
	NewUser(ip string)
	GlobalLog(ip string, msg ...any)
}

type Log struct{}

func (l Log) Warn(msg ...any) {
	fmt.Printf("[warn] %v \n", msg...)
}
func (l Log) Error(msg ...any) {
	fmt.Printf("[error] %v \n", msg...)
}
func (l Log) Info(msg ...any) {
	fmt.Printf("[info] %v \n", msg...)
}
func (l Log) NewUser(ip string) {
	fmt.Printf("[system-info] 有新IP加入辣! %v \n", ip)
}
func (l Log) GlobalLog(ip string, msg ...any) {
	fmt.Printf("[system-info] %v ", ip)
	fmt.Printf("在全局喊话: %v \n", msg...)
}

var Logger = new(Log)
