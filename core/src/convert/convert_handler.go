package convert

import (
	"bytes"
	"github.com/chai2010/webp"
	"image"
	"image_officeization/core/src/common"
	"image_officeization/core/src/watermark"
	"os"
	"time"
)

func Run(params ConvertInputParams) {
	for i := 0; i < len(params.Paths); i++ {
		runCurrent(params, i)
		// 必须在当前迭代暂停一段时间，以等待图片保存完毕。
		// 如果图片源存在多个文件名一样的图片（如a.jpeg/a.png），则其中一张会消失不见
		time.Sleep(time.Millisecond * 1000)
	}

}
func runCurrent(params ConvertInputParams, i int) {
	_, _, srcFormat := common.GetFileNameAndFormat(params.Paths[i])
	if common.SameFormat(srcFormat, common.ImgFormatName[params.ImageFormatType]) {
		println("convert same to same format Image")
		os.Exit(common.ExitSameFormatConvert)
	}
	var dstNew image.Image
	// 若原图片是webp格式 webp -> another
	if common.IsWebP(srcFormat) {
		webpImg, err := os.Open(params.Paths[i])
		dstNew, err = webp.Decode(webpImg)
		err = webpImg.Close()
		if err != nil {
			println(err.Error())
			os.Exit(common.ExitWebpDecode)
		}
	} else {
		dstNew = watermark.DstNew(params.Paths[i])
	}
	formatName := common.ImgFormatName[params.ImageFormatType]
	encodeImgBuf := bytes.Buffer{}
	if params.ImageFormatType == common.WEBP {
		// 若目标图片是webp格式 another -> webp
		encodeImgBuf = encodeWebPImg(dstNew)
	}
	common.SaveImgFileByFormat(params.Paths[i], params.OutDir, dstNew, formatName, params.ImageFormatType, encodeImgBuf)
	//var a rune
	//_, err := fmt.Scanln(&a)
	//if err != nil {
	//	return
	//}
}

func encodeWebPImg(mDecode image.Image) bytes.Buffer {
	buf := bytes.Buffer{}
	var err error
	if err = webp.Encode(&buf, mDecode, &webp.Options{Lossless: true, Quality: 100}); err != nil {
		println(err)
		os.Exit(common.ExitWebpEncode)
	}
	return buf
}
