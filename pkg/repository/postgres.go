package repository

import (
	"fmt"
	"github.com/LuxAeterna-git/my-parcer/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Pg struct {
	db *gorm.DB
}

func NewPg() *Pg {
	dsn := "host=localhost user=user password=test dbname=goods port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,  // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})
	if err != nil {
		fmt.Errorf("err while initilaizing db: %s", err)
	}

	db.AutoMigrate(&model.Good{})
	return &Pg{db: db}
}

func (s *Pg) Store(product model.Good) {
	s.db.Create(&product)
}

func (s *Pg) FindAll() []model.Good {
	var products []model.Good
	s.db.Find(&products)
	return products
}

func (s *Pg) DeleteByID(id int) {
	s.db.Delete(&model.Good{}, id)
}
