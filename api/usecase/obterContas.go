package usecase

import (
	"api/model"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var client = redis.NewClient(&redis.Options{
	Addr:	  "localhost:6379",
	Password: "",
	DB:		  0, 
	Protocol: 2, 
})

func ObterContas() []model.Conta {

	fmt.Printf("Obtendo contas...\n")

	contas := []model.Conta{}

	res, err := client.LRange(ctx, "contas", 0, -1).Result()

	if err != nil {
		fmt.Printf("\nFalha ao buscar contas no redis: %s\n", err)
		os.Exit(1)
	}

	for _, v := range res {
		var conta model.Conta
		err = json.Unmarshal([]byte(v), &conta)

		if (err != nil) {
			fmt.Println("Erro ao parsear a conta")
			continue
		}

		contas = append(contas, conta)
	}

	fmt.Println("Contas obtidas: ", contas)
	return contas
}


