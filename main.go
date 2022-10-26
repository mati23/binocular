package main

import (
	"fyne.io/fyne/v2"
	"github.com/mati23/binocular/connectors"
	"github.com/mati23/binocular/view"
)

func main() {
	windowSize := fyne.Size{1600.0, 1000.0}
	dockerClient := connectors.DockerClient()

	view.MainView(windowSize, *dockerClient)

}
