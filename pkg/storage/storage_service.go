package storage

import (
	"nginx_mate_go/pkg/constant"
	"nginx_mate_go/pkg/path"
)

var storage = NewFsStorage(path.NginxStorage)

func init() {
	storage.AddNamespace(constant.StorageNsStaticTcpRoute)
}

func Acquire() *FsStorage {
	return storage
}
