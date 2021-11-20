package main

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("main.db"), &gorm.Config{})

    // drop if exists
    db.Migrator().DropTable(&Product{})
    db.AutoMigrate(&Product{})

	if err != nil {
		panic("failed to connect database")
	}
}

func Test_insert(t *testing.T) {
	tests := []struct {
		name    string
		product Product
	}{
		{
			name: "insert product 1",
			product: Product{
				Name:  "product1",
				Price: 100,
			},
		},
		{
			name: "insert product 2",
			product: Product{
				Name:  "product2",
				Price: 200,
			},
		},
		{
			name: "insert product 3",
			product: Product{
				Name:  "product3",
				Price: 200,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            err := insert(db, &tt.product)
            if err != nil {
                t.Errorf("insert product error: %v", err)
            }
			if tt.product.ID == uint(0) {
				t.Errorf("insert product failed: product id is %d", tt.product.ID)
			}
		})
	}
    var total int64
    db.Where("true").Find(&Product{}).Count(&total)
    if total != 3 {
        t.Errorf("insert product failed: total is %d", total)
    }
}
