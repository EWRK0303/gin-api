package models

import (
	"gorm.io/gorm"
)

type Bucket struct {
	gorm.Model
	Bucketname string //bucket名字
	Userid     string //哪个用户创建的
}

type File struct {
	gorm.Model
	BucketId string //文件所属的bucket_id
	FileName string //文件名字
}
