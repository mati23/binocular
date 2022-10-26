package view

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/mati23/binocular/connectors"
	"github.com/mati23/binocular/domain"
)

type ContainerItem struct {
	DockerContainer     *types.Container
	DockerContainerView *fyne.Container
	StatusLabel         canvas.Text
	ActionButton        fyne.CanvasObject
	ActionButtonPointer *fyne.CanvasObject
	ControlContainer    fyne.Container
	DockerClient        *client.Client
}

func (containerItem *ContainerItem) CreateContainerItem() {

	containerItem.ControlContainer = *container.NewGridWithColumns(2)
	containerItem.StatusLabel = *canvas.NewText("Initial Text", color.Black)

	command := domain.RunnableCommand{Command: []string{"docker", "start", strings.Replace(containerItem.DockerContainer.Names[0], "/", "", -1)}}
	containerItem.ActionButton = CreateImageHeaderButtons(containerItem.updateContainerItem, command, "play-solid")
	containerItem.ActionButtonPointer = &containerItem.ActionButton

	containerItem.ControlContainer = *container.NewGridWithColumns(2, &containerItem.StatusLabel, containerItem.ActionButton)

	if containerItem.DockerContainer.State == "running" {
		fmt.Println("Container Running...")
		containerItem.buildRunningContainerView()

	} else {
		containerItem.buildStoppedContainerView()
	}

	containerItem.ControlContainer.Refresh()
	containerName := canvas.NewText(containerItem.DockerContainer.Names[0], color.White)
	containerName.TextStyle.Bold = true
	containerID := canvas.NewText(containerItem.DockerContainer.ID, color.White)
	arrowIcon := CreateIconByImageName("arrow-right", "30", "15")
	spacer := layout.NewSpacer()

	spacer.Resize(fyne.NewSize(0.0, 0.0))

	hbox := container.New(layout.NewHBoxLayout(), spacer, containerName, arrowIcon, containerID, layout.NewSpacer())
	containerItem.DockerContainerView = container.NewGridWithColumns(2, hbox, &containerItem.ControlContainer)
}

func (containerItem *ContainerItem) updateContainerItem() {
	fmt.Println("Content Being Updated...")

	if containerItem.DockerContainer.State == "running" {
		fmt.Println("Container Running...")
		containerItem.buildStoppedContainerView()
	} else {
		containerItem.buildRunningContainerView()
	}

	containerItem.DockerContainerView.Refresh()
}

func (containerItem *ContainerItem) buildRunningContainerView() {
	println("Running Container")
	containerItem.StatusLabel.Text = "Running"
	containerItem.StatusLabel.Color = color.RGBA{R: 0, G: 255, B: 1, A: 255}
	containerItem.StatusLabel.TextStyle.Bold = true

	command := domain.RunnableCommand{Command: []string{"docker", "stop", strings.Replace(containerItem.DockerContainer.Names[0], "/", "", -1)}}
	containerItem.ActionButton = CreateImageHeaderButtons(containerItem.updateContainerItem, command, "stop-solid")

	containerItem.ControlContainer.RemoveAll()
	containerItem.ControlContainer = *container.NewGridWithColumns(2, &containerItem.StatusLabel, CreateImageHeaderButtons(containerItem.updateContainerItem, command, "stop-solid"))
	containerItem.ControlContainer.Refresh()

	containerItem.DockerContainer = connectors.GetContainerById(containerItem.DockerContainer.ID, containerItem.DockerClient)
}
func (containerItem *ContainerItem) buildStoppedContainerView() {
	println("Stopping Container")
	containerItem.StatusLabel.Text = "Stopped"
	containerItem.StatusLabel.Color = color.RGBA{R: 255, G: 1, B: 1, A: 255}
	containerItem.StatusLabel.TextStyle.Bold = true

	command := domain.RunnableCommand{Command: []string{"docker", "start", strings.Replace(containerItem.DockerContainer.Names[0], "/", "", -1)}}

	containerItem.ControlContainer.RemoveAll()
	containerItem.ControlContainer = *container.NewGridWithColumns(2, &containerItem.StatusLabel, CreateImageHeaderButtons(containerItem.updateContainerItem, command, "play-solid"))
	containerItem.ControlContainer.Refresh()

	containerItem.DockerContainer = connectors.GetContainerById(containerItem.DockerContainer.ID, containerItem.DockerClient)
}
