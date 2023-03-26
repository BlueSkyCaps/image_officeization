package convert

import "image_officeization/core/src/common"

type ConvertInputParams struct {
	// 原图片绝对路径
	Paths []string
	// 输出目录
	OutDir string
	// 想得到的转换格式
	ImageFormatType common.ImageFormatType
}
