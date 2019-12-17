package entities

import (
	"time"

	"github.com/jinzhu/gorm"
)

//NumberTicket Entity
type NumberTicket struct {
	gorm.Model
	ID       int `gorm:"unique;AUTO_INCREMENT"`
	Number   int
	Date     time.Time
	Used     bool
	Category Category
}
