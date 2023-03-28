package common

// PictureAnchor 图片添加水印的锚点位置
type PictureAnchor int

const (
	// Center 水印的锚点位置在中间
	Center PictureAnchor = iota
	TopLeft
	TopRight
	BottomLeft
	BottomRight
)

// 要执行的单元
const (
	// Watermark 执行添加水印的操作
	Watermark = iota
	// Resize 执行图片缩放
	Resize
	// Convert 执行转换格式
	Convert
)

// WatermarkType 添加水印的类型，图片水印还是文字水印
type WatermarkType int

const (
	// TextWatermark 文字水印
	TextWatermark WatermarkType = iota
	// PictureWatermark 图片水印
	PictureWatermark WatermarkType = iota
)

// ImageFormatType 图片格式（后缀）
type ImageFormatType int

const (
	PNG ImageFormatType = iota
	JPEG
	JPG
	GIF
	BMP
	TIF
)

var ImgFormatName = []string{"png", "jpeg", "jpg", "gif", "bmp", "tif"}
