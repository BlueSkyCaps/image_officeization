package common

import (
	"bytes"
	"github.com/disintegration/imaging"
	"image"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func SaveImgFile(filePath string, saveOutDir string, dst image.Image) {
	fullName, _, _ := GetFileNameAndFormat(filePath)
	var savePathF = saveOutDir + "/" + strconv.FormatInt(time.Now().Unix(), 10) + "-" + fullName
	// 使用默认的EncodeOption。目标格式若是jpg/jpeg，设置图片质量为100。
	err := imaging.Save(dst, savePathF, imaging.JPEGQuality(100))
	if err != nil {
		println("saveImgFile error：" + err.Error())
		return
	}
}

func SaveImgFileByTheFormat(filePath string, saveOutDir string, dst image.Image, formatName string,
	formatType ImageFormatType, encodeImgBuf bytes.Buffer) {
	_, name, _ := GetFileNameAndFormat(filePath)
	var savePathF = saveOutDir + "/" + strconv.FormatInt(time.Now().Unix(), 10) + "-" + name + "." + formatName
	var err error
	if formatType == WEBP {
		if err = OverrideFile(savePathF, encodeImgBuf); err != nil {
			println(err)
			os.Exit(ExitFileOp)
		}
		return
	}
	// 使用默认的EncodeOption。目标格式若是jpg/jpeg，设置图片质量为100。
	err = imaging.Save(dst, savePathF, imaging.JPEGQuality(100))
	if err != nil {
		println("saveImgFile error：" + err.Error())
		os.Exit(ExitImagingSave)
		return
	}
}

func SaveImgFileByCompress(filePath string, saveOutDir string, dst image.Image, quality ImgCompressQuality) {
	_, name, _ := GetFileNameAndFormat(filePath)
	// 无论是何种格式，imaging自行处理，统一成jpeg的格式保存
	var savePathF = saveOutDir + "/" + strconv.FormatInt(time.Now().Unix(), 10) + "-" + name + "." + "jpeg"
	// 通过imaging提供的EncodeOption，设置JPEG的质量数值，达到压缩体积的目的。
	err := imaging.Save(dst, savePathF,
		imaging.JPEGQuality(ImgCompressJPEGValues[quality]))
	if err != nil {
		println("saveImgFile error：" + err.Error())
		return
	}
}

// GetFileNameAndFormat 根据文件路径获取文件名、文件名（不含后缀格式）以及后缀格式
func GetFileNameAndFormat(filePath string) (fullName string, name string, format string) {
	filePath = strings.ReplaceAll(filePath, "\\", "/")
	//fullName = filePath[strings.LastIndex(filePath, "/")+1:]
	fullName = path.Base(filePath)
	lastIndex := strings.LastIndex(fullName, ".")
	name = fullName[0:lastIndex]
	format = strings.TrimLeft(path.Ext(fullName), ".")
	return fullName, name, format
}

func IsWebP(srcFormat string) bool {
	return strings.ToLower(srcFormat) == ImgFormatName[6]
}

func SameFormat(srcFormat string, dstFormat string) bool {
	return strings.ToLower(srcFormat) == strings.ToLower(dstFormat)

}

// OverrideFile 打开一个文件往里覆盖数据。文件不存在则创建且则覆盖数据。
func OverrideFile(filePath string, buf bytes.Buffer) error {
	err := os.WriteFile(filePath, buf.Bytes(), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
