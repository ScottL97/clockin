package models

type QrcodeInfo struct {
	Url string
}


func GetQrcodeInfo() QrcodeInfo {
	rtn := QrcodeInfo{Url: "/static/img/qrcode.png"}
	return rtn
}
