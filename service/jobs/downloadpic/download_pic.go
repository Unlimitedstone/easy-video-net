package downloadpic

import (
	"easy-video-net/global"
	"easy-video-net/models/query"
	"fmt"
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm/clause"
	"path/filepath"
)

type DownloadPic struct {
	Cron string
}

func DownloadPicToFile(url, path, name string) error {
	client := resty.New()
	resp, err := client.R().SetOutput(filepath.Join(path, name)).Get(url)
	if err != nil {
		return fmt.Errorf("client error %s", err.Error())
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("bad status: %s", resp.Status())
	}
	return nil
}

const filepathPrefix = "assets/static/video/%s.jpg"

func (d *DownloadPic) Start() {
	db := query.Use(global.Db)
	movieMessageList, err := db.MovieMessageModel.Debug().Where(
		db.MovieMessageModel.ImageURL.IsNotNull()).Where(db.MovieMessageModel.ImageSrc.IsNull()).Find()
	if err != nil {
		fmt.Println(fmt.Sprintf("query movieMessageModel error: %v", err))
	}

	count := len(movieMessageList)
	successCount := 0
	for _, movieMessage := range movieMessageList {
		imageSrc := fmt.Sprintf(filepathPrefix, movieMessage.UniID)
		if err := DownloadPicToFile(movieMessage.ImageURL, "./", imageSrc); err != nil {
			fmt.Println(fmt.Sprintf("download movieMessageModel error: %v", err))
			continue
		}
		movieMessage.ImageSrc = imageSrc
		successCount++
	}

	if err := query.Use(global.Db).MovieMessageModel.Debug().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"image_src"}),
	}).Create(movieMessageList...); err != nil {
		fmt.Println(fmt.Sprintf("update movieMessageModel error: %v", err))
	}

	fmt.Println(fmt.Sprintf("success count: %d, all count: %d", successCount, count))
}

func (d *DownloadPic) GetCron() string {
	return d.Cron
}

func (d *DownloadPic) Name() string {
	return "download_pic"
}
