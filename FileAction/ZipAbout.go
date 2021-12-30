/*
@Time : 2021/12/30 6:02 下午
@Author : yuyunqing
@File : ZipAbout
@Software: GoLand
*/
package FileAction

import (
	"bytes"
	"compress/zlib"
	"io"
)

//进行zlib压缩
func DoZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

//进行zlib解压缩
func DoZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}
