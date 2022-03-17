package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct {
	// Name       string  `json:"name,omitempty"` //空字段不輸出
	ID string `json:"id"`
	// Item       OrderItem `json:"item"`
	Item       []OrderItem `json:"items"`
	Quantity   int         `json:"quantity"`
	TotalPrice float64     `json:"total_price"`
}

func marshal() {
	o := Order{
		ID:         "1234",
		Quantity:   3,
		TotalPrice: 20,
		Item: []OrderItem{
			{
				ID:    "item_1",
				Name:  "learn go",
				Price: 15,
			},
			{
				ID:    "item_2",
				Name:  "interview",
				Price: 10,
			},
		},
	}

	ret, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", ret)
}

func unmarshal() {
	s := `{"id":"1234","items":[{"id":"item_1","name":"learn go","price":15},{"id":"item_2","name":"interview","price":10}],"quantity":3,"total_price":20}`
	var o Order
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
}

func parseNLP() {
	res := `{
		"data":[
			{
				"synonym":"",
				"weight":"0.6",
				"word":"真絲",
				"tag":"材質"
			},
			{
				"synonym":"",
				"weight":"0.8",
				"word":"韓都衣舍",
				"tag":"品牌"
			},
			{
				"synonym":"連身裙;連衣裙",
				"weight":"1.0",
				"word":"連衣裙",
				"tag":"品類"
			}			
		]
	}`

	// m := make(map[string]interface{})
	// err := json.Unmarshal([]byte(res), &m)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", m)
	// fmt.Printf("%+v\n", m["data"].([]interface{})[2].(map[string]interface{})["synonym"])

	m := struct {
		Data []struct {
			Synonym string `json:"synonym"`
			Tag     string `json:"tag"`
		} `json:"data"`
	}{}
	err := json.Unmarshal([]byte(res), &m)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", m.Data[2].Tag)
}

func main() {
	parseNLP()
}
