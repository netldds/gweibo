package services

import (
	"gweibo/models"
	"net/url"
	"time"
)

//ROOT https://weibo.com/p/
var (
	HomePageReq = &models.WeiRequests{
		Req: []models.WeiRequest{
			{
				Method:  models.GET,
				Params:  []string{models.PID, "home"},
				Timeout: time.Second * 10,
				Query: url.Values{
					//"stat_date":   []string{"201907"},
					"pids":        []string{"Pl_Official_MyProfileFeed__20"},
					"ajaxpagelet": []string{"1"},
				},
			},
			{
				Method:  models.GET,
				Params:  []string{"aj", "mblog", "getlongtext"},
				Timeout: time.Second * 10,
				Query: url.Values{
					"mid": []string{""},
				},
			},
		},
	}
)
