package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"strconv"
	"log"
	"encoding/json"
	"sync"
)

type Purchase struct {
	CustomerName string
	Price float64
}

type Data struct {
	List []Purchase
}

type Response struct {
	Data Data
}

func main(){
	var wg sync.WaitGroup

	for _, region := range os.Args[1:] {
		wg.Add(1)

		region_id, _ := strconv.Atoi(region)
		go print_purchases(region_id, &wg)
	}

	wg.Wait()
}

func print_purchases(region_id int, wg *sync.WaitGroup) {
	defer wg.Done()

	file, _ := os.Create(fmt.Sprintf("%d.txt", region_id))
	defer file.Close()

	for page_num := 1; page_num <= 100; page_num++ {
		fmt.Printf("Region %d, page %d\n", region_id, page_num)

		url := fmt.Sprintf("https://zakupki.gov.ru/api/mobile/proxy/917/epz/order/extendedsearch/results.html?morphology=on&pageNumber=%d&sortDirection=false&fz44=on&pc=on&currencyId=-1&sortBy=UPDATE_DATE&regions=%d", page_num, region_id)
		resp, err := http.Get(url)

		if err != nil{
			log.Fatal(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil{
			log.Fatal(err)
		}

		var r Response
		ser_err := json.Unmarshal(body, &r)
		if ser_err != nil{
			log.Fatal(ser_err)
		}
		
		for _, purchase := range r.Data.List{
			text := fmt.Sprintf("Region %d; customer %v; price %f\n", region_id, purchase.CustomerName, purchase.Price)
			file.WriteString(text)
		}
	}
}
