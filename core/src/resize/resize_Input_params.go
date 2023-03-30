/*
Package resize缩放图片尺寸，以像素为宽高单位。
采用Imaging库封装大量代码去判断图片格式,包括encode和decode操作。
支持jpeg/jpg png gif bmp tif，但不支持webp格式。
同样可以进行放大，但不会超过原图片最大尺寸。
*/
package resize

import (
	"image"
)

// ResizeInputParams 图片缩放需要的数据
type ResizeInputParams struct {
	// 宽和高（像素）
	WH image.Point
	// 原图片绝对路径
	Paths []string
	// 输出目录
	OutDir string
}
