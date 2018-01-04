package main
import (
	"os"

	"github.com/urfave/cli"
)


func main() {
	app := cli.NewApp()
	app.Name = "transcoder"
	app.Usage = "transcoder is the daemon that converts media files into streaming files."
	app.Action = func(c *cli.Context) error {
		//TODO: Run as daemon service
		return nil
	}

	app.Run(os.Args)
}