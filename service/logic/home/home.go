package home

import (
	"easy-video-net/global"
	receive "easy-video-net/interaction/receive/home"
	response "easy-video-net/interaction/response/home"
	"easy-video-net/models/contribution/video"
	"easy-video-net/models/home/rotograph"
	"easy-video-net/models/model"
	"easy-video-net/models/query"
	"fmt"
)

func GetHomeInfo(data *receive.GetHomeInfoReceiveStruct) (results interface{}, err error) {
	//获取主页轮播图
	rotographList := new(rotograph.List)
	err = rotographList.GetAll()
	if err != nil {
		return nil, err
	}

	//获取主页推荐视频
	videoList := new(video.VideosContributionList)
	err = videoList.GetHoneVideoList(data.PageInfo)

	if err != nil {
		return nil, err
	}
	res := &response.GetHomeInfoResponse{}
	res.Response(rotographList, videoList)

	movieMessages, err := query.Use(global.Db).MovieMessageModel.Find()
	if err != nil {
		return nil, err
	}

	return buildResp(res, movieMessages), nil
}

func buildResp(resp *response.GetHomeInfoResponse, movieMessages []*model.MovieMessageModel) *response.GetHomeInfoResponse {
	movieMessageInfo := make([]response.MovieMessageInfo, 0, len(movieMessages))
	for _, v := range movieMessages {
		movieMessageInfo = append(movieMessageInfo, response.MovieMessageInfo{
			ID:          v.ID,
			UniID:       v.UniID,
			Name:        v.Name,
			ImageURL:    v.ImageURL,
			ImageSrc:    fmt.Sprintf("%s/%s", global.Config.ProjectUrl, v.ImageSrc),
			Score:       v.Score,
			Type:        v.Type,
			Common:      v.Common,
			Source:      v.Source,
			CreateTime:  v.CreateTime,
			UpdateTime:  v.UpdateTime,
			Introduce:   v.Introduce,
			Alias:       v.Alias_,
			Duration:    v.Duration,
			ReleaseDate: v.ReleaseDate,
			Director:    v.Director,
			Actor:       v.Actor,
			Writer:      v.Writer,
			Country:     v.Country,
			Imdb:        v.Imdb,
		})
	}
	resp.MovieMessages = movieMessageInfo
	return resp
}
