/**
 * @author: hqd
 * @description: base model
 * @file: base
 * @date: 2021-02-07 13:10
 */

package models

import (
	"time"

	"github.com/hqd8080/go-blog-server/models/form"
	"github.com/jinzhu/gorm"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"
)

// 创建时间
type CreateField struct {
	CreatedAt time.Time `json:"created_at"`
}

// 修改时间
type UpdateField struct {
	UpdatedAt time.Time `json:"updated_at"`
}

// 删除时间
type DeleteField struct {
	DeletedAt *time.Time `json:"-"`
}

// 是否删除
type IsDeleteField struct {
	IsDeleted uint `json:"is_deleted"`
}

// base model
type Model struct {
	IsDeleted uint       `json:"is_deleted"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func queryBuilder(db *gorm.DB, pageParam *form.PaginationParam, out interface{}) (int, error) {
	if pageParam != nil {
		total, err := query(db, pageParam.PageNum, pageParam.PageSize, out)
		if err != nil {
			return 0, err
		}
		return total, nil
	}
	return 0, db.Find(out).Error
}

func query(db *gorm.DB, pageIndex, pageSize int, out interface{}) (int, error) {
	var count int
	if err := db.Count(&count).Error; err != nil {
		return 0, err
	}
	if count == 0 {
		return 0, nil
	}
	db = db.Where("deleted_at IS NULL")
	err := db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(out).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
