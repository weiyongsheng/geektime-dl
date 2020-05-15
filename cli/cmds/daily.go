package cmds

import (
	"github.com/mmzou/geektime-dl/cli/application"
	"github.com/mmzou/geektime-dl/downloader"
	"github.com/urfave/cli"
)

//NewLessonCommand login command
func NewDailyCommand() []cli.Command {
	return []cli.Command{
		{
			Name:      "daily",
			Usage:     "搜索每日课程",
			UsageText: appName + " daily",
			Action:    dailyAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "phone",
					Usage:       "登录手机号",
					Destination: &LoginConfig.phone,
				},
				cli.StringFlag{
					Name:        "password",
					Usage:       "登录密码",
					Destination: &LoginConfig.password,
				},
				cli.StringFlag{
					Name:        "gcid",
					Usage:       "GCID Cookie",
					Destination: &LoginConfig.gcid,
				},
				cli.StringFlag{
					Name:        "gcess",
					Usage:       "GCESS Cookie",
					Destination: &LoginConfig.gcess,
				},
				cli.StringFlag{
					Name:        "serverId",
					Usage:       "SERVERID Cookie",
					Destination: &LoginConfig.serverID,
				},
			},
		},
	}
}

func dailyAction(c *cli.Context) error {
	lessons, err := application.SearchDailyLesson("mysql")

	if err != nil {
		return err
	}

	downloadData := downloader.Data{
		Title: "每日课程搜索“mysql”",
	}
	downloadData.Type = "每日课程"
	downloadData.Data = extractVideoDownloadData(lessons, 0)

	downloadData.PrintInfo()

	return nil
}
