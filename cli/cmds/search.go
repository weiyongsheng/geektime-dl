package cmds

import (
	"errors"

	"github.com/mmzou/geektime-dl/cli/application"
	"github.com/mmzou/geektime-dl/downloader"
	"github.com/urfave/cli"
)

//NewSearchCommand login command
func NewSearchCommand() []cli.Command {
	return []cli.Command{
		{
			Name:      "ds",
			Usage:     "搜索每日课程",
			UsageText: appName + " ds [搜索词]",
			Action:    searchAction,
		},
	}
}

func searchAction(c *cli.Context) error {
	if c.NArg() == 0 {
		cli.ShowSubcommandHelp(c)
		println()
		return errors.New("请输入【搜索词】")
	}

	kwd := c.Args().Get(0)
	lessons, err := application.SearchDailyLesson(kwd)

	if err != nil {
		return err
	}

	downloadData := downloader.Data{
		Title: "每日课程搜索词“" + kwd + "”",
	}
	downloadData.Type = "每日课程"
	downloadData.Data = extractVideoDownloadData(lessons, 0)

	// printExtractDownloadData(lessons)

	downloadData.PrintInfo()

	return nil
}
