package models

import (
	"gorm.io/gorm"
)

type Bucket struct {
	gorm.Model
	Bucketname string //bucket名字
	Userid     string //哪个用户创建的
}

type file struct {
	gorm.Model
	bucket_id int    //文件所属的bucket_id
	fileName  string //文件名字
}

// 检查 Bucket 是否属于用户
func IsBucketOwnedByUser(username, bucketName string) bool {
	var count int64
	database.DB.Model(&models.UserBucket{}).
		Where("username = ? AND bucket_name = ?", username, bucketName).
		Count(&count)
	return count > 0
}
