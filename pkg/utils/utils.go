package utils

import (
	"encoding/json"
	"errors"
	"github.com/altairtt/pkg/storage"
)

type ToolsI interface {
	InnerStruct(hombre []storage.Male, mujer []storage.Female) (inner []interface{}, err error)
	CreateTeam(inner []interface{}) (newTeam []storage.Team, err error)
}

type Tools struct{}

func NewTools() *Tools {
	t := new(Tools)

	return t
}

//InnerStruct. Method to join Male and Female array structs
func (t *Tools) InnerStruct(males []storage.Male, females []storage.Female) (inner []interface{}, err error) {
	nMales := len(males)
	nFemales := len(females)

	nTotal := nMales + nFemales
	if nTotal == 0 {
		err = errors.New("Structs arrays are empty. Impossible to follow up")
		return inner, err
	}

	inner = make([]interface{}, 0, nTotal)

	for i := range males {
		inner = append(inner, males[i])
	}
	for i := range females {
		inner = append(inner, females[i])
	}

	return
}

//CreateTeam. Method to transform an interface array to a People array
func (t *Tools) CreateTeam(inner []interface{}) (newTeam []storage.Team, err error) {
	newTeam = []storage.Team{}

	data, _ := json.Marshal(inner)

	err = json.Unmarshal(data, &newTeam)
	if err != nil {
		return
	}

	return
}
