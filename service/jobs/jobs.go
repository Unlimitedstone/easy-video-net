package jobs

import (
	"easy-video-net/jobs/downloadpic"
	"fmt"
	"github.com/robfig/cron/v3"
)

type Jobs interface {
	Start()
	GetCron() string
	Name() string
}

var JobsGetter = func() []Jobs {
	return []Jobs{
		&downloadpic.DownloadPic{Cron: "/xxxxxx"},
	}
}

func Run() {
	jobList := JobsGetter()
	for _, job := range jobList {
		go func(job Jobs) {
			c := cron.New(cron.WithSeconds())
			_, err := c.AddFunc(job.GetCron(), func() {
				job.Start()
			})
			if err != nil {
				fmt.Println(fmt.Sprintf("Job: %s, run error: %s", job.Name(), err.Error()))
			}
		}(job)
	}
}
