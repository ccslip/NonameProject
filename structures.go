package main

import (
	"context"
	"time"

	"github.com/google/uuid"
)

//const cdekurl = "http://api.cdek.ru"
//const account = "VuTPedBjnRFFODZw2SE8ZuONsQYNMZwd"
//const secure = "cpHyV2qCgHuZhyycpOCZ4LFFSs1jZ7PO"

const tokenFile = "token.json"

const cdekurl = "http://api.edu.cdek.ru"
const account = "wqGwiQx0gg8mLtiEKsUinjVSICCjtTEP"
const secure = "RmAmgvSgSl1yirlz9QupbzOJVqhCxcP5"

var Art map[int32]string

type Box struct {
	Key    int64
	Params PackageReq
}

var Boxes = []Box{
	{1176, PackageReq{Length: 12, Width: 14, Height: 7, Weight: 50}},
	{1862, PackageReq{Length: 19, Width: 14, Height: 7, Weight: 60}},
	{6000, PackageReq{Length: 20, Width: 20, Height: 15, Weight: 100}},
	{6864, PackageReq{Length: 26, Width: 22, Height: 12, Weight: 130}},
	{8000, PackageReq{Length: 20, Width: 20, Height: 20, Weight: 130}},
	{27000, PackageReq{Length: 30, Width: 30, Height: 30, Weight: 270}},
}

func init() {
	Art = map[int32]string{
		5291: "2013",
		8055: "2017",
		5439: "2001",
		7668: "2022",
		5274: "2021",
		7827: "2023",
		8036: "2400",
		8067: "2401",
		7788: "1008",
		7842: "3004",
		7985: "3005",
		5164: "HDR-15-24",
		5211: "HDR-30-24",
		5214: "HDR-60-24",
		4566: "3012",
		4998: "3003",
		5563: "3006",
		7675: "1DA25A",
		7676: "1DA40A",
		7921: "РТР-1140",
		7834: "3DA60A",
		7772: "R0903N",
		7773: "R0904N",
		7678: "3DA40A",
		7677: "3DA25A",
		7681: "РТР-325",
		7835: "РТР-360",
		3419: "HDR-100-24",
	}
}

type CityAutocomplete struct {
	CityUUID    uuid.UUID `json:"city_uuid"`
	Code        int32     `json:"code"`
	FullName    string    `json:"full_name"`
	CountryCode string    `json:"country_code"`
}

// App struct
type App struct {
	ctx context.Context
}

type Phone struct {
	Number     string `json:"number"`
	Additional string `json:"additional"`
}
type Officeimage struct {
	Number int32  `json:"number"`
	URL    string `json:"url"`
}

type Worktime struct {
	Day  int64  `json:"day"`
	Time string `json:"time"`
}

type Worktimeexception struct {
	Datestart string `json:"date_start"`
	Dateend   string `json:"date_end"`
	Timestart string `json:"time_start"`
	Timeend   string `json:"time_end"`
	Isworking bool   `json:"is_working"`
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
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Depth  float64 `json:"depth"`
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
	Min string `json:"min"`
	Max string `json:"max"`
}

type TariffCode struct {
	TariffCode        int               `json:"tariff_code"`
	TariffName        string            `json:"tariff_name"`
	TariffDescription string            `json:"tariff_description"`
	DeliveryMode      int32             `json:"delivery_mode"`
	DeliverySum       float64           `json:"delivery_sum"`
	PeriodMin         int32             `json:"period_min"`
	PeriodMax         int32             `json:"period_max"`
	CalendarMin       int32             `json:"calendar_min"`
	CalendarMax       int32             `json:"calendar_max"`
	DeliveryDateRange DeliveryDateRange `json:"delivery_date_range"`
}

type TariffList struct {
	TariffCodes []TariffCode `json:"tariff_codes"`
	Errors      []Error      `json:"errors"`
	Warnings    []Warning    `json:"warnings"`
}

type Tariff struct {
	TariffName                string                    `json:"tariff_name"`
	WeightMin                 float64                   `json:"weight_min"`
	WeightMax                 float64                   `json:"weight_max"`
	WeightCalcMax             float64                   `json:"weight_calc_max"`
	LengthMin                 int32                     `json:"length_min"`
	LengthMax                 int32                     `json:"length_max"`
	WidthMin                  int32                     `json:"width_min"`
	WidthMax                  int32                     `json:"width_max"`
	HeightMin                 int32                     `json:"height_min"`
	HeightMax                 int32                     `json:"height_max"`
	OrderTypes                []string                  `json:"order_types"`
	PayerContragentType       []string                  `json:"payer_contragent_type"`
	SenderContragentType      []string                  `json:"sender_contragent_type"`
	RecipientContragentType   []string                  `json:"recipient_contragent_type"`
	DeliveryModes             []DeliveryMode            `json:"delivery_modes"`
	AdditionalOrderTypesParam AdditionalOrderTypesParam `json:"additional_order_types_param"`
}
type DeliveryMode struct {
	DeliveryMode     string `json:"delivery_mode"`
	DeliveryModeName string `json:"delivery_mode_name"`
	TariffCode       int32  `json:"tariff_code"`
}

type AdditionalOrderTypesParam struct {
	WithoutAdditionalOrderType bool     `json:"without_additional_order_type"`
	AdditionalOrderTypes       []string `json:"additional_order_types"`
}
type AllTariffList struct {
	TariffCodes []Tariff `json:"tariff_codes"`
}

type Location struct {
	Countrycode string  `json:"country_code"`
	Regioncode  int32   `json:"region_code"`
	Region      string  `json:"region"`
	Citycode    int32   `json:"city_code"`
	City        string  `json:"city"`
	Fiasguid    string  `json:"fias_guid"`
	Postalcode  string  `json:"postal_code"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	Address     string  `json:"address"`
	Addressfull string  `json:"address_full"`
	Cityuuid    string  `json:"city_uuid"`
}
type City struct {
	Code           int32      `json:"code"`             // Код населенного пункта СДЭК
	CityUUID       uuid.UUID  `json:"city_uuid"`        // Идентификатор населенного пункта в ИС СДЭК
	City           string     `json:"city"`             // Название населенного пункта
	FiasGuid       *uuid.UUID `json:"fias_guid"`        // Уникальный идентификатор ФИАС (может отсутствовать)
	CountryCode    string     `json:"country_code"`     // Код страны (ISO_3166-1_alpha-2)
	Country        string     `json:"country"`          // Название страны
	Region         string     `json:"region"`           // Название региона
	RegionCode     *int       `json:"region_code"`      // Код региона СДЭК (может отсутствовать)
	FiasRegionGuid *uuid.UUID `json:"fias_region_guid"` // Устаревшее поле
	SubRegion      *string    `json:"sub_region"`       // Название района региона
	Longitude      *float64   `json:"longitude"`        // Долгота
	Latitude       *float64   `json:"latitude"`         // Широта
	TimeZone       *string    `json:"time_zone"`        // Часовой пояс
	PaymentLimit   float64    `json:"payment_limit"`    // Ограничение наложенного платежа
}
type Deliverypoints struct {
	Code string `json:"code"`
	//Name                  string              `json:"name"`
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

type Country struct {
	Code int64     `json:"code"` // внутренний код страны
	ISO  string    `json:"iso"`  // ISO-код (например "RU")
	Name string    `json:"name"` // название страны
	UUID uuid.UUID `json:"uuid"` // UUID страны
}

type Region struct {
	UUID uuid.UUID `json:"uuid"` // UUID региона
	Code int64     `json:"code"` // внутренний код региона
	Name string    `json:"name"` // название региона
}

type Autocomplete struct {
	UUID                        uuid.UUID `json:"uuid"`                        // UUID населённого пункта
	Code                        int64     `json:"code"`                        // внутренний код
	Name                        string    `json:"name"`                        // короткое имя (город)
	FullName                    string    `json:"full_name"`                   // полное название (например "г. Москва, Россия")
	Country                     Country   `json:"country"`                     // вложенный объект Country
	Region                      Region    `json:"region"`                      // вложенный объект Region
	AvailableAdaptiveSuggestion bool      `json:"availableAdaptiveSuggestion"` // сделал с большой буквы, чтобы поле сериализовалось
}
type Token struct {
	AccessToken string    `json:"access_token"`
	ExpiresAt   time.Time `json:"expires_at"`
}

type token struct {
	Access_Token string `json:"access_token"`
	Token_Type   string `json:"token_type"`
	Expires_In   int64  `json:"expires_in"`
	Scope        string `json:"scope"`
	Jti          string `json:"jti"`
}
type RequestTariffSchema struct {
	Date                 string       `json:"date"`
	Type                 int32        `json:"type"`
	AdditionalOrderTypes []int32      `json:"additional_order_types"`
	Currency             int32        `json:"currency"`
	Lang                 string       `json:"lang"`
	TariffCode           int32        `json:"tariff_code"`
	FromLocation         LocationReq  `json:"from_location"`
	ToLocation           LocationReq  `json:"to_location"`
	Packages             []PackageReq `json:"packages"`
}

type LocationReq struct {
	Code           int32   `json:"code"`
	PostalCode     string  `json:"postal_code"`
	CountryCode    string  `json:"country_code"`
	City           string  `json:"city"`
	Address        string  `json:"address"`
	ContragentType string  `json:"contragent_type"`
	Longitude      float64 `json:"longitude"`
	Latitude       float64 `json:"latitude"`
}

type PackageReq struct {
	Weight float64 `json:"weight"`
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}
type Orders struct {
	Id             int64   `json:"id"`
	Contacts       []Field `json:"contacts"`
	Options        string  `json:"options"`
	Price          float64 `json:"price"`
	Currency       string  `json:"currency"`
	Date           string  `json:"date"`
	Sentdate       string  `json:"sentdate"`
	Statusdate     string  `json:"statusdate"`
	Note           string  `json:"note"`
	Email          string  `json:"email"`
	Delivery       string  `json:"delivery"`
	Delivery_Price float64 `json:"delivery_price"`
	Payment        string  `json:"payment"`
	Tracking_Num   string  `json:"tracking_num"`
	Status         string  `json:"status"`
	Userid         int64   `json:"userid"`
}

type Field struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Label string `json:"label"`
}

type Options struct {
	DeliverPrice string `json:"delivery_price"`
	Payed        string `json:"payed"`
	PayDate      int32  `json:"paydate"`
	PaymentURL   string `json:"paymentURL"`
	PaymentHash  string `json:"paymentHash"`
	//PaymentID    bool   `json:"paymentID"`
}

type SkladDver struct {
	Type          int32            `json:"type"`                  //Тип заказа.
	Number        string           `json:"number"`                //Номера заказа в ИС Клиента. Только для заказов "интернет-магазин".
	TariffCode    int32            `json:"tariff_code"`           //Код тарифа
	ShipmentPoint string           `json:"shipment_point"`        //Код ПВЗ СДЭК, на который будет производиться самостоятельный привоз клиентом. Обязательное поле, если заказ с тарифом "от склада". Не может использоваться одновременно с from_location
	Recipient     Recipient        `json:"recipient"`             //Получатель
	ToLocation    ToLocation       `json:"to_location,omitempty"` //Адрес получения. Заполняется, если тариф "до двери".
	Packages      []Packages       `json:"packages"`              //Список информации по местам (упаковкам). Количество мест в заказе может быть от 1 до 255
	Print         string           `json:"print"`                 //Enum: "WAYBILL" "BARCODE" Тип печатной формы, которую необходимо сформировать по заказу
	OfficeList    []Deliverypoints `json:"office_list"`
	TariffList    TariffList       `json:"tariff_list"`
}

type SkladSklad struct {
	Type          int32            `json:"type"`           //Тип заказа.
	Number        string           `json:"number"`         //Номера заказа в ИС Клиента. Только для заказов "интернет-магазин".
	TariffCode    int32            `json:"tariff_code"`    //Код тарифа
	ShipmentPoint string           `json:"shipment_point"` //Код ПВЗ СДЭК, на который будет производиться самостоятельный привоз клиентом. Обязательное поле, если заказ с тарифом "от склада". Не может использоваться одновременно с from_location
	DeliveryPoint string           `json:"delivery_point"`
	Recipient     Recipient        `json:"recipient"` //Получатель
	Packages      []Packages       `json:"packages"`  //Список информации по местам (упаковкам). Количество мест в заказе может быть от 1 до 255
	Print         string           `json:"print"`     //Enum: "WAYBILL" "BARCODE" Тип печатной формы, которую необходимо сформировать по заказу
	OfficeList    []Deliverypoints `json:"office_list"`
	TariffList    TariffList       `json:"tariff_list"`
}

type ZakazDTO struct { //ФОРМА ДЛЯ РАЗМЕЩЕНИЯ ЗАКАЗА
	Type                 int32   `json:"type"` //Тип заказа.
	AdditionalOrderTypes []int32 `json:"additional_order_types"`
	Number               string  `json:"number"` //Номера заказа в ИС Клиента. Только для заказов "интернет-магазин".
	AccompanyingNumber   string  `json:"accompanying_number"`
	TariffCode           int32   `json:"tariff_code"`    //Код тарифа
	Comment              string  `json:"comment"`        //Комментарий к заказу
	ShipmentPoint        string  `json:"shipment_point"` //Код ПВЗ СДЭК, на который будет производиться самостоятельный привоз клиентом. Обязательное поле, если заказ с тарифом "от склада". Не может использоваться одновременно с from_location
	//DeliveryPoint            string    `json:"delivery_point,"`
	DateInvoice              string    `json:"date_invoice"`
	ShipperName              string    `json:"shipper_name"`
	ShipperAddress           string    `json:"shipper_address"`
	DeliveryRecipientCost    Money     `json:"delivery_recipient_cost"`
	DeliveryRecipientCostAdv []Money1  `json:"delivery_recipient_cost_adv"`
	Sender                   Recipient `json:"sender"`
	Seller                   Seller    `json:"seller"`    //Реквизиты истинного продавца
	Recipient                Recipient `json:"recipient"` //Получатель
	//FromLocation             FromLocation `json:"from_location"` //Адрес отправления. Заполняется, если тариф "от двери".
	ToLocation      ToLocation `json:"to_location,omitempty"` //Адрес получения. Заполняется, если тариф "до двери".
	Packages        []Packages `json:"packages"`              //Список информации по местам (упаковкам). Количество мест в заказе может быть от 1 до 255
	IsClientReturn  bool       `json:"is_client_return"`      //Признак клиентского возврата
	HasReverseOrder bool       `json:"has_reverse_order"`     //Признак необходимости создания реверсного заказа
	DeveloperKey    string     `json:"developer_key"`         //Ключ разработчика
	Print           string     `json:"print"`                 //Enum: "WAYBILL" "BARCODE" Тип печатной формы, которую необходимо сформировать по заказу
	WidgetToken     string     `json:"widget_token"`          //Токен CMS, содержащий дополнительные данные
}

type Packages struct {
	Number    string  `json:"number"`     // обязательное поле
	Weight    float64 `json:"weight"`     // обязательное поле (в граммах)
	Length    float64 `json:"length"`     // длина (см), может отсутствовать
	Width     float64 `json:"width"`      // ширина (см), может отсутствовать
	Height    float64 `json:"height"`     // высота (см), может отсутствовать
	Comment   string  `json:"comment"`    // комментарий к упаковке
	Items     []Item  `json:"items"`      // список товаров
	PackageID string  `json:"package_id"` // уникальный id упаковки в СДЭК
}

type Item struct {
	Name        string  `json:"name"`         // обязательное поле
	WareKey     string  `json:"ware_key"`     // обязательное поле
	Marking     string  `json:"marking"`      // маркировка (честный знак)
	Payment     Money   `json:"payment"`      // обязательное поле
	Weight      float64 `json:"weight"`       // обязательное поле
	WeightGross int32   `json:"weight_gross"` // вес брутто
	Amount      int32   `json:"amount"`       // обязательное поле
	NameI18n    string  `json:"name_i18n"`    // для международных заказов
	Brand       string  `json:"brand"`
	CountryCode string  `json:"country_code"`
	Material    string  `json:"material"`
	WifiGsm     bool    `json:"wifi_gsm"`
	URL         string  `json:"url"`
	Cost        float64 `json:"cost"`       // обязательное поле
	FeacnCode   string  `json:"feacn_code"` // код ТН ВЭД
	JewelUIN    string  `json:"jewel_uin"`  // УИН ювелирного изделия
	Used        bool    `json:"used"`       // признак б/у
}

type Money struct {
	Value   float64 `json:"value"`    // обязательное поле, сумма платежа включая НДС
	VatSum  float64 `json:"vat_sum"`  // сумма НДС (обязательна, если задан vat_rate)
	VatRate int32   `json:"vat_rate"` // ставка НДС (0,5,7,10,12,20 или null)
}

type Money1 struct {
	Threshold int32   `json:"threshold"`
	Sum       float64 `json:"sum"`      // обязательное поле, сумма платежа включая НДС
	VatSum    float64 `json:"vat_sum"`  // сумма НДС (обязательна, если задан vat_rate)
	VatRate   int32   `json:"vat_rate"` // ставка НДС (0,5,7,10,12,20 или null)
}

type FromLocation struct {
	Code           int32   `json:"code"`             //Код населенного пункта СДЭК (метод "Список населенных пунктов")
	CityUUID       string  `json:"city_uuid"`        //Идентификатор города в ИС СДЭК
	City           string  `json:"city"`             //Название населенного пункта
	FiasGuid       string  `json:"fias_guid"`        //Идентификатор ФИАС населенного пункта
	CountryCode    string  `json:"country_code"`     //Код страны в формате ISO_3166-1_alpha-2
	Country        string  `json:"country"`          //Название страны населенного пункта
	Region         string  `json:"region"`           //Название региона
	RegionCode     int32   `json:"region_code"`      //Код региона СДЭК (см. метод "Список регионов")
	FiasRegionGuid string  `json:"fias_region_guid"` //Уникальный идентификатор ФИАС региона населенного пункта
	SubRegion      string  `json:"sub_region"`       //Название района региона
	Longitude      float64 `json:"longitude"`        //Долгота
	Latitude       float64 `json:"latitude"`         //Широта
	TimeZone       string  `json:"time_zone"`        //Часовой пояс населенного пункта
	PaymentLimit   float64 `json:"payment_limit"`    //Ограничение на сумму наложенного платежа в населенном пункте
	Address        string  `json:"address"`          // required Строка адреса
	PostalCode     string  `json:"postal_code"`      //Почтовый индекс
}

type ToLocation struct {
	Code           int32   `json:"code,omitempty"`
	CityUUID       string  `json:"city_uuid,omitempty"` // string вместо uuid.UUID
	City           string  `json:"city,omitempty"`
	FiasGUID       string  `json:"fias_guid,omitempty"` // string вместо uuid.UUID
	CountryCode    string  `json:"country_code,omitempty"`
	Country        string  `json:"country,omitempty"`
	Region         string  `json:"region,omitempty"`
	RegionCode     int32   `json:"region_code,omitempty"`
	FiasRegionGUID string  `json:"fias_region_guid,omitempty"` // string вместо uuid.UUID
	SubRegion      string  `json:"sub_region,omitempty"`
	Longitude      float64 `json:"longitude,omitempty"`
	Latitude       float64 `json:"latitude,omitempty"`
	TimeZone       string  `json:"time_zone,omitempty"`
	PaymentLimit   float64 `json:"payment_limit,omitempty"`
	Address        string  `json:"address,omitempty"`
	PostalCode     string  `json:"postal_code,omitempty"`
}

type Recipient struct {
	Company              string     `json:"company"`
	Name                 string     `json:"name"` // required
	ContragentType       string     `json:"contragent_type"`
	PassportSeries       string     `json:"passport_series"`
	PassportNumber       string     `json:"passport_number"`
	PassportDateOfIssue  string     `json:"passport_date_of_issue"` // string в формате date
	PassportOrganization string     `json:"passport_organization"`
	Tin                  string     `json:"tin"`
	PassportDateOfBirth  string     `json:"passport_date_of_birth"` // string в формате date
	Email                string     `json:"email"`
	Phones               []PhoneDto `json:"phones"` // Массив телефонов
}

type PhoneDto struct {
	Number     string `json:"number"`
	Additional string `json:"additional"`
}

type Seller struct {
	Name          string `json:"name"`           //Наименование истинного продавца
	INN           string `json:"inn"`            //ИНН истинного продавца
	Phone         string `json:"phone"`          //Телефон истинного продавца
	OwnershipForm string `json:"ownership_form"` //Код формы собственности
	Address       string `json:"address"`        //Адрес истинного продавца
}

type OrderDTO struct {
	ContragentType  string `json:"contragent_type"`  //тип контрагента
	DeliveryAddress string `json:"delivery_address"` //адрес доставки
	Company         string `json:"company"`          //название компании
	Message         string `json:"message"`          //комментарий к заказу
	INN             string `json:"inn"`              //инн
	Phone           string `json:"phone"`            //телефон
	Email           string `json:"email"`            //Адрес эл. почты
	FullName        string `json:"fullname"`         //ФИО
	Index           string `json:"index"`            //индекс адреса доставки
	Street          string `json:"street"`           //улица доставки
	Dom             string `json:"dom"`              //дом доставки
	CDEKAddress     string `json:"cdek_address"`     //отформатированный адрес для сдэка
}
type MetriksData struct {
	Metriks []interface{} `json:"metriks"`
}
type MetriksDB struct {
	Raw    string
	First  int32
	Second int32
}

type PrintOrdersRequest struct {
	Orders    []PrintOrderDto `json:"orders"`
	CopyCount int32           `json:"copy_count"` // число копий, по умолчанию 1
	Format    string          `json:"formaty"`    // формат печати: A4, A5, A6, A7 (по умолчанию A4)
	Lang      string          `json:"lang"`       // язык печатной формы: RUS или ENG

}

// Заказ
type PrintOrderDto struct {
	OrderUUID  string `json:"order_uuidy"` // UUID заказа, обязателен если нет cdek_number
	CdekNumber int64  `json:"cdek_number"` // номер заказа СДЭК, обязателен если нет order_uuid
}
type RootEntityDto struct {
	Entity          EntityDto          `json:"entity"`
	Requests        []RequestDto1      `json:"requests"`         // информация о запросах
	RelatedEntities []RelatedEntityDto `json:"related_entities"` // связанные сущности
}

// Информация о сущности
type EntityDto struct {
	UUID string `json:"uuid"`
}

// Информация о запросе
type RequestDto1 struct {
	RequestUUID string       `json:"request_uuid"`
	Type        string       `json:"type"`      // CREATE, UPDATE, DELETE, AUTH, GET, CREATE_CLIENT_RETURN
	DateTime    string       `json:"date_time"` // <date-time>
	State       string       `json:"state"`     // ACCEPTED, WAITING, SUCCESSFUL, INVALID
	Errors      []ErrorDto1  `json:"errors"`    // ошибки
	Warnings    []WarningDto `json:"warnings"`  // предупреждения
}

// Ошибка
type ErrorDto1 struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Предупреждение
type WarningDto struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Связанная сущность
type RelatedEntityDto struct {
	UUID       string `json:"uuid"`         // идентификатор
	Type       string `json:"type"`         // return_order, direct_order, ...
	URL        string `json:"url"`          // только для waybill и barcode
	CreateTime string `json:"create_time"`  // <date-time>
	CdekNumber string `json:"cdek_numbery"` // номер заказа
	Date       string `json:"date"`         // дата доставки
	TimeFrom   string `json:"time_from"`    // время "с"
	TimeTo     string `json:"time_to"`      // время "по"
}

/*
func testzakaz() RootEntityDto {
	var phone PhoneDto
	var item1 Item
	var itemtest []Item
	var zakaz SkladDver
	var mesto Packages
	zakaz.Number = "1234"

	zakaz.Recipient.ContragentType = "LEGAL_ENTITY"
	zakaz.Recipient.Company = "ООО ТЕСТ"
	zakaz.Recipient.Tin = "7801223344"
	phone.Number = "+79115557744"
	zakaz.Recipient.Phones = append(zakaz.Recipient.Phones, phone)
	zakaz.Recipient.Name = "ИВАНОВ ИВАН ИВАНОВИЧ"

	zakaz.Type = 1
	//zakaz.TariffCode = 111
	zakaz.TariffCode = 137 //Посылка склад-дверь
	zakaz.ShipmentPoint = "SPB55"
	//zakaz.DeliveryPoint = "MSK537"
	zakaz.ToLocation.Address = "125504, г Москва, г. Москва, Коровинское шоссе, д 1А, 12"
	zakaz.ToLocation.Code = 44

	mesto.Number = "1"
	mesto.Weight = 0.6
	mesto.Height = 25
	mesto.Length = 17
	mesto.Width = 7
	item1.Amount = 1
	item1.WareKey = "2023"
	item1.Cost = 100
	item1.Payment.Value = 0
	item1.Name = "РЕГУЛЯТОР СКОРОСТИ"
	item1.Weight = 0.4
	itemtest = append(itemtest, item1)
	mesto.Items = append(mesto.Items, itemtest...)
	zakaz.Packages = append(zakaz.Packages, mesto)
	zakaz.Print = "BARCODE"
	//zakaz.ToLocation.Address = cdekaddr

	return SendZakaz(zakaz)
}
*/
