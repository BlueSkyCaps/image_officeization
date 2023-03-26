package test

import (
	"github.com/disintegration/imaging"
	"log"
)

func ResizeTest() {
	// 打开原始图片
	src, err := imaging.Open("static/test/img/liuyifei.jpeg")
	if err != nil {
		log.Fatalf("Failed to open image: %v", err)
	}

	// 创建一个新的图片，大小和原始图片一样
	resImg := imaging.Resize(src, 2800, 2452, imaging.Lanczos)
	// 将缩放后的图片输出到本地
	err = imaging.Save(resImg, "static/test/img/output_resize.jpg")
	if err != nil {
		panic(err)
	}
}
