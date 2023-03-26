package watermark

import (
	"github.com/disintegration/imaging"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image_officeization/core/src/common"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Run(params WaterInputParams) {
	if params.WatermarkType == common.TextWatermark {
		addText(params)
		return
	}
	addImage()
}

func addImage() {}

func addText(params WaterInputParams) {
	for i := 0; i < len(params.Paths); i++ {
		dst := DstNew(params.Paths[i])
		fontObj := parseFontObj(params.TextWaterInputParams.FontPath)
		textWidthMax, texts := advanceTextWidth(fontObj, params.TextWaterInputParams)
		drawTextContextSet(fontObj, textWidthMax, texts, dst, params)
		// 将带有水印的图片输出到本地
		err := imaging.Save(dst, params.OutDir+"/out_watermark_"+strconv.Itoa(i)+".jpg")
		if err != nil {
			return
		}
	}
}

// 根据设置的字体大小计算文字水印所占的宽度，获取最大的宽度的那一行文本（文本水印有时不止一行）,以及分割了的多行字符串
func advanceTextWidth(fontObj *truetype.Font, params TextWaterInputParams) (int, []string) {
	face := truetype.NewFace(fontObj, &truetype.Options{
		Size: params.FontSize,
	})
	texts := strings.Split(params.Text, "\n")
	textWidthMax := 0
	for i := 0; i < len(texts); i++ {
		textWidthTmp := 0
		for _, ch := range texts[i] {
			// GlyphAdvance方法返回字符的Advance（前进宽度）和是否存在该字符的布尔值。
			advance, ok := face.GlyphAdvance(ch)
			if ok {
				// 需要将Advance向右移6位，因为Advance的单位是1/64像素。
				textWidthTmp += int(advance >> 6)
			}
		}
		if textWidthTmp > textWidthMax {
			textWidthMax = textWidthTmp
		}
	}
	return textWidthMax, texts
}

// 设置绘制文字水印的上下文
func drawTextContextSet(fontObj *truetype.Font, textWidthMax int, texts []string, dst *image.NRGBA, params WaterInputParams) {
	// 创建一个新的绘图上下文
	c := freetype.NewContext()
	c.SetFont(fontObj)
	/*
		// 设置像素每英寸，这个值不具有参考意义，最终取的仍是原图片的DPI。
		// 网络图片输出一般都是72，而打印时清晰的精度一般为300
	*/
	c.SetDPI(72)
	// 设置字体大小
	c.SetFontSize(params.TextWaterInputParams.FontSize)
	// 设置要绘制的图片尺寸，取得就是原图片尺寸
	c.SetClip(dst.Bounds())
	// 设置需要去绘制的目标图片，就是原图片
	c.SetDst(dst)
	// 设置要去绘制的图片颜色，因为要绘制文字，这就是文字的颜色
	c.SetSrc(image.NewUniform(color.RGBA{R: 255, G: 255, B: 255, A: 255}))
	// 开始根据字体大小的增量来绘制当前行的文本
	var currentAddY int
	for i := 0; i < len(texts); i++ {
		drawTextHandler(dst, textWidthMax, c, params, currentAddY, i+1, len(texts), texts[i])
		currentAddY = currentAddY + int(params.TextWaterInputParams.FontSize)
	}
}

func drawTextHandler(dst *image.NRGBA, textWidth int, c *freetype.Context, params WaterInputParams,
	currentAddY int, currentLineIndex int, lineTotal int, currentTextLine string) {
	var pt fixed.Point26_6
	switch params.Anchor {
	case common.Center:
		pt = freetype.Pt(dst.Bounds().Max.X/2-textWidth/2+params.Offset.Y,
			dst.Bounds().Max.Y/2-int(params.TextWaterInputParams.FontSize)+currentAddY+params.Offset.Y)

	case common.TopLeft:
		pt = freetype.Pt(0+params.Offset.X, 0+currentAddY+int(params.TextWaterInputParams.FontSize)+params.Offset.Y)

	case common.TopRight:
		pt = freetype.Pt(dst.Bounds().Max.X-textWidth+params.Offset.X,
			0+int(params.TextWaterInputParams.FontSize)+currentAddY+params.Offset.Y)

	case common.BottomLeft:
		/*
			// 底部位置，因为要考虑多行文本，第一行的应该根据当前行算出最大的增量值，越往后的行增量值递减
			// 避免多行的时候，越是在前面的行反而在越后面绘制（被倒序绘制了）
		*/
		currentAddY = -(int(params.TextWaterInputParams.FontSize) * (lineTotal - currentLineIndex + 1))
		pt = freetype.Pt(0+params.Offset.X,
			dst.Bounds().Max.Y+currentAddY+params.Offset.Y)

	case common.BottomRight:
		/*
			// 底部位置，因为要考虑多行文本，第一行的应该根据当前行算出最大的增量值，越往后的行增量值递减
			// 避免多行的时候，越是在前面的行反而在越后面绘制（被倒序绘制了）
		*/
		currentAddY = -(int(params.TextWaterInputParams.FontSize) * (lineTotal - currentLineIndex + 1))
		pt = freetype.Pt(dst.Bounds().Max.X-textWidth+params.Offset.X,
			dst.Bounds().Max.Y+currentAddY+params.Offset.Y)
	default:
		// 没有指定有效的Anchor，使用自定义位置
		pt = freetype.Pt(params.Point.X+params.Offset.X,
			params.Point.Y+params.Offset.Y)
	}
	// 绘在位置点制水印文字
	_, err := c.DrawString(currentTextLine, pt)
	if err != nil {
		return
	}
}

// 返回要解析的字体对象
func parseFontObj(fontPath string) *truetype.Font {
	// 读取字体文件
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		panic(err)
	}
	// 解析字体文件
	parseFont, err := freetype.ParseFont(fontBytes)
	return parseFont
}

// DstNew 创建要操作的图片，从原始图片中来
func DstNew(srcImgPath string) *image.NRGBA {
	// 打开原始图片
	src, err := imaging.Open(srcImgPath)
	if err != nil {
		log.Fatalf("Failed to open image: %v", err)
	}
	// 创建一个新的图片，大小和原始图片一样
	dst := imaging.New(src.Bounds().Max.X, src.Bounds().Max.Y, color.NRGBA{})
	// 复制原始图片到新图片
	dst = imaging.Paste(dst, src, image.Pt(0, 0))
	return dst
}
