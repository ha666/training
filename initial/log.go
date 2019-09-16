package initial

import "github.com/ha666/logs"

func initLog() {
	logs.SetLogger(logs.AdapterConsole, `{"level":7}`)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
}
