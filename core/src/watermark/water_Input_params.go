/*
Package watermark 图片添加水印。
采用Imaging库封装大量代码去判断图片格式,包括encode和decode操作。
采用freetype库进行文本水印绘制和字体设置。
源图片支持jpeg/jpg png gif bmp tif，但不支持webp格式（通过webp库多次转码将耗时，且是没必要的，可通过convert单元转换格式）。
支持根据Anchor(common.PictureAnchor)锚点来绘制水印位置，也可通过image.Point自定义坐标点，但优先使用Anchor。
目前支持文字水印类型，绘制图片水印不保证后续会添加。
RGBA设置颜色，A为透明度。
字体需要指定，系统已安装的字体绝对路径位置，windows通常为：c:/windows/font/xxx.tif
也可以修改源码，尝试通过Golang获取有效的字体路径，并指定给WaterInputParams.TextWaterInputParams.FontPath
*/
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
