package main

import (
	"fyne.io/fyne/v2"
	"github.com/mati23/binocular/connectors"
	"github.com/mati23/binocular/view"
)

var DEFAULT_WIDTH float32 = 1600.0
var DEFAULT_HEIGHT float32 = 1000.0

func main() {
	windowSize := fyne.Size{Width: DEFAULT_WIDTH, Height: DEFAULT_HEIGHT}
	dockerClient := connectors.DockerClient()

	view.MainView(windowSize, *dockerClient)

}
