package models

import (
	"log"
	"strings"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(data ConfigFile) {

	var builder strings.Builder
	builder.WriteString(data.DataBase.IP)
	builder.WriteString(data.DataBase.PORT)
	builder.WriteString(data.DataBase.DB)
	urlDB := builder.String()

	opt, err := redis.ParseURL(urlDB)
	if err != nil {
		log.Fatalf("Error al conectarse a la base de datos: %s", err)
	}

	// client := redis.NewClient(opt)
	redis.NewClient(opt)

}
