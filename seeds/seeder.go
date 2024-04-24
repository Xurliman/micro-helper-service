package seeds

import (
	"gorm.io/gorm"
	"log"
	"reflect"
)

type Seed struct {
	db *gorm.DB
}

func seed(seed Seed, seedMethodName string) {
	method := reflect.ValueOf(seed).MethodByName(seedMethodName)

	if !method.IsValid() {
		log.Fatalf("%s is not a valid function", seedMethodName)
	}
	log.Printf("Seeding %s ...", seedMethodName)
	method.Call(nil)
	log.Printf("Seeding %s succeeded", seedMethodName)
}

func Execute(db *gorm.DB, seedMethodNames ...string) {
	seedStruct := Seed{db: db}
	seedType := reflect.TypeOf(seed)

	if len(seedMethodNames) == 0 {
		log.Printf("Seeding all seeders")
		for i := 0; i < seedType.NumMethod(); i++ {
			method := seedType.Method(i)
			log.Printf("Seeding %s", method.Name)
			seed(seedStruct, method.Name)
		}
	}

	for _, seedMethodName := range seedMethodNames {
		seed(seedStruct, seedMethodName)
	}
}
