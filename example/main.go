package main

import (
	"fmt"
	xxl "github.com/xxl-job/xxl-job-executor-go"
	"github.com/xxl-job/xxl-job-executor-go/example/task"
	"github.com/xxl-job/xxl-job-executor-go/log"
	"os"
	"path"
)

func main() {
	logdir := "./"

	logOption := log.NewOptions()
	logOption.LogFileDir = logdir
	logOption.InfoFileName = "info.log"

	exec := xxl.NewExecutor(
		xxl.ServerAddr("http://127.0.0.1:8080/xxl-job-admin"),
		xxl.AccessToken(""),            //请求令牌(默认为空)
		xxl.ExecutorIp("127.0.0.1"),    //可自动获取
		xxl.ExecutorPort("9999"),       //默认9999（非必填）
		xxl.RegistryKey("golang-jobs"), //执行器名称
		xxl.SetLogDir(logdir),
		xxl.SetLogger(log.New(logOption)), //自定义日志
	)
	exec.Init()
	//设置日志查看handler
	exec.LogHandler(func(req *xxl.LogReq) *xxl.LogRes {
		localFile := path.Join(logdir, fmt.Sprintf("%v.log", req.LogID))
		var message string
		file, err := os.Open(localFile)
		if err != nil {
			message = "没有读都数据"
		} else {
			fileinfo, _ := file.Stat()
			filesize := fileinfo.Size()
			buffer := make([]byte, filesize)
			file.Read(buffer)
			message = string(buffer)
		}
		defer file.Close()

		return &xxl.LogRes{Code: 200, Msg: "", Content: xxl.LogResContent{
			FromLineNum: req.FromLineNum,
			ToLineNum:   2,
			LogContent:  message,
			IsEnd:       true,
		}}
	})
	//注册任务handler
	exec.RegTask("task.test", task.Test)
	err := exec.Run()
	if err != nil {
		log.Errorf("start failed:%v", err.Error())
	}
}

////xxl.Logger接口实现
//type logger struct{}
//
//func (l *logger) Info(format string, a ...interface{}) {
//	fmt.Println(fmt.Sprintf("自定义日志 - "+format, a...))
//}
//
//func (l *logger) Error(format string, a ...interface{}) {
//	log.Println(fmt.Sprintf("自定义日志 - "+format, a...))
//}
