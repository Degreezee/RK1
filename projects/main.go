package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Input struct {
	Quantity *int `json:"quantity"`
}

type Output struct {
	Result string `json:"result"`
}

// Обработчик HTTP-запроса
func QuantityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write([]byte("method not allowed"))
		return
	}

	var input Input

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	if input.Quantity == nil {
		w.WriteHeader(400)
		w.Write([]byte("quantity is missing"))
		return
	}

	var output Output

	switch {
	case *input.Quantity < 1:
		output.Result = "Гости не приехали"
	case *input.Quantity == 1:
		output.Result = "стандарт"
	case *input.Quantity == 2:
		output.Result = "комфорт"
	case *input.Quantity == 3:
		output.Result = "бизнес"
	case *input.Quantity == 4:
		output.Result = "люкс"
	case *input.Quantity > 4:
		output.Result = "Необходимо несколько номеров"
	default:
		w.WriteHeader(400)
		w.Write([]byte("unknown operator"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	respBytes, _ := json.Marshal(output)
	w.Write(respBytes)
}

func main() {
	http.HandleFunc("/room_type", QuantityHandler)

	fmt.Println("starting server...")
	err := http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
