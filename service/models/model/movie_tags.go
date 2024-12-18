package model

import (
	"time"
)

const TableNameMovieTagsModel = "movie_tags"

// MovieTagsModel mapped from table <movie_tags>
type MovieTagsModel struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MovieID    string    `gorm:"column:movie_id;not null" json:"movie_id"`
	Tag        string    `gorm:"column:tag" json:"tag"`
	Source     string    `gorm:"column:source;not null;comment:标签来源" json:"source"` // 标签来源
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName MovieTagsModel's table name
func (*MovieTagsModel) TableName() string {
	return TableNameMovieTagsModel
}
