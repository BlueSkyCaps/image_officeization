package convert

import (
	"github.com/disintegration/imaging"
	"image_officeization/core/src/common"
	"image_officeization/core/src/watermark"
	"strconv"
)

func Run(params ConvertInputParams) {
	for i := 0; i < len(params.Paths); i++ {
		dstNew := watermark.DstNew(params.Paths[i])
		formatName := common.ImgFormatName[params.ImageFormatType]
		if params.ImageFormatType == common.JPEG {
			err := imaging.Save(dstNew, params.OutDir+"/out_convert_"+strconv.Itoa(i)+"."+formatName,
				imaging.JPEGQuality(100))
			if err != nil {
				panic("error in convert save:" + err.Error())
			}
			return
		}
		err := imaging.Save(dstNew, params.OutDir+"/out_convert_"+strconv.Itoa(i)+"."+formatName)
		if err != nil {
			panic("error in convert save:" + err.Error())
		}

	}
}
