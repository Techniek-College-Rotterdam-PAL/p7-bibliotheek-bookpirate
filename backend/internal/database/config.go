package database

import (
	"errors"
	"github.com/xo/dburl"
	"gorm.io/gorm"
)

var (
	siteDb *gorm.DB
	bookDb *gorm.DB
)

func Site() *gorm.DB  { return siteDb }
func Books() *gorm.DB { return bookDb }

type LoadDbCfg struct {
	SourceDsns []*dburl.URL

	//LogEnableStackTrace bool
	//LogPrefix           string
	//LogFile             *os.File
}

func checkConfig(cfg *LoadDbCfg) error {
	switch {
	case cfg == nil:
		return errors.New("invalid configuration: config was nil")
	case len(cfg.SourceDsns) == 0:
		return errors.New("invalid configuration: no source dsns provided")
	default:
		return nil
	}
}
