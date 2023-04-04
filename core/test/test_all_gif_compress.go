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
	f, err := os.Open("C:\\Users\\BlueSkyCarry\\Desktop\\2S60X}}KL%HZ`H8$MX8MIRA.gif")
	if err != nil {
		panic(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	// 解码GIF文件
	gifImg, err := gif.DecodeAll(f)
	if err != nil {
		panic(err)
	}

	// 压缩每一帧图像
	for i, frame := range gifImg.Image {
		var u = strconv.Itoa(i)
		u2 := string(u[len(u)-1])
		if u2 == "1" || u2 == "3" || u2 == "5" || u2 == "9" {
			continue
		}
		//nrgba := imaging.Resize(frame, frame.Rect.Max.X, frame.Rect.Max.Y, imaging.MitchellNetravali)
		//paletted := image.NewPaletted(nrgba.Rect, palette.WebSafe)
		//draw.Draw(paletted, paletted.Rect, frame, paletted.Rect.Min, draw.Over)
		// 调整每帧图像的质量
		buf := new(bytes.Buffer)
		_ = imaging.Encode(buf, frame, imaging.GIF, imaging.GIFNumColors(256))
		//decodeJpeg, _, err := image.Decode(buf)
		//buf.Reset()
		//if err != nil {
		//	return
		//}
		//err = gif.Encode(buf, decodeJpeg, nil)
		//if err != nil {
		//	return
		//}
		decodeGif, err := gif.Decode(buf)
		if err != nil {
			return
		}
		inputGifBoss.Image = append(inputGifBoss.Image, decodeGif.(*image.Paletted))
		inputGifBoss.Delay = append(inputGifBoss.Delay, 0)
	}
	inputGifBoss.Config = image.Config{ColorModel: inputGifBoss.Config.ColorModel, Width: inputGifBoss.Image[0].Bounds().Dx(),
		Height: inputGifBoss.Image[0].Bounds().Dy()}
	newFile, err := os.Create("static/test/img/compressed.gif")
	if err != nil {
		panic(err)
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
