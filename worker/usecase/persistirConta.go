package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"
	"worker/model"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var client = redis.NewClient(&redis.Options{
	Addr:	  "localhost:6379",
	Password: "",
	DB:		  0, 
	Protocol: 2, 
})

func PersistirConta(v model.ContaAvro) {

	fmt.Printf("\nSolicitação de conta: %s", v.Nome)

	simularDelayDaPersistencia()

	j, err :=  json.Marshal(v)

	if err != nil {
		fmt.Printf("Falha ao converter valor para json: %s\n", err)
		os.Exit(1)
	}

	err = client.LPush(ctx, "contas", j).Err()

	if (err != nil) {
		fmt.Printf("\nFalha ao persistir conta: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nConta persistida com sucesso: %s\n", v)
}

func simularDelayDaPersistencia() {
	time.Sleep(30 * time.Second) 
}
