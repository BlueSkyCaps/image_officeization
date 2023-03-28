package convert

import (
	"image_officeization/core/src/common"
	"image_officeization/core/src/watermark"
)

func Run(params ConvertInputParams) {
	for i := 0; i < len(params.Paths); i++ {
		dstNew := watermark.DstNew(params.Paths[i])
		formatName := common.ImgFormatName[params.ImageFormatType]
		common.SaveImgFileByFormat(params.Paths[i], params.OutDir, dstNew, formatName, params.ImageFormatType)
	}
}
