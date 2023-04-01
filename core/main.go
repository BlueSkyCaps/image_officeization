package main

import (
	"image_officeization/core/src"
	"image_officeization/core/src/common"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if len(args) >= 3 {
		// args[0]外部命令行调用时的此进程文件名
		execType, _ := strconv.ParseInt(args[1], 10, 32)
		// 调用方程序把空格替换成"?"以便命令行参数能够有效传递而不会被截断，此处将其复原回来
		decodeString := strings.ReplaceAll(args[2], "?", " ")
		//err := os.WriteFile("C:\\Users\\BlueSkyCarry\\Desktop\\aa.txt", []byte(decodeString), fs.ModePerm)
		//if err != nil {
		//	return
		//}
		if execType == common.Watermark {
			src.InitWatermarkInput(decodeString)
		} else if execType == common.Resize {
			src.InitResizeInput(decodeString)
		} else if execType == common.Convert {
			src.InitConvertInput(decodeString)
		} else if execType == common.Compress {
			src.InitCompressInput(decodeString)
		} else {
			os.Exit(common.ExitExecTypeNotMatch)
		}
		os.Exit(common.ExitSuccess)
	}
	println("please pass args!")
	// 非成功退出
	os.Exit(common.ExitExecTypeNotMatch)
	//_, _ = fmt.Scanln()
}
