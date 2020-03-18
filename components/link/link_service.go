package link

import (
	"log"

	"github.com/teris-io/shortid"
)

func saveLink(model *createLinkModel) error {
	log.Println(model)

	var err error
	model.Content, err = shortid.Generate()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
