package model

import (
	"5ub5i5t/bughuntingtoolbox/database"
	"html"
	"strings"

	"gorm.io/gorm"
)

type Domain struct {
	gorm.Model
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
