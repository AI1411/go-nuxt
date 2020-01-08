package api

import (
	"context"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"os"
)

func FetchMostPopularVideos() echo.HandlerFunc {
	return func(c echo.Context) error {

		key := os.Getenv("API_KEY")

		ctx := context.Background()
		//起動の条件
		yts, err := youtube.NewService(ctx, option.WithAPIKey(key))
		if err != nil {
			logrus.Fatalf("Error creating new youtube service: %v", err)
		}

		call := yts.Videos.List("id,snippet").Chart("mostPopular").MaxResults(3)
		//続きの動画を取得するためにpageTokenをリクエスト時に送信する
		pageToken := c.QueryParam("pageToken")
		if len(pageToken) > 0 {
			call = call.PageToken(pageToken)
		}

		//youtube APIの実行
		res, err := call.Do()
		if err != nil {
			logrus.Fatalf("Error calling Youtube API: %v", err)
		}
		return c.JSON(fasthttp.StatusOK, res)
	}
}
