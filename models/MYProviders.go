package models

type MYProviderOrdersParams struct {
	Orders []MYOrder `json:"orders"`
}

type MYOrder struct {
	OrderID      int            `json:"order_id"`
	OrderDetails MYOrderDetails `json:"order_details"`
	Items        []MYItem       `json:"items"`
}

type MYOrderDetails struct {
	OrderLength float64   `json:"order_length"`
	OrderWidth  float64   `json:"order_width"`
	OrderHeight float64   `json:"order_height"`
	OrderWeight float64   `json:"order_weight"`
	OrderStatus string    `json:"order_status"`
	Consignee   Consignee `json:"consignee"`
	Pickup      Pickup    `json:"pickup"`
}

type MYItem struct {
	ItemID          int     `json:"item_product_id"`
	ItemDescription string  `json:"item_desc"`
	ItemCategory    string  `json:"item_category"`
	ItemSku         string  `json:"item_sku"`
	ItemQuantity    int     `json:"item_quantity"`
	ItemPrice       float64 `json:"item_price_value"`
	ItemCurrency    string  `json:"item_price_currency"`
}

type Consignee struct {
	ConsigneeID          int    `json:"consignee_id"`
	ConsigneeName        string `json:"consignee_name"`
	ConsigneePhoneNumber string `json:"consignee_phone_number"`
	ConsigneeCountry     string `json:"consignee_country"`
	ConsigneeAddress     string `json:"consignee_address"`
	ConsigneePostal      int    `json:"consignee_postal"`
	ConsigneeState       string `json:"consignee_state"`
	ConsigneeCity        string `json:"consignee_city"`
	ConsigneeProvince    string `json:"consignee_province"`
	ConsigneeEmail       string `json:"consignee_email"`
}

type Pickup struct {
	PickupID          int    `json:"pickup_id"`
	PickupName        string `json:"pickup_name"`
	PickupPhoneNumber string `json:"pickup_phone_number"`
	PickupCountry     string `json:"pickup_country"`
	PickupAddress     string `json:"pickup_address"`
	PickupPostal      int    `json:"pickup_postal"`
	PickupState       string `json:"pickup_state"`
	PickupCity        string `json:"pickup_city"`
	PickupProvince    string `json:"pickup_province"`
}

type MYOrderStatus struct {
	OrderID     int    `json:"order_id"`
	OrderStatus string `json:"order_status"`
}
