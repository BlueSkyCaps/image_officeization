package test

import (
	"encoding/json"
	"image"
	"image/color"
	"image_officeization/core/src"
	"image_officeization/core/src/common"
	"image_officeization/core/src/watermark"
)

func Run() {
	var waterInputParams = watermark.WaterInputParams{
		Anchor: common.Center,
		Offset: image.Point{},
		// 原图片绝对路径
		Paths: []string{"static/test/img/liuyifei.png"},
		// 输出目录
		OutDir: "static/test/img",
		// 水印类型
		WatermarkType: common.TextWatermark,
		TextWaterInputParams: watermark.TextWaterInputParams{
			FontPath: "C:/Windows/Fonts/simhei.ttf",
			// 字体大小
			FontSize: 64,
			// RGBA颜色与透明度
			RGBA: color.RGBA{R: 255, G: 0, B: 0, A: 255},
			// 文字水印值
			Text: "@刘亦菲-LiuYiFei-520\n我爱你\n你是我的生",
		},
	}

	bytes, err := json.Marshal(waterInputParams)
	if err != nil {
		return
	}
	src.InitWatermarkInput(string(bytes))
}
