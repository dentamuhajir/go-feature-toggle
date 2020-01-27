package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type ToggleModel struct {
	gorm.Model
	Key string `gorm:"size:50"`
	On  bool
}

type DatabaseConfig struct {
	Provider, Name string
}

var (
	db                 *gorm.DB
	err                error
	dbProvider, dbName string = "sqlite3", "feature-toggle.db"
)

func (db *DatabaseConfig) configure(p, n string) {
	db.Provider = p
	db.Name = n
}

func Init() (db *gorm.DB, err error) {
	cfg := DatabaseConfig{}
	cfg.configure(dbProvider, dbName)
	db, err = gorm.Open(cfg.Provider, cfg.Name)
	return
}

func Close() error {
	return db.Close()
}

func Migration(db *gorm.DB) (bool, float64) {

	if db.HasTable("toggle_models") {
		db.DropTable("toggle_models")
	}

	db.AutoMigrate(&ToggleModel{})

	var toggles []ToggleModel = []ToggleModel{
		ToggleModel{Key: "feature-new-style-v2", On: true},
		ToggleModel{Key: "feature-show-verified-article", On: false},
		ToggleModel{Key: "feature-show-sidebar", On: false},
	}

	totalMigrate := 3

	for _, toggle := range toggles {
		db.Create(&toggle)
	}

	return true, float64(totalMigrate)
}

// func FeatureIsOn(keyName string) bool {
// 	db = DBConnect()

// 	var toggles ToggleModel

// 	if db.Where("key = ?", keyName).Find(&toggles).RecordNotFound() {
// 		panic("Record not found")
// 	}

// 	return toggles.On
// }

type Feature struct {
	Key string
}

func FeatureIsOn(f Feature) bool {
	keyName := f.Key
	db, _ = Init()

	var toggles ToggleModel

	if db.Where("key = ?", keyName).Find(&toggles).RecordNotFound() {
		panic("Record not found")
	}

	return toggles.On

	//return f.Key
}
