package resize

import (
	"github.com/disintegration/imaging"
	"image_officeization/core/src/common"
	"log"
)

func Run(params ResizeInputParams) {
	for i := 0; i < len(params.Paths); i++ {
		// 打开原始图片
		src, err := imaging.Open(params.Paths[i])
		if err != nil {
			log.Fatalf("Failed to open image: %v", err)
		}

		// 创建一个新的图片，大小和原始图片一样
		resImg := imaging.Resize(src, params.WH.X, params.WH.Y, imaging.Lanczos)
		// 将缩放后的图片输出到本地
		common.SaveImgFile(params.Paths[i], params.OutDir, resImg)
	}
}
