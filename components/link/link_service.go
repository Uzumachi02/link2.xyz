package link

import (
	"errors"
	"log"

	"github.com/Uzumachi02/link2.xyz/store/postgresql"
	"github.com/teris-io/shortid"
)

func saveLink(model *createLinkModel) (string, error) {
	log.Println(model)

	hasTargetURL := len(model.TargetURL) != 0
	hasContent := len(model.Content) != 0

	if !hasTargetURL && !hasContent {
		return "", errors.New("is empty")
	}

	uniqid, err := getUniqID()
	if err != nil {
		return "", err
	}

	return uniqid, nil
}

func getUniqID() (string, error) {
	linkRepo := postgresql.NewLinkRepository()
	repeat := 1

	for {
		uniqid, err := shortid.Generate()
		if err != nil {
			return "", err
		}

		has, err := linkRepo.HasUniqid(uniqid)
		if err != nil {
			return "", err
		}

		if !has {
			return uniqid, nil
		}

		if repeat > 50 {
			return "", errors.New("not uniqid")
		}

		repeat++
	}
}
