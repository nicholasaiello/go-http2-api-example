package entity

type Account struct {
	ID          string `json:"accountId"`
	EncodedID   string `json:"encAccountId"`
	Name        string `json:"accountLongName"`
	ShortName   string `json:"accountShortName"`
	Description string `json:"acctDesc"`
	Mode        string `json:"accountMode"`
	Status      string `json:"acctStatus"`
	Type        string `json:"acctType"`
	Created     string `json:"acctCreationDate"`
}
