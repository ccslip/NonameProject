package main

import (
	"embed"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

type CityAutocomplete struct {
	CityUUID    uuid.UUID `json:"city_uuid"`
	Code        int32     `json:"code"`
	FullName    string    `json:"full_name"`
	CountryCode string    `json:"country_code"`
}

type Phone struct {
	Number     string `json:"number"`
	Additional string `json:"additional"`
}
type Officeimage struct {
	Number int64  `json:"number"`
	URL    string `json:"url"`
}

type Worktime struct {
	Day  int64  `json:"day"`
	Time string `json:"time"`
}

type Worktimeexception struct {
	Datestart string    `json:"date_start"`
	Dateend   string    `json:"date_end"`
	Timestart Timestart `json:"time_start"`
	Timeend   Timeend   `json:"time_end"`
	Isworking bool      `json:"is_working"`
}

type Timestart struct {
	Hour   int32 `json:"hour"`
	Minute int32 `json:"minute"`
	Second int32 `json:"second"`
	Nano   int32 `json:"nano"`
}

type Timeend struct {
	Hour   int32 `json:"hour"`
	Minute int32 `json:"minute"`
	Second int32 `json:"second"`
	Nano   int32 `json:"nano"`
}

type Dimension struct {
	Width  int32 `json:"width"`
	Height int32 `json:"height"`
	Depth  int32 `json:"depth"`
}

type Error struct {
	Code           string `json:"code"`
	Additionalcode string `json:"additional_code"`
	Message        string `json:"message"`
}

type Warning struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type DeliveryDateRange struct {
	Min time.Time `json:"min"`
	Max time.Time `json:"max"`
}

type TariffCode struct {
	TariffCode        int               `json:"tariff_code"`
	TariffName        string            `json:"tariff_name"`
	TariffDescription string            `json:"tariff_description"`
	DeliveryMode      int               `json:"delivery_mode"`
	DeliverySum       float64           `json:"delivery_sum"`
	PeriodMin         int               `json:"period_min"`
	PeriodMax         int               `json:"period_max"`
	CalendarMin       int               `json:"calendar_min"`
	CalendarMax       int               `json:"calendar_max"`
	DeliveryDateRange DeliveryDateRange `json:"delivery_date_range"`
}

type TariffList struct {
	TariffCodes []TariffCode `json:"tariff_codes"`
	Errors      []Error      `json:"errors"`
	Warnings    []Warning    `json:"warnings"`
}

type Location struct {
	Countrycode string    `json:"country_code"`
	Regioncode  int32     `json:"region_code"`
	Region      string    `json:"region"`
	Citycode    int32     `json:"city_code"`
	City        string    `json:"city"`
	Fiasguid    uuid.UUID `json:"fias_guid"`
	Postalcode  string    `json:"postal_code"`
	Longitude   float64   `json:"longitude"`
	Latitude    float64   `json:"latitude"`
	Address     string    `json:"address"`
	Addressfull string    `json:"address_full"`
	Cityuuid    uuid.UUID `json:"city_uuid"`
}
type City struct {
	Code           int        `json:"code"`                       // Код населенного пункта СДЭК
	CityUUID       uuid.UUID  `json:"city_uuid"`                  // Идентификатор населенного пункта в ИС СДЭК
	City           string     `json:"city"`                       // Название населенного пункта
	FiasGuid       *uuid.UUID `json:"fias_guid,omitempty"`        // Уникальный идентификатор ФИАС (может отсутствовать)
	CountryCode    string     `json:"country_code"`               // Код страны (ISO_3166-1_alpha-2)
	Country        string     `json:"country"`                    // Название страны
	Region         string     `json:"region"`                     // Название региона
	RegionCode     *int       `json:"region_code,omitempty"`      // Код региона СДЭК (может отсутствовать)
	FiasRegionGuid *uuid.UUID `json:"fias_region_guid,omitempty"` // Устаревшее поле
	SubRegion      *string    `json:"sub_region,omitempty"`       // Название района региона
	Longitude      *float64   `json:"longitude,omitempty"`        // Долгота
	Latitude       *float64   `json:"latitude,omitempty"`         // Широта
	TimeZone       *string    `json:"time_zone,omitempty"`        // Часовой пояс
	PaymentLimit   float64    `json:"payment_limit"`              // Ограничение наложенного платежа
}
type Deliverypoints struct {
	Code                  string              `json:"code"`
	Address               string              `json:"address"`
	Name                  string              `json:"name"`
	UUID                  string              `json:"uuid"`
	Addresscomment        string              `json:"address_comment"`
	Neareststation        string              `json:"nearest_station"`
	Nearestmetrostation   string              `json:"nearest_metro_station"`
	Worktime              string              `json:"work_time"`
	Phones                []Phone             `json:"phones"`
	Email                 string              `json:"email"`
	Note                  string              `json:"note"`
	Type                  string              `json:"type"`
	Ownercode             string              `json:"owner_code"`
	Takeonly              bool                `json:"take_only"`
	Ishandout             bool                `json:"is_handout"`
	Isreception           bool                `json:"is_reception"`
	Isdressingroom        bool                `json:"is_dressing_room"`
	Ismarketplace         bool                `json:"is_marketplace"`
	Isltl                 bool                `json:"is_ltl"`
	Havecashless          bool                `json:"have_cashless"`
	Havecash              bool                `json:"have_cash"`
	Havefastpaymentsystem bool                `json:"have_fast_payment_system"`
	Allowedcod            bool                `json:"allowed_cod"`
	Site                  string              `json:"site"`
	Officeimagelist       []Officeimage       `json:"office_image_list"`
	Worktimelist          []Worktime          `json:"work_time_list"`
	Worktimeexceptionlist []Worktimeexception `json:"work_time_exception_list"`
	Weightmin             float64             `json:"weight_min"`
	Weightmax             float64             `json:"weight_max"`
	Dimensions            []Dimension         `json:"dimensions"`
	Errors                []Error             `json:"errors"`
	Warnings              []Warning           `json:"warnings"`
	Location              Location            `json:"location"`
	Distance              int64               `json:"distance"`
	Fulfillment           bool                `json:"fulfillment"`
}

type Token struct {
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

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
	baseURL := "http://api.edu.cdek.ru/v2/oauth/token"

	// Добавляем query параметры
	params := url.Values{}
	params.Add("grant_type", "client_credentials")
	params.Add("client_id", "wqGwiQx0gg8mLtiEKsUinjVSICCjtTEP")
	params.Add("client_secret", "RmAmgvSgSl1yirlz9QupbzOJVqhCxcP5")

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

type token struct {
	Access_Token string `json:"access_token"`
	Token_Type   string `json:"token_type"`
	Expires_In   int64  `json:"expires_in"`
	Scope        string `json:"scope"`
	Jti          string `json:"jti"`
}

const tokenFile = "token.json"

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

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "desktop",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
