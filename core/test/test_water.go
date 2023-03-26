package test

//
//import (
//	"github.com/disintegration/imaging"
//	"github.com/golang/freetype"
//	"github.com/golang/freetype/truetype"
//	"golang.org/x/image/math/fixed"
//	"image"
//	"image/color"
//	"image_officeization/core/src/common"
//	"image_officeization/core/src/watermark"
//	"io/ioutil"
//	"log"
//	"strconv"
//)
//
//func WaterTest(opType int, params interface{}) {
//	if opType == 0 {
//		inputParams := params.(watermark.TextWaterInputParams)
//		addText(inputParams)
//		return
//	}
//	addImage()
//}
//
//func addImage() {}
//
//func addText(params watermark.TextWaterInputParams) {
//	for i := 0; i < len(params.WaterInput.Paths); i++ {
//		dst := dstNew(params.WaterInput.Paths[i])
//		fontObj := parseFontObj(params.FontPath)
//		textWidth := advanceTextWidth(fontObj, params)
//		drawTextContextSet(fontObj, textWidth, dst, params)
//		// 将带有水印的图片输出到本地
//		err := imaging.Save(dst, params.WaterInput.OutDir+"/out_watermark_"+strconv.Itoa(i)+".jpg")
//		if err != nil {
//			return
//		}
//	}
//}
//
//func advanceTextWidth(fontObj *truetype.Font, params watermark.TextWaterInputParams) int {
//	face := truetype.NewFace(fontObj, &truetype.Options{
//		Size: params.FontSize,
//	})
//	textWidth := 0
//	for _, ch := range params.Text {
//		// GlyphAdvance方法返回字符的Advance（前进宽度）和是否存在该字符的布尔值。
//		advance, ok := face.GlyphAdvance(ch)
//		if ok {
//			// 需要将Advance向右移6位，因为Advance的单位是1/64像素。
//			textWidth += int(advance >> 6)
//		}
//	}
//	return textWidth
//}
//
//// 设置绘制文字水印的上下文
//func drawTextContextSet(fontObj *truetype.Font, textWidth int, dst *image.NRGBA, params watermark.TextWaterInputParams) {
//	// 创建一个新的绘图上下文
//	c := freetype.NewContext()
//	c.SetFont(fontObj)
//	/*
//		// 设置像素每英寸，这个值不具有参考意义，最终取的仍是原图片的DPI。
//		// 网络图片输出一般都是72，而打印时清晰的精度一般为300
//	*/
//	c.SetDPI(72)
//	// 设置字体大小
//	c.SetFontSize(params.FontSize)
//	// 设置要绘制的图片尺寸，取得就是原图片尺寸
//	c.SetClip(dst.Bounds())
//	// 设置需要去绘制的目标图片，就是原图片
//	c.SetDst(dst)
//	// 设置要去绘制的图片颜色，因为要绘制文字，这就是文字的颜色
//	c.SetSrc(image.NewUniform(color.RGBA{R: 255, G: 255, B: 255, A: 255}))
//	// 在靠近图片中心（非正中心，因为无法计算完整字体水印所占的宽度）添加水印
//	var pt fixed.Point26_6
//	switch params.WaterInput.Anchor {
//	case common.Center:
//		pt = freetype.Pt(dst.Bounds().Max.X/2-textWidth/2+params.WaterInput.Offset.Y,
//			dst.Bounds().Max.Y/2-int(params.FontSize)+params.WaterInput.Offset.Y)
//
//	case common.TopLeft:
//		pt = freetype.Pt(0+params.WaterInput.Offset.X, 0+int(params.FontSize)+params.WaterInput.Offset.Y)
//
//	case common.TopRight:
//		pt = freetype.Pt(dst.Bounds().Max.X-textWidth+params.WaterInput.Offset.X,
//			0+int(params.FontSize)+params.WaterInput.Offset.Y)
//
//	case common.BottomLeft:
//		pt = freetype.Pt(0+params.WaterInput.Offset.X,
//			dst.Bounds().Max.Y-int(params.FontSize)+params.WaterInput.Offset.Y)
//
//	case common.BottomRight:
//		pt = freetype.Pt(dst.Bounds().Max.X-textWidth+params.WaterInput.Offset.X,
//			dst.Bounds().Max.Y-int(params.FontSize)+params.WaterInput.Offset.Y)
//	}
//	// 绘在位置点制水印文字
//	_, err := c.DrawString(params.Text, pt)
//	if err != nil {
//		return
//	}
//}
//
//// 返回要解析的字体对象
//func parseFontObj(fontPath string) *truetype.Font {
//	// 读取字体文件
//	fontBytes, err := ioutil.ReadFile(fontPath)
//	if err != nil {
//		panic(err)
//	}
//	// 解析字体文件
//	parseFont, err := freetype.ParseFont(fontBytes)
//	return parseFont
//}
//
//// 创建要操作的图片，从原始图片中来
//func dstNew(srcImgPath string) *image.NRGBA {
//	// 打开原始图片
//	src, err := imaging.Open(srcImgPath)
//	if err != nil {
//		log.Fatalf("Failed to open image: %v", err)
//	}
//	// 创建一个新的图片，大小和原始图片一样
//	dst := imaging.New(src.Bounds().Max.X, src.Bounds().Max.Y, color.NRGBA{})
//	// 复制原始图片到新图片
//	dst = imaging.Paste(dst, src, image.Pt(0, 0))
//	return dst
//}
