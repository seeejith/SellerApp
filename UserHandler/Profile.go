package UserHandler

type Profile struct {
	Id           int    `gorm:"primary_key";"AUTO_INCREMENT",json:",omitempty"`
	Name         string `json:",omitempty"`
	Age          int    `json:",omitempty"`
	Address      string `json:",omitempty"`
	MobileNumber string `json:",omitempty"`
	Email        string `json:",omitempty"`
}
