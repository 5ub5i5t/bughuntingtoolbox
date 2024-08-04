package model

import (
	"5ub5i5t/bughuntingtoolbox/database"
	"html"
	"strings"

	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
	//ID uint `gorm:"primaryKey"`
	//Target string `gorm:"type:text" json:"target"`
	//Domain string `gorm:"type:text" json:"domain"`
	//Target string `gorm:"UNIQUE_INDEX:compositeindex;type:text;not null;json:target"`
	//Domain string `gorm:"UNIQUE_INDEX:compositeindex;type:text;not null;json:domain"`
	//Target string `gorm:"UNIQUE_INDEX:compositeindex;type:text;not null" json:"target"`
	//Domain string `gorm:"UNIQUE_INDEX:compositeindex;type:text;not null" json:"domain"`
	// "index:idx_name,unique"
	Target string `gorm:"index:idx_unique,unique" type:"text;not null" json:"target"`
	Domain string `gorm:"index:idx_unique,unique" type:"text;not null" json:"domain"`
}

func (domain *Domain) Save() (*Domain, error) {
	err := database.Database.Create(&domain).Error
	if err != nil {
		return &Domain{}, err
	}
	return domain, nil
}

func (domain *Domain) BeforeSave(*gorm.DB) error {
	domain.Target = html.EscapeString(strings.TrimSpace(domain.Target))
	domain.Domain = html.EscapeString(strings.TrimSpace(domain.Domain))
	return nil
}
