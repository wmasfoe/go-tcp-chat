package log

import (
	"chat-server/pkg"
	"fmt"
	"reflect"
)

type ClassicLog interface {
	Warn(msg ...any)
	Error(msg ...any)
	Info(msg ...any)
}

type ConnectLog interface {
	NewUser(ip string)
	GlobalLog(ip string, msg ...any) string
}

type Log struct{}

func baseLog(t string, args ...any) {
	logType := fmt.Sprintf("[%v]", t)

	val := reflect.ValueOf(args)

	fmt.Printf("%v ", logType)

	if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
		for _, arg := range args {
			fmt.Printf("%v ", arg)
		}
	}

	fmt.Print("\n")
}

func (l Log) Warn(msg ...any) {
	baseLog("warn", msg...)
}
func (l Log) Error(msg ...any) {
	baseLog("error", msg...)
}
func (l Log) Info(msg ...any) {
	baseLog("info", msg...)
}
func (l Log) NewUser(ip string) {
	baseLog("system-info", fmt.Sprintf("有新IP加入辣! %v", ip))
}
func (l Log) UserExit(ip string) {
	baseLog("system-info", fmt.Sprintf("%v 已下线", ip))
}

func (l Log) GlobalLog(ip string, msg ...any) {
	fmt.Printf("[system-info] %v ", ip)
	fmt.Printf("在全局喊话: %v \n", msg...)
}

func (l Log) GlobalLogStr(ip string, msg ...string) string {
	joinedMsg := pkg.JoinStr(msg...)
	return fmt.Sprintf("[system-info] %v 在全局喊话: %v", ip, joinedMsg)
}
func (l Log) UserExitStr(ip string) string {
	return fmt.Sprintf("[system-info] %v 已下线 \n", ip)
}
func (l Log) NewUserStr(ip string) string {
	return fmt.Sprintf("[system-info] 有新IP加入辣! %v \n", ip)
}

var Logger = new(Log)
