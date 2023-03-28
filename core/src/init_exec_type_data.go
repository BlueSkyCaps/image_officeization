package src

import (
	"encoding/json"
	"image_officeization/core/src/convert"
	"image_officeization/core/src/resize"
	"image_officeization/core/src/watermark"
	"os"
)

func InitWatermarkInput(jsonWaterInputData string) {
	// 反序列化为WaterInputParams
	// 解析JSON字符串为结构体实例
	var waterInputParams watermark.WaterInputParams
	err := json.Unmarshal([]byte(jsonWaterInputData), &waterInputParams)
	if err != nil {
		println(err.Error())
		os.Exit(-3)
	}
	watermark.Run(waterInputParams)
}

func InitResizeInput(jsonResizeInputData string) {
	// 反序列化为ResizeInputParams
	// 解析JSON字符串为结构体实例
	var resizeInputParams resize.ResizeInputParams
	err := json.Unmarshal([]byte(jsonResizeInputData), &resizeInputParams)
	if err != nil {
		println(err.Error())
		os.Exit(-3)
	}
	resize.Run(resizeInputParams)
}

func InitConvertInput(jsonConvertInputData string) {
	// 反序列化为ResizeInputParams
	// 解析JSON字符串为结构体实例
	var convertInputParams convert.ConvertInputParams
	err := json.Unmarshal([]byte(jsonConvertInputData), &convertInputParams)
	if err != nil {
		println(err.Error())
		os.Exit(-3)
	}
	convert.Run(convertInputParams)
}
