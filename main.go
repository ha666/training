package main

import (
	"github.com/ha666/golibs"
	"github.com/ha666/logs"
	"github.com/ha666/training/addSub"
	"github.com/ha666/training/initial"
	_ "github.com/ha666/training/initial"
	"runtime"
	"time"
)

const VERSION = "2019.906.1713"

func init() {

	//region 初始化系统设置
	runtime.GOMAXPROCS(runtime.NumCPU())
	//endregion

	//region 输出当前系统信息

	//region 输出golang版本信息
	logs.Info("【go】version:%s", runtime.Version())
	//endregion

	//region 输出系统信息
	logs.Info("【sys】os:%s", runtime.GOOS)
	logs.Info("【sys】cpu:%d", runtime.NumCPU())
	//endregion

	//region 输出网络信息
	logs.Info("【net】ip:%s", initial.ServerIP)
	//endregion

	//region 输出应用信息
	logs.Info("【app】path:%s", golibs.GetCurrentDirectory())
	logs.Info("【app】version:%s", VERSION)
	//endregion

	//endregion
}

func main() {
	defer time.Sleep(1e9)
	addSub.Do()
}
