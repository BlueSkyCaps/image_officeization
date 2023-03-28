package common

import (
	"github.com/disintegration/imaging"
	"image"
	"path"
	"strconv"
	"strings"
	"time"
)

func SaveImgFile(filePath string, saveOutDir string, dst *image.NRGBA) {
	filePath = strings.ReplaceAll(filePath, "\\", "/")
	fullName := path.Base(filePath)
	var savePathF = saveOutDir + "/" + strconv.FormatInt(time.Now().Unix(), 10) + "-" + fullName
	err := imaging.Save(dst, savePathF)
	if err != nil {
		println("saveImgFile error：" + err.Error())
		return
	}
}

func SaveImgFileByFormat(filePath string, saveOutDir string, dst *image.NRGBA, formatName string, formatType ImageFormatType) {
	filePath = strings.ReplaceAll(filePath, "\\", "/")
	fullName := path.Base(filePath)
	lastIndex := strings.LastIndex(fullName, ".")
	name := fullName[0:lastIndex]
	var savePathF = saveOutDir + "/" + strconv.FormatInt(time.Now().Unix(), 10) + "-" + name + "." + formatName
	var err error
	// 使用默认的EncodeOption。目标格式若是jpg/jpeg，设置图片质量为100。
	err = imaging.Save(dst, savePathF, imaging.JPEGQuality(100))
	if err != nil {
		println("saveImgFile error：" + err.Error())
		return
	}
}
