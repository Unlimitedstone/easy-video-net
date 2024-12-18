package response

import (
	"easy-video-net/models/contribution/video"
	"easy-video-net/models/home/rotograph"
	"easy-video-net/utils/conversion"
	"time"
)

// 首页轮播图
type rotographInfo struct {
	Title string `json:"title"`
	Cover string `json:"cover"`
	Color string `json:"color"`
	Type  string `json:"type"`
	ToId  uint   `json:"to_id"`
}
type rotographInfoList []rotographInfo

// VideoInfo 首页视频
type VideoInfo struct {
	ID            uint      `json:"id"`
	Uid           uint      `json:"uid" `
	Title         string    `json:"title" `
	Video         string    `json:"video"`
	Cover         string    `json:"cover" `
	VideoDuration int64     `json:"video_duration"`
	Label         []string  `json:"label"`
	Introduce     string    `json:"introduce"`
	Heat          int       `json:"heat"`
	BarrageNumber int       `json:"barrageNumber"`
	Username      string    `json:"username"`
	CreatedAt     time.Time `json:"created_at"`
}

type MovieMessageInfo struct {
	ID          int64     `json:"id"`
	UniID       string    `json:"uni_id"`    // 唯一id
	Name        string    `json:"name"`      // 电影名字
	ImageURL    string    `json:"image_url"` // 电影url
	ImageSrc    string    `json:"image_src"` // 本地图片
	Score       string    `json:"score"`     // 电影评分
	Type        string    `json:"type"`      // 类型：Move:电影，TV：电视剧，Anime：动漫
	Common      string    `json:"common"`    // 精选评论
	Source      string    `json:"source"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
	Introduce   string    `json:"introduce"`    // 简介
	Alias       string    `json:"alias"`        // 别名
	Duration    int32     `json:"duration"`     // 时长（分钟）
	ReleaseDate string    `json:"release_date"` // 上映日期
	Director    string    `json:"director"`     // 导演
	Actor       string    `json:"actor"`        // 主演
	Writer      string    `json:"writer"`       // 编剧
	Country     string    `json:"country"`      // 制片国家
	Imdb        string    `json:"imdb"`         // imdb id
}

type videoInfoList []VideoInfo

type GetHomeInfoResponse struct {
	Rotograph rotographInfoList `json:"rotograph"`
	VideoList videoInfoList     `json:"videoList"`
	// TODO 外包装
	MovieMessages []MovieMessageInfo `json:"movie_messages"`
}

// TODO
// 重写response
func (r *GetHomeInfoResponse) Response(rotographList *rotograph.List, videoList *video.VideosContributionList) {
	//处理轮播图
	rl := make(rotographInfoList, 0)
	for _, rk := range *rotographList {
		cover, _ := conversion.FormattingJsonSrc(rk.Cover)
		info := rotographInfo{
			Title: rk.Title,
			Cover: cover,
			Color: rk.Color,
			Type:  rk.Type,
			ToId:  rk.ToId,
		}
		rl = append(rl, info)
	}
	r.Rotograph = rl
	//处理视频
	vl := make(videoInfoList, 0)
	for _, lk := range *videoList {
		cover, _ := conversion.FormattingJsonSrc(lk.Cover)
		videoSrc, _ := conversion.FormattingJsonSrc(lk.Video)
		info := VideoInfo{
			ID:            lk.ID,
			Uid:           lk.Uid,
			Title:         lk.Title,
			Video:         videoSrc,
			Cover:         cover,
			VideoDuration: lk.VideoDuration,
			Label:         conversion.StringConversionMap(lk.Label),
			Introduce:     lk.Introduce,
			Heat:          lk.Heat,
			BarrageNumber: len(lk.Barrage),
			Username:      lk.UserInfo.Username,
			CreatedAt:     lk.CreatedAt,
		}
		vl = append(vl, info)
	}
	r.VideoList = vl

}
