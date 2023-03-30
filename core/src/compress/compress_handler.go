package compress

import (
	"github.com/disintegration/imaging"
	"image_officeization/core/src/common"
	"os"
	"time"
)

func Run(params CompressInputParams) {
	for i := 0; i < len(params.Paths); i++ {
		// 打开原始图片
		src, err := imaging.Open(params.Paths[i])
		if err != nil {
			println("Failed to open image: " + err.Error())
			os.Exit(common.ExitFileOp)
		}
		// 根据压缩质量将图片输出到本地
		common.SaveImgFileByCompress(params.Paths[i], params.OutDir, src, params.ImgCompressQuality)
		time.Sleep(1000 * time.Millisecond)
	}
}
