package UserHandler

import "time"

type Auth struct {
	Id            int     `gorm:"primary_key";"AUTO_INCREMENT"`
	LastLoginTime time.Time
	ProfileId     int
}
