package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
)

func (t *Token) IsValid() bool {
	return time.Now().Before(t.ExpiresAt)
}

func loadToken(filename string) (*Token, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var token Token
	if err := json.Unmarshal(data, &token); err != nil {
		return nil, err
	}
	return &token, nil
}
func saveToken(filename string, token *Token) error {
	data, err := json.MarshalIndent(token, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func requestNewToken() (*Token, error) {
	var resptoken token
	baseURL := cdekurl + "/v2/oauth/token"

	// Добавляем query параметры
	params := url.Values{}
	params.Add("grant_type", "client_credentials")
	params.Add("client_id", account)
	params.Add("client_secret", secure)

	fullURL := baseURL + "?" + params.Encode()

	resp, err := http.Post(fullURL, "application/json", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(respBody), &resptoken)
	if err != nil {
		panic(err)
	}

	return &Token{
		AccessToken: resptoken.Access_Token,
		ExpiresAt:   time.Now().Add(1 * time.Hour), // токен на 1 час
	}, nil
}

func tokens() *Token {
	var token *Token
	var err error

	token, err = loadToken(tokenFile)
	if err != nil || !token.IsValid() {
		token, err = requestNewToken()
		if err != nil {
			panic(err)
		}
		if err := saveToken(tokenFile, token); err != nil {
			panic(err)
		}
	}
	return token
}

func GetAllTariffs() {
	token := tokens()
	//params := url.Values{}
	//params.Add("city_code", name)
	//params.Add("country_code", "RU")
	fullURL := cdekurl + "/v2/calculator/alltariffs"
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(body))
	var tar AllTariffList
	err = json.Unmarshal([]byte(body), &tar)
	if err != nil {
		fmt.Println("ОШИБКА UNMARSHALL TARIFF")
		panic(err)
	}
	for _, x := range tar.TariffCodes {
		fmt.Println(x)
		fmt.Println("------------------------------------------")
	}
	//fmt.Println(tar.TariffCodes[0].TariffName)
}

func GetPurchaseDB(number int64, db *sql.DB) ([]Item, PackageReq) { //получаем список товаров, возвращаем общий вес, список товаров и размер коробки
	var korobas PackageReq
	var Items []Item
	var mass float64
	var size int64
	rows, err := db.Query(`SELECT p_id,name,count,options FROM rcserver_cityrton2.iY56x1_shopkeeper3_purchases WHERE order_id=?`, number)
	if err != nil {
		log.Fatal("Ошибка запроса:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var data MetriksData
		var item Item
		var Pid int32
		var Name string
		var Count float64
		var options string

		err := rows.Scan(&Pid, &Name, &Count, &options) // перечисли все нужные поля
		if err != nil {
			log.Fatal("Ошибка при чтении строки:", err)
		}
		item.Name = html.UnescapeString(Name)
		if val, ok := Art[Pid]; ok {
			item.WareKey = val
		} else {
			item.WareKey = "777"
		}
		item.Amount = int32(Count)
		err = json.Unmarshal([]byte(options), &data)
		if err != nil {
			fmt.Println("ОШИБКА UNMARSHALL OPTIONS PURCHASE")
			log.Fatal(err)
		}
		parts := strings.Split(data.Metriks[0].(string), "|")
		x, _ := strconv.Atoi(parts[0])
		height, _ := strconv.Atoi(parts[1])
		depth, _ := strconv.Atoi(parts[2])
		width, _ := strconv.Atoi(parts[3])
		size += ((int64(height)) / 100 * (int64(depth)) / 100 * (int64(width)) / 100) * int64(Count)
		//x *= int(Count)
		mass += float64(x) * float64(Count) //общий вес грузового места
		item.Weight = float64(x)            //вес 1 ед товара
		item.Cost = float64(100)
		item.Payment.Value = 0
		Items = append(Items, item)
		//fmt.Println(item)
	}
	size = int64(float64(size) * 1.3)
	for _, x := range Boxes {
		if x.Key > size {
			korobas.Height = x.Params.Height
			korobas.Length = x.Params.Length
			korobas.Width = x.Params.Width
			korobas.Weight = x.Params.Weight
			break
		}
	}
	korobas.Weight += (mass)
	return Items, korobas

}
func OnlyDigits(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r == rune('/') {
			b.WriteRune(r)
		}
		if unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return strings.TrimSpace(b.String())
}

func OnlyStreet(s string) string {
	var Prefixes = []string{
		"пр-кт", "ул", "пр-д", "пер", "шоссе", "проезд", "б-р", "бульвар", "наб", "пл", "им.", "И.С.", "им",
	}
	//fmt.Println(s)
	var inx []int
	parts := strings.Split(s, " ")
	//var result string
	for i, v := range parts {
		for _, y := range Prefixes {
			if v == y {

				inx = append(inx, i)
			}
		}
	}
	delMap := make(map[int]struct{})
	for _, i := range inx {
		delMap[i] = struct{}{}
	}
	var result []string
	for i, v := range parts {
		if _, found := delMap[i]; !found {
			result = append(result, v)
		}
	}
	return strings.Join(result, " ") + ","
}

func GetOrderDB(number int64, db *sql.DB) (Recipient, string, string, string) {
	var orderdto Recipient
	var contacts string
	var options string
	var email string
	var field []Field
	var option Options
	var phone PhoneDto
	var index string
	var deliveryddress string
	row := db.QueryRow(`SELECT contacts,options,email FROM rcserver_cityrton2.iY56x1_shopkeeper3_orders WHERE id=?`, number)
	row.Scan(&contacts, &options, &email)
	//fmt.Println(number)
	err := json.Unmarshal([]byte(contacts), &field)
	if err != nil {
		fmt.Println("Ошибка парсинга поля contacts")
		panic(err)
	}
	if options != "Array" {
		err = json.Unmarshal([]byte(options), &option)
		if err != nil {
			fmt.Println("Ошибка парсинга поля options")
			//panic(err)
		}
	}
	street := OnlyStreet(field[10].Value) //улица доставки
	dom := OnlyDigits(field[11].Value)    //дом
	cdekaddress := street + " " + dom
	orderdto.Email = email
	index = field[8].Value
	switch field[1].Value {
	case "1": //физ лицо
		deliveryddress = field[21].Value //полный адрес доставки с индексом
		orderdto.ContragentType = "INDIVIDUAL"
		phone.Number = field[16].Value
		orderdto.Phones = append(orderdto.Phones, phone)
		orderdto.Name = field[19].Value
	case "2": //юр лицо
		deliveryddress = field[30].Value //полный адрес доставки с индексом
		orderdto.ContragentType = "LEGAL_ENTITY"
		orderdto.Company = html.UnescapeString(field[16].Value)
		orderdto.Tin = field[17].Value
		phone.Number = field[26].Value
		orderdto.Phones = append(orderdto.Phones, phone)
		orderdto.Name = field[28].Value
	case "3": //ИП
		deliveryddress = field[28].Value //полный адрес доставки с индексом
		orderdto.ContragentType = "LEGAL_ENTITY"
		orderdto.Company = "ИП " + html.UnescapeString(field[16].Value)
		orderdto.Tin = field[17].Value
		phone.Number = field[19].Value
		orderdto.Phones = append(orderdto.Phones, phone)
		orderdto.Name = field[26].Value
	}
	return orderdto, deliveryddress, cdekaddress, index
}

func RequestTariff() {
	var d RequestTariffSchema
	d.TariffCode = 483
	d.Lang = "rus"
	d.Type = 1
	d.FromLocation.Code = 137
	d.FromLocation.ContragentType = "LEGAL_ENTITY"
	d.ToLocation.Code = 438
	d.ToLocation.ContragentType = "INDIVIDUAL"
	d.Packages = append(d.Packages, PackageReq{Weight: 2000.0, Length: 30, Width: 30, Height: 30})
	token := tokens()
	//params := url.Values{}
	//params.Add("city_code", name)
	//params.Add("country_code", "RU")
	fullURL := cdekurl + "/v2/calculator/tariff"
	jsonData, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	/*
		var tar TariffList
		err = json.Unmarshal([]byte(body), &tar)
		if err != nil {
			panic(err)
		}

			for _, x := range tar.TariffCodes {
				fmt.Println(x.TariffName, x.DeliverySum, x.TariffCode)
				fmt.Println("------------------------------------------")
			}
	*/
}

func RequestTariffList(codetolocation int32, contragenttype string, ves PackageReq) TariffList {
	fmt.Println("ВЫЗОВ ИЗ RequestTariffList", contragenttype)
	var d RequestTariffSchema
	d.Lang = "rus"
	d.Type = 1
	d.FromLocation.Code = 137
	d.FromLocation.ContragentType = "LEGAL_ENTITY"
	d.ToLocation.Code = codetolocation
	d.ToLocation.ContragentType = contragenttype
	d.Packages = append(d.Packages, ves)
	token := tokens()
	//params := url.Values{}
	//params.Add("city_code", name)
	//params.Add("country_code", "RU")
	fullURL := cdekurl + "/v2/calculator/tarifflist"
	jsonData, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(body))
	var tar TariffList
	err = json.Unmarshal([]byte(body), &tar)
	if err != nil {
		fmt.Println("ОШИБКА UNMARSHALL REQUESTTARIFFLIST")
		panic(err)
	}

	//	for _, x := range tar.TariffCodes {
	//		fmt.Printf("Название тарифа:%s, Режим тарифа:%d", x.TariffName, x.DeliveryMode)
	//		fmt.Println("------------------------------------------")
	//	}
	return tar
}

func GetCityIndex(name string) City { //получение кода города по индексу
	//fmt.Println("ВЫЗОВ ИЗ GO", name)
	token := tokens()
	params := url.Values{}
	params.Add("postal_code", name)
	//params.Add("country_code", "RU")
	fullURL := cdekurl + "/v2/location/cities" + "?" + params.Encode()
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var city []City
	//fmt.Println(string(body))
	err = json.Unmarshal([]byte(body), &city)
	if err != nil {
		fmt.Println("ОШИБКА UNMARSHALL GETCITYINDEX")
		panic(err)
	}
	if len(city) != 0 {
		return city[0]
	} else {
		fmt.Println("Город не найден")
		return City{}
	}
	//fmt.Println(city[0].Code)
	//return []string{"Москва", "Санкт-Петербург"}
}

func OfficesList(citycode string, index string, code string) []Deliverypoints {
	//fmt.Println("ВЫЗОВ ИЗ GO", name)
	token := tokens()
	params := url.Values{}
	params.Add("city_code", citycode)
	params.Add("postal_code", index)
	params.Add("code", code) //код ПВЗ
	//params.Add("country_code", "RU")
	fullURL := cdekurl + "/v2/deliverypoints" + "?" + params.Encode()
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(body))
	var offices []Deliverypoints
	err = json.Unmarshal([]byte(body), &offices)
	if err != nil {
		fmt.Println("ОШИБКА UNMARSHALL OFFICELIST")
		panic(err)
	}
	return offices
	//return []string{"Москва", "Санкт-Петербург"}
}

func MinTariffSkladDver(a TariffList) TariffCode {
	var d []TariffCode
	var tar TariffCode
	for x, _ := range a.TariffCodes {
		if a.TariffCodes[x].DeliveryMode == 3 {
			d = append(d, a.TariffCodes[x])
		}
	}
	tar = d[0]
	for _, t := range d {
		if t.DeliverySum < tar.DeliverySum {
			tar = t
		}
	} //sort.Slice()
	return tar

}

func MinTariffSkladSklad(a TariffList) TariffCode {
	var d []TariffCode
	var tar TariffCode
	for x, _ := range a.TariffCodes {
		if a.TariffCodes[x].DeliveryMode == 4 {
			d = append(d, a.TariffCodes[x])
		}
	}
	tar = d[0]
	for _, t := range d {
		if t.DeliverySum < tar.DeliverySum {
			tar = t
		}
	} //sort.Slice()
	return tar
}

func Zakaz(number int64, db *sql.DB) ZakazDTO {
	var zakaz ZakazDTO
	var mesto Packages
	recip, fullad, cdekad, index := GetOrderDB(number, db)
	item, korob := GetPurchaseDB(number, db)
	city := GetCityIndex(index)
	officelist := OfficesList(strconv.Itoa(int(city.Code)), "", "")
	xx := RequestTariffList(city.Code, recip.ContragentType, korob)
	zakaz.Recipient = recip
	zakaz.Type = 1
	mesto.Weight = korob.Weight
	mesto.Height = korob.Height
	mesto.Length = korob.Length
	mesto.Width = korob.Width
	mesto.Number = "1"
	mesto.Items = append(mesto.Items, item...)
	zakaz.Packages = append(zakaz.Packages, mesto)
	zakaz.Print = "BARCODE"
	zakaz.City = city
	zakaz.OfficeList = officelist
	zakaz.TariffList = xx
	//fmt.Println(xx)
	if len(officelist) != 0 {
		for x, _ := range officelist {
			if strings.Contains(officelist[x].Location.Address, cdekad) {
				zakaz.Number = strconv.Itoa(int(number))
				//zakaz.ShipmentPoint = "SPB55"
				zakaz.ShipmentPoint = "SPB261" //Энергетиков 8, к1, только в боевом листе
				zakaz.DeliveryPoint = officelist[x].Code
				zakaz.TariffCode = int32(MinTariffSkladSklad(xx).TariffCode)
				//fmt.Println(zakaz.OfficeList)
				return zakaz
			}
		}
	}

	zakaz.Number = strconv.Itoa(int(number))
	//zakaz.ShipmentPoint = "SPB55"
	zakaz.ShipmentPoint = "SPB261"
	zakaz.ToLocation.Address = fullad
	zakaz.ToLocation.Code = city.Code
	zakaz.TariffCode = int32(MinTariffSkladDver(xx).TariffCode)
	return zakaz
}

func SendZakaz(zakaz ZakazDTO) RootEntityDto {
	token := tokens()
	fullURL := cdekurl + "/v2/orders"
	jsonData, err := json.Marshal(zakaz)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var x RootEntityDto
	err = json.Unmarshal([]byte(body), &x)
	if err != nil {
		fmt.Println("ОШИБКА UNMARSHALL SENDZAKAZ")
		panic(err)
	}
	return x
}

func RequestInfo(id string) {
	token := tokens()
	//params.Add("country_code", "RU")
	fullURL := cdekurl + "/v2/orders/" + id
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func RequestBarCode(id string) {
	var d PrintOrdersRequest
	d.Orders = append(d.Orders, PrintOrderDto{OrderUUID: id})
	d.Format = "A6"
	d.CopyCount = 1
	d.Lang = "RUS"
	token := tokens()
	fullURL := cdekurl + "/v2/print/barcodes"
	jsonData, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) []CityAutocomplete {
	//fmt.Println("ВЫЗОВ ИЗ GO", name)
	token := tokens()
	params := url.Values{}
	params.Add("name", name)
	params.Add("country_code", "RU")
	fullURL := cdekurl + "/v2/location/suggest/cities" + "?" + params.Encode()
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(body))
	var cit []CityAutocomplete
	err = json.Unmarshal([]byte(body), &cit)
	if err != nil {
		panic(err)
	}
	//return []string{"Москва", "Санкт-Петербург"}
	return cit
}

func (a *App) OfficesList(name string) []Deliverypoints {
	//fmt.Println("ВЫЗОВ ИЗ GO", name)
	token := tokens()
	params := url.Values{}
	params.Add("city_code", name)
	params.Add("country_code", "RU")
	fullURL := cdekurl + "/v2/deliverypoints" + "?" + params.Encode()
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var offices []Deliverypoints
	err = json.Unmarshal([]byte(body), &offices)
	if err != nil {
		panic(err)
	}
	//return []string{"Москва", "Санкт-Петербург"}
	return offices
}

func (a *App) LoadZakaz(x int64) interface{} {
	//var db *sql.DB
	//var err error
	//dsn := "root:secret123@tcp(127.0.0.1:3306)/cdekdb?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err = sql.Open("mysql", dsn)
	//if err != nil {
	username := "rcserver_cityrton2"
	password := "5$QSkS%hntL33o3Ey9aaoqVvOcP4xF"
	host := "cityron.ru"
	port := "3306"
	dbname := "rcserver_cityrton2"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	// открываем соединение
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}
	defer db.Close()

	//}
	xx := Zakaz(x, db)
	fmt.Println(xx)
	return xx
}

func (a *App) Printm(x Deliverypoints) {
	fmt.Println(x)
}
