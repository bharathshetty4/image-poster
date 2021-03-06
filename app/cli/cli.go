package cli

import (
	"fmt"
	"log"
	"os"

	cliV2 "github.com/urfave/cli/v2"
	"go.uber.org/zap"

	commands "github.com/mayflyman4/image-poster/app/cli/commands"
)

// TODO: get the version from a File
const (
	cliVersion = "0.0.1"
)

func StartCLI() {
	app := cliV2.NewApp()

	//set cli properties
	setAppProperties(app)

	//add all the commands
	addCommands(app)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("Unable to start the application, ", zap.Error(err))
	}
}

//SetAppProperties adds CLI properties
func setAppProperties(app *cliV2.App) {
	app.Name = "Image Poster"
	app.HelpName = "image-poster"
	app.Usage = "CLI application to post an Image to multiple Platforms like Flickr, 500px etc simultaneously."
	app.UsageText = "image-poster command [command options]"

	app.Description = "A CLI utility to upload the given Image file to some of the platforms like Flickr, 500px etc" +
		" with \n   the given tags using a single command."
	app.CustomAppHelpTemplate = "Image-poster "
	app.Version = cliVersion
	//TODO: BKS autocompletion of the app is not happening
	app.EnableBashCompletion = true
	app.CustomAppHelpTemplate = fmt.Sprintf("%s \nMore information at: https://github.com/mayflyman4/image-poster\n", cliV2.AppHelpTemplate)

	//Author related
	authtorBks := cliV2.Author{Name: "Bharath Kumar", Email: "shettybharath4@gmail.com"}
	app.Authors = append(app.Authors, &authtorBks)
}

//addCommands add commands list to the CLI commands
func addCommands(app *cliV2.App) {
	app.Commands = []*cliV2.Command{
		// commands.AddCommandStartServer(),
		commands.AddCommandAddPlatform(),
	}
}
