package domain

import (
	"fmt"
	"strconv"
)

type Image struct {
	Prefix     string
	ID         string
	Repository string
	Tag        string
	Created    string
	Size       int64
}

func (image Image) PrintImageStats() {
	fmt.Printf("Image stats: %s-%s-%s-%s-%s-%s\n",
		image.Prefix, image.ID, image.Repository, image.Tag, image.Created, strconv.FormatInt(image.Size, 10))
}

func (image Image) ImageKey() string {
	return image.Prefix + "-" + image.ID
}
