package task

import (
	"context"
	"fmt"
	xxl "github.com/xxl-job/xxl-job-executor-go"
	"github.com/xxl-job/xxl-job-executor-go/log"
)

func Test(cxt context.Context, param *xxl.RunReq, logger log.Logger) (msg string) {
	logger.Info("xxxx")
	fmt.Println("test one task" + param.ExecutorHandler + " param：" + param.ExecutorParams + " log_id:" + xxl.Int64ToStr(param.LogID))
	logger.Info("vvss")
	return "test done"
}
