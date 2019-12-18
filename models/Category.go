package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Category Entity
type Category struct {
	gorm.Model
	ID      int    `gorm:"unique;AUTO_INCREMENT"`
	Name    string `gorm:"type:varchar(MAX)"`
	Created time.Time
	Updated time.Time
}
