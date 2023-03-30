/*
Package convert 图片格式转换。
采用Imaging库封装大量代码去判断图片格式,包括encode和decode操作。
imaging内部采用stdlib标准库进行图片的编解码，因此jpeg/jpg png gif bmp tif格式可以直接进行互相转化。
采用webp库进行webp格式与其它格式之间的互转，因此打包编译时，程序体积多占了~8M。
*/
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
