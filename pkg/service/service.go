package service

import (
	"github.com/altairtt/pkg/storage"
	"github.com/altairtt/pkg/utils"
	"log"
	"sort"
)

type ServiceI interface {
	Run() error
}

type Service struct {
	Tools *utils.Tools
}

func NewService(tools *utils.Tools) *Service {
	s := new(Service)

	s.Tools = tools

	return s

}

//Run. Method to starts the service
func (s *Service) Run() error {

	//Step 1. Inner two differents structs
	people, err := s.Tools.InnerStruct(storage.Males, storage.Females)
	if err != nil {
		return err
	}

	//Step 2. Create new team
	newTeam, err := s.Tools.CreateTeam(people)
	if err != nil {
		return err
	}

	//Step 3. Order the team by age
	sort.SliceStable(newTeam, func(i, j int) bool {
		return newTeam[i].Age < newTeam[j].Age
	})

	log.SetFlags(0)
	log.Println("Team components sorted by age:")

	for _, p := range newTeam {
		log.Printf("%s %s %v",p.Name, p.Surname, p.Age)
	}

	return nil
}
