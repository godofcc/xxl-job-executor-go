package task

import (
	"context"
	"fmt"
	xxl "github.com/xxl-job/xxl-job-executor-go"
)

func Test(cxt context.Context, param *xxl.RunReq, logger xxl.Logger) (msg string) {
	logger.Info("xxxx")
	fmt.Println("test one task" + param.ExecutorHandler + " param：" + param.ExecutorParams + " log_id:" + xxl.Int64ToStr(param.LogID))
	return "test done"
}
