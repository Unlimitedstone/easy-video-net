package model

import (
	"time"
)

const TableNameMovieMessageModel = "movie_message"

// MovieMessageModel mapped from table <movie_message>
type MovieMessageModel struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UniID       string    `gorm:"column:uni_id;not null;comment:唯一id" json:"uni_id"`                   // 唯一id
	Name        string    `gorm:"column:name;not null;comment:电影名字" json:"name"`                       // 电影名字
	ImageURL    string    `gorm:"column:image_url;comment:电影url" json:"image_url"`                     // 电影url
	ImageSrc    string    `gorm:"column:image_src" json:"image_src"`                                   // 本地资源src
	Score       string    `gorm:"column:score;comment:电影评分" json:"score"`                              // 电影评分
	Type        string    `gorm:"column:type;not null;comment:类型：Move:电影，TV：电视剧，Anime：动漫" json:"type"` // 类型：Move:电影，TV：电视剧，Anime：动漫
	Common      string    `gorm:"column:common;comment:精选评论" json:"common"`                            // 精选评论
	Source      string    `gorm:"column:source;not null" json:"source"`
	CreateTime  time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
	Introduce   string    `gorm:"column:introduce;comment:简介" json:"introduce"`         // 简介
	Alias_      string    `gorm:"column:alias;comment:别名" json:"alias"`                 // 别名
	Duration    int32     `gorm:"column:duration;comment:时长（分钟）" json:"duration"`       // 时长（分钟）
	ReleaseDate string    `gorm:"column:release_date;comment:上映日期" json:"release_date"` // 上映日期
	Director    string    `gorm:"column:director;comment:导演" json:"director"`           // 导演
	Actor       string    `gorm:"column:actor;comment:主演" json:"actor"`                 // 主演
	Writer      string    `gorm:"column:writer;comment:编剧" json:"writer"`               // 编剧
	Country     string    `gorm:"column:country;comment:制片国家" json:"country"`           // 制片国家
	Imdb        string    `gorm:"column:imdb;comment:imdb id" json:"imdb"`              // imdb id
}

// TableName MovieMessageModel's table name
func (*MovieMessageModel) TableName() string {
	return TableNameMovieMessageModel
}
