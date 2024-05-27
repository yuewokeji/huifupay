package request

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"mime/multipart"
	"net/http"
)

func NewFileRequest(url string, data interface{}) *FileRequest {
	req := &BaseRequest{}
	req.SetURL(url)
	req.SetData(data)

	freq := &FileRequest{
		BaseRequest: req,
	}
	return freq
}

type FileRequest struct {
	*BaseRequest

	file        io.Reader
	writeToBody func() error
}

func (f *FileRequest) SetFile(file io.Reader) {
	f.file = file
}

func (f *FileRequest) StartWriteToBody() error {
	if nil != f.writeToBody {
		return f.writeToBody()
	}
	return nil
}

// Build 构建http.Request
// https://paas.huifu.com/partners/api#/shgl/shjj/api_shjj_shtpsc?id=%e6%8e%a5%e5%8f%a3%e8%af%b4%e6%98%8e
// 最近更新时间：2024.3.1
func (f *FileRequest) Build() (req *http.Request, err error) {
	pr, pw := io.Pipe()
	writer := multipart.NewWriter(pw)
	fn := func() error {
		defer pw.Close()
		if f.file != nil {
			part, err := writer.CreateFormFile("file", "huifupay_file_upload")
			if err != nil {
				return errors.Wrap(err, "create form file")
			}

			_, err = io.Copy(part, f.file)
			if err != nil {
				return errors.Wrap(err, "copy file")
			}
		}

		data, err := json.Marshal(f.GetData())
		if err != nil {
			return errors.Wrap(err, "marshal `data`")
		}
		writer.WriteField("data", string(data))
		writer.WriteField("sys_id", f.GetSysID())
		writer.WriteField("product_id", f.GetProductID())
		err = writer.Close()
		if err != nil {
			return errors.Wrap(err, "close multipart writer")
		}
		return nil
	}
	f.writeToBody = fn

	req, err = http.NewRequest("POST", f.url, pr)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return
}
