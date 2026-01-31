package storage

import(
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct{
	Host string
	Port string
	Password string
	User string
	DBname string
	SSLmode string
}

func NewConnection(config *Config)(*gorm.DB,error){
	dsn: fmt.Sprintf(
		"host=%s port=%s password=%s dbname=%s sslmode=%s
		",config.Host,config.Port,config.Passowrd,config.DBname,config.SSLmode
	)
	db,err:=gorm.Open(postgres.Open(dsn),&form.COnfig{})

	if err!=nil{
		return db,err
	}
	return db,nil
}