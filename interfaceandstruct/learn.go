package interfaceandstruct

import "fmt"

type ICitizenService interface {
	Validate(c *Citizen) bool
	CreateCitizenCard(c *Citizen) error
}

type ThaiCitizenService struct {

}

func NewThaiCitizenService() *ThaiCitizenService {
	return &ThaiCitizenService{}
}

func (svc *ThaiCitizenService) Validate(c *Citizen) bool {
	return len(c.Firstname) > 0 && len(c.Lastname) > 0 && len(c.CitizenID) > 0
}

func (svc *ThaiCitizenService) CreateCitizenCard(c *Citizen) error {
	println(fmt.Sprintf("successfully create Thai citizen id for ID=%s", c.CitizenID))
	return nil
}

type JapanCitizenService struct {
	
}

func NewJapanCitizenService() *JapanCitizenService {
	return &JapanCitizenService{}
}

func (svc *JapanCitizenService) Validate(c *Citizen) bool {
	return len(c.Firstname) > 0 && len(c.Lastname) > 0 && len(c.CitizenID) > 0
}

func (svc *JapanCitizenService) CreateCitizenCard(c *Citizen) error {
	println(fmt.Sprintf("successfully create Japan citizen id for ID=%s", c.CitizenID))
	return nil
}

func Learn() {
	citizen := NewCitizen("Narongsak", "Suwan", "1122334455")
	citizenSvc := NewThaiCitizenService()

	err := createCitizenCard(citizenSvc, citizen)
	if err != nil {
		printError(err)
	}
}

func createCitizenCard(svc ICitizenService, c *Citizen) error {
	if !svc.Validate(c) {
		return fmt.Errorf("Citizen data is valid")
	}

	err := svc.CreateCitizenCard(c)
	if err != nil {
		return err
	}
	return nil
}

type Citizen struct {
	Firstname	string	`json:"firstname"`
	Lastname	string	`json:"lastname"`
	CitizenID	string	`json:"citizenId"`
}

func NewCitizen(firstname, lastname, citizenID string) *Citizen {
	return &Citizen{
		Firstname: firstname,
		Lastname: lastname,
		CitizenID: citizenID,
	}
}

func printError(err error) {
	fmt.Println("error = ", err.Error())
}