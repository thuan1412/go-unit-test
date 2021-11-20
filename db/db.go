package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
    gorm.Model
    Name string
    Price uint
}

func insert(db *gorm.DB, product *Product) error  {
    return db.Create(product).Error
}

func main() {
    db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })

    err = db.Unscoped().Where("true").Delete(&Product{}).Error
    if err != nil {
        panic(err)
    }
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&Product{})
    db.Create(&Product{Name: "test-1", Price: 110})
    db.Create(&Product{Name: "test-2", Price: 120})
    db.Create(&Product{Name: "test-3", Price: 130})
    db.Create(&Product{Name: "test-4", Price: 140})
    db.Create(&Product{Name: "test-5", Price: 150})
    db.Create(&Product{Name: "test-6", Price: 160})
}
