/*
 * This file is part of the XXX distribution (https://github.com/xxxx or http://xxx.github.io).
 * Copyright (c) 2015 Mateus Arruda de Medeiros.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 3.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */
package view

import (
	"bytes"
	"container/list"
	"encoding/base64"
	"fmt"
	"log"
	"os/exec"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/mati23/binocular/connectors"
	"github.com/mati23/binocular/domain"
)

var images = list.New()

func AddImageToList(image domain.Image) {
	images.PushFront(image)
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func fillGrid(dockerImages []types.ImageSummary, dockerClient client.Client) []*widget.AccordionItem {
	accordionItems := []*widget.AccordionItem{}

	for _, dockerImage := range dockerImages {
		image := domain.Image{
			Prefix:     "image",
			ID:         dockerImage.ID,
			Repository: dockerImage.RepoDigests[0],
			Tag:        dockerImage.RepoTags[0],
			Created:    strconv.FormatInt(dockerImage.Created, 10),
			Size:       dockerImage.Size,
		}

		filter := filters.NewArgs(filters.KeyValuePair{"ancestor", image.ID})
		containerLabels := []types.Container{}
		dockerContainers := connectors.GetContainersList(&dockerClient, filter)
		println("Printing containers\n")
		if len(dockerContainers) > 0 {
			containerLabels = append(containerLabels, dockerContainers...)
		}

		command := domain.RunnableCommand{Command: []string{"docker", "run", image.ID}}
		headderButton := CreateImageHeaderButtons(nil, command, "play-solid", "Run")

		var accordionItem = widget.NewAccordionItem("",
			container.NewGridWithRows(1, headderButton, widget.NewLabel(image.ID)))
		if len(containerLabels) > 0 {
			rowContainerLabels := buildRowContainerWithLabels(containerLabels, dockerClient)

			accordionItem = widget.NewAccordionItem(dockerImage.ID,
				container.NewGridWithRows(2, headderButton, rowContainerLabels))
			accordionItems = append(accordionItems, accordionItem)

		} else {
			accordionItem = widget.NewAccordionItem(dockerImage.ID,
				container.NewGridWithRows(1, headderButton))
			accordionItems = append(accordionItems, accordionItem)
		}

	}

	return accordionItems
}

func buildRowContainerWithLabels(containerLabels []types.Container, dockerClient client.Client) *fyne.Container {

	var containerItems = []ContainerItem{}
	for _, containerItem := range containerLabels {
		container := &containerItem
		containerItemView := ContainerItem{DockerContainer: container, DockerClient: &dockerClient}
		containerItemView.CreateContainerItem()
		containerItems = append(containerItems, containerItemView)
	}

	var containerViews = []fyne.CanvasObject{}
	for _, containerItem := range containerItems {
		if containerItem.DockerContainerView != nil {
			containerViews = append(containerViews, containerItem.DockerContainerView)
		}

	}

	rowContainer := container.NewGridWithRows(len(containerViews), containerViews...)
	return rowContainer
}

func runCommandWithTerminalOutput(command ...string) {
	cmd := exec.Command(command[0], command[1], command[2])
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("out: ", outb.String())

}

func MainView(windowSize fyne.Size, dockerClient client.Client) {
	dockerImages := connectors.GetImageList(&dockerClient)

	app := app.New()
	window := app.NewWindow("Docker Gui")
	window.SetFixedSize(true)
	window.Resize(windowSize)

	imagesGrid := container.New(layout.NewGridLayout(1))
	list := fillGrid(dockerImages, dockerClient)
	accordion := widget.NewAccordion(list...)

	imagesGrid.Add(accordion)

	tabs := container.NewAppTabs(
		container.NewTabItem("Images", imagesGrid),
		container.NewTabItem("Containers", widget.NewLabel("Container")),
		container.NewTabItem("Volumes", widget.NewLabel("Volumes")),
	)
	window.SetContent(tabs)

	fmt.Printf("Starting Excecution\n")
	window.ShowAndRun()
	fmt.Printf("Execution Finished\n")
}
