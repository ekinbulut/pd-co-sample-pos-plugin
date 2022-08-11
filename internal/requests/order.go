package requests

import "time"

type Order struct {
	Token    string `json:"token"`
	Code     string `json:"code"`
	Comments struct {
		CustomerComment string `json:"customerComment"`
		VendorComment   string `json:"vendorComment"`
	} `json:"comments"`
	CreatedAt time.Time `json:"createdAt"`
	Customer  struct {
		Email                  string        `json:"email"`
		FirstName              string        `json:"firstName"`
		LastName               string        `json:"lastName"`
		MobilePhone            string        `json:"mobilePhone"`
		Code                   string        `json:"code"`
		ID                     string        `json:"id"`
		MobilePhoneCountryCode string        `json:"mobilePhoneCountryCode"`
		Flags                  []interface{} `json:"flags"`
	} `json:"customer"`
	Delivery struct {
		Address struct {
		} `json:"address"`
		ExpectedDeliveryTime time.Time `json:"expectedDeliveryTime"`
		ExpressDelivery      bool      `json:"expressDelivery"`
		RiderPickupTime      time.Time `json:"riderPickupTime"`
	} `json:"delivery"`
	Discounts []struct {
	} `json:"discounts"`
	ExpeditionType  string    `json:"expeditionType"`
	ExpiryDate      time.Time `json:"expiryDate"`
	ExtraParameters struct {
		Property1 string `json:"property1"`
		Property2 string `json:"property2"`
	} `json:"extraParameters"`
	LocalInfo struct {
		CountryCode            string `json:"countryCode"`
		CurrencySymbol         string `json:"currencySymbol"`
		Platform               string `json:"platform"`
		PlatformKey            string `json:"platformKey"`
		CurrencySymbolPosition string `json:"currencySymbolPosition"`
		CurrencySymbolSpaces   string `json:"currencySymbolSpaces"`
		DecimalDigits          string `json:"decimalDigits"`
		DecimalSeparator       string `json:"decimalSeparator"`
		Email                  string `json:"email"`
		Phone                  string `json:"phone"`
		ThousandsSeparator     string `json:"thousandsSeparator"`
		Website                string `json:"website"`
	} `json:"localInfo"`
	Payment struct {
		Status              string `json:"status"`
		Type                string `json:"type"`
		RemoteCode          string `json:"remoteCode"`
		RequiredMoneyChange string `json:"requiredMoneyChange"`
		VatID               string `json:"vatId"`
		VatName             string `json:"vatName"`
	} `json:"payment"`
	Test               bool        `json:"test"`
	ShortCode          string      `json:"shortCode"`
	PreOrder           bool        `json:"preOrder"`
	Pickup             interface{} `json:"pickup"`
	PlatformRestaurant struct {
		ID string `json:"id"`
	} `json:"platformRestaurant"`
	Price struct {
		DeliveryFees                     []interface{} `json:"deliveryFees"`
		GrandTotal                       string        `json:"grandTotal"`
		MinimumDeliveryValue             string        `json:"minimumDeliveryValue"`
		PayRestaurant                    string        `json:"payRestaurant"`
		RiderTip                         string        `json:"riderTip"`
		SubTotal                         string        `json:"subTotal"`
		VatTotal                         string        `json:"vatTotal"`
		Comission                        string        `json:"comission"`
		ContainerCharge                  string        `json:"containerCharge"`
		DeliveryFee                      string        `json:"deliveryFee"`
		CollectFromCustomer              string        `json:"collectFromCustomer"`
		DiscountAmountTotal              string        `json:"discountAmountTotal"`
		DeliveryFeeDiscount              string        `json:"deliveryFeeDiscount"`
		ServiceFeePercent                string        `json:"serviceFeePercent"`
		ServiceFeeTotal                  string        `json:"serviceFeeTotal"`
		ServiceTax                       int           `json:"serviceTax"`
		ServiceTaxValue                  int           `json:"serviceTaxValue"`
		DifferenceToMinimumDeliveryValue string        `json:"differenceToMinimumDeliveryValue"`
		VatVisible                       bool          `json:"vatVisible"`
		VatPercent                       string        `json:"vatPercent"`
	} `json:"price"`
	Products []struct {
	} `json:"products"`
	CorporateOrder  bool   `json:"corporateOrder"`
	CorporateTaxID  string `json:"corporateTaxId"`
	IntegrationInfo struct {
	} `json:"integrationInfo"`
	MobileOrder bool          `json:"mobileOrder"`
	WebOrder    bool          `json:"webOrder"`
	Vouchers    []interface{} `json:"vouchers"`
}
