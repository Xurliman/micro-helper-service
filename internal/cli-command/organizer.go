package cli_command

import (
	"flag"
	"github.com/Xurliman/banking-microservice/internal/migrations"
	seeds "github.com/Xurliman/banking-microservice/internal/seeders"
	"gorm.io/gorm"
)

func Organize(db *gorm.DB) {
	flag.Parse()
	arguments := flag.Args()

	if len(arguments) >= 1 {
		switch arguments[0] {
		case "seed":
			seeds.Execute(db, arguments[1:]...)
		case "migrate":
			migrations.Execute(db)
		}
	}
}
