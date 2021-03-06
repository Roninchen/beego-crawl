package controllers

import (
	"crawl_movie/models"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/logs"
	"time"
)

type CrawlMovieController struct {
	beego.Controller
}

func (c *CrawlMovieController) CrawlMovie() {

	//连接到redis
	models.ConnectRedis("redis服务器不告诉你嘿嘿嘿")

	//爬虫入口url
	sUrl := "https://movie.douban.com/subject/27133303/"
	models.PutinQueue(sUrl)
	go models.IP66()
	for {
		//models.IP66()
		length := models.GetQueueLength()
		if length == 0 {
			break //如果url队列为空 则退出当前循环
		}

		sUrl = models.PopfromQueue()
		//我们应当判断sUrl是否应该被访问过
		if models.IsVisit(sUrl) {
			continue
		}
		logs.SetLogger(sUrl)
		go models.Run(sUrl)
		time.Sleep(2000)
	}
}