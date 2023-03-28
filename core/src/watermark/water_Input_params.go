package watermark

import (
	"image"
	"image/color"
	"image_officeization/core/src/common"
)

// WaterInputParams 给图片处理水印所需要传递的数据
type WaterInputParams struct {
	// 图片添加水印的锚点位置，当此参数指定了有效值，WaterInputParams.point被忽略
	Anchor common.PictureAnchor
	// 图片添加水印的坐标（像素），当WaterInputParams.anchor指定了有效值，此参数被忽略
	Point image.Point
	// 设置偏移量（像素）
	Offset image.Point
	// 原图片绝对路径
	Paths []string
	// 输出目录
	OutDir string
	// 水印类型
	WatermarkType         common.WatermarkType
	TextWaterInputParams  TextWaterInputParams
	ImageWaterInputParams ImageWaterInputParams
}

// TextWaterInputParams 给图片处理水印的类型为文本，所需要传递的数据
type TextWaterInputParams struct {
	// 选取的字体路径
	FontPath string
	// 字体大小
	FontSize float64
	// RGBA颜色与透明度
	RGBA color.RGBA
	// 文字水印值
	Text string
}

// ImageWaterInputParams 给图片处理水印的类型为图片水印，所需要传递的数据
type ImageWaterInputParams struct {
	// 图片水印的绝对路径
	WaterImgPath string
}
