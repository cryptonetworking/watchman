package watchman

import (
	"fmt"
	"github.com/itsabgr/go-handy"
	"gorm.io/gorm"
	"time"
)

type model struct {
	gorm.Model
	Time     time.Time `gorm:"index"`
	Instance string    `gorm:"index"`
	Log      string
}
type _gorm struct {
	db *gorm.DB
}

func (g _gorm) Close() error {
	return nil
}

func (g _gorm) Put(date time.Time, instance string, data ...any) error {
	return g.db.Create(&model{
		Time:     date,
		Instance: instance,
		Log:      fmt.Sprint(data...),
	}).Error
}

func Gorm(dialect gorm.Dialector) Output {
	db, err := gorm.Open(dialect)
	handy.Throw(err)
	err = db.AutoMigrate(&model{})
	handy.Throw(err)
	return _gorm{
		db: db,
	}
}
