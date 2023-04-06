package test

import (
	"bytes"
	"github.com/disintegration/imaging"
	"image"
	"image/gif"
	"os"
	"strconv"
)

func GifAllTest() {
	inputGifBoss := &gif.GIF{}
	// 打开原始GIF文件
	f, err := os.Open("C:\\Users\\BlueSkyCarry\\Desktop\\微信图片_20230405212507.gif")
	if err != nil {
		panic(err.Error())
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err.Error())
		}
	}(f)

	// 解码GIF文件
	gifImg, err := gif.DecodeAll(f)
	if err != nil {
		panic(err.Error())
	}

	// 迭代每一帧图像
	for i, frame := range gifImg.Image {
		_ = i
		var u = strconv.Itoa(i)
		u2 := string(u[len(u)-1])
		if u2 == "1" || u2 == "3" || u2 == "5" || u2 == "7" || u2 == "9" {
			continue
		}
		// 调整每帧图像
		buf := new(bytes.Buffer)
		_ = imaging.Encode(buf, frame, imaging.GIF, imaging.GIFNumColors(256))
		decodeGif, err := gif.Decode(buf)
		if err != nil {
			return
		}
		//inputGifBoss.Image = append(inputGifBoss.Image, frame)
		inputGifBoss.Image = append(inputGifBoss.Image, decodeGif.(*image.Paletted))
		inputGifBoss.Delay = append(inputGifBoss.Delay, 0)
	}
	inputGifBoss.Config = image.Config{ColorModel: inputGifBoss.Config.ColorModel, Width: inputGifBoss.Image[0].Bounds().Dx(),
		Height: inputGifBoss.Image[0].Bounds().Dy()}
	newFile, err := os.Create("static/test/img/compressed.gif")
	if err != nil {
		panic(err.Error())
	}
	defer func(newFile *os.File) {
		err := newFile.Close()
		if err != nil {

		}
	}(newFile)

	// 将GIF图像编码写入文件
	if err := gif.EncodeAll(newFile, inputGifBoss); err != nil {
		panic(err.Error())
	}
}
