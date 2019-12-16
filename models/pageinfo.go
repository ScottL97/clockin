package models

type PageInfo struct {
	Name string
	Email string
}

func GetPageInfo() PageInfo {
	rtn := PageInfo{Name: "Scott", Email: "2687538884@qq.com"}
	return rtn
}
