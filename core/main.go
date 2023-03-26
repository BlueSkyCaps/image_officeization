package main

import (
	"image_officeization/core/src"
	"image_officeization/core/src/common"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		// args[0]外部命令行调用时的此进程文件名
		execType, _ := strconv.ParseInt(args[1], 10, 32)
		if execType == common.Watermark {
			src.InitWatermarkInput(args[2])
		} else if execType == common.Resize {
			src.InitResizeInput(args[2])
		} else if execType == common.Convert {
			src.InitConvertInput(args[2])
		}
	}
}