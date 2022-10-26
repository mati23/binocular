package view

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/mati23/binocular/domain"
)

func CreateImageHeaderButtons(f func(), command domain.RunnableCommand, iconName ...string) fyne.CanvasObject {
	bytes, err := ioutil.ReadFile("resources/icons/" + iconName[0] + ".png")
	if err != nil {
		panic(err)
	}

	resource := fyne.NewStaticResource(
		"icon",
		bytes,
	)
	var button = widget.NewButtonWithIcon("", resource, func() {
		print("Tapped icon button\n")
		runCommandWithTerminalOutput(command.Command...)
		if f != nil {
			f()
		}
	})
	if len(iconName) > 1 {
		button = widget.NewButtonWithIcon(iconName[1], resource, func() {
			print("Tapped icon button")
		})
	}

	return button
}

func CreateIconByImageName(imageName ...string) fyne.CanvasObject {
	bytes, err := ioutil.ReadFile("resources/icons/" + imageName[0] + ".png")
	if err != nil {
		panic(err)
	}

	resource := fyne.NewStaticResource(
		"icon",
		bytes,
	)
	image := canvas.NewImageFromResource(resource)
	if len(imageName) > 1 {
		imageWidth, err := strconv.Atoi(imageName[1])
		if err != nil {
			panic(err)
		}
		imageHeight, err := strconv.Atoi(imageName[2])
		if err != nil {
			panic(err)
		}
		image.SetMinSize(fyne.NewSize(float32(imageWidth), float32(imageHeight)))
	}

	return image
}
