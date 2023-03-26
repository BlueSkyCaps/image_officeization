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
