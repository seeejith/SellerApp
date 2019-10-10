package UserHandler

type Access struct {
	Id        int     `gorm:"primary_key";"AUTO_INCREMENT"`
	Read      bool    `sql:"default:true"`
	Write     bool    `sql:"default:true"`
	ProfileId int
}
