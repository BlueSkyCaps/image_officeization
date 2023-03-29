package common

// ExitCode 退出状态码
const (
	// ExitSuccess 成功退出
	ExitSuccess = iota
	// ExitExecTypeNotMatch 操作单元未匹配
	ExitExecTypeNotMatch = iota
	// ExitJsonUnmarshal json解析问题
	ExitJsonUnmarshal = iota
	// ExitImagingSave Imaging库保存问题
	ExitImagingSave = iota
	// ExitSameFormatConvert 相同格式图片请求转换
	ExitSameFormatConvert = iota
	// ExitWebpDecode webp库 decode问题
	ExitWebpDecode = iota
	// ExitWebpEncode webp库 encode问题
	ExitWebpEncode = iota
	// ExitFileOp 文件（图片）操作问题
	ExitFileOp = iota
)
