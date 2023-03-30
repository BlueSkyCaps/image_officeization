/*
Package compress 压缩图片体积大小。
采用Imaging库封装大量代码去判断图片格式,包括encode和decode操作。
源图片支持jpeg/jpg png gif bmp tif，但不支持webp格式（通过webp库多次转码将耗时，且是没必要的，可通过convert包转换格式）。
imaging内部采用stdlib标准库进行图片的编解码，无论是何种格式，都将其编码成jpeg格式，并通过压缩级别设置imaging.JPEGQuality，
改变jpeg的图像质量，以此方式来达到压缩效果。此方式压缩后体积减小明显且质量可观。
所以，jpeg/jpeg gif png bmp tif的源图片都会以jpeg格式输出。
ImgCompressQuality.High设置的数值>80，压缩效果很好。ImgCompressQuality.Low能明显观察到图片清晰度减低，体积压缩程度更多。
经过测试，无论是1.3M的png图片还是400kb的jpeg图片，压缩等级为High，输出体积均为160kb左右；压缩等级为Low，输出体积为50kb左右
*/
package compress

import "image_officeization/core/src/common"

type CompressInputParams struct {
	// 原图片绝对路径
	Paths []string
	// 输出目录
	OutDir string
	// 压缩质量级别
	ImgCompressQuality common.ImgCompressQuality
}
