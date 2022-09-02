package sap_api_output_formatter

type PurchaseOrder struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	PurchaseOrder string `json:"purchase_order"`
	Deleted       string `json:"deleted"`
}

type Header struct {
	PurchaseOrder               string      `json:"PurchaseOrder"`
	CompanyCode                 string      `json:"CompanyCode"`
	PurchaseOrderType           string      `json:"PurchaseOrderType"`
	PurchasingProcessingStatus  string      `json:"PurchasingProcessingStatus"`
    CreationDate                string      `json:"CreationDate"`
    LastChangeDateTime          string      `json:"LastChangeDateTime"`
    Supplier                    string      `json:"Supplier"`
    Language                    string      `json:"Language"`
    PaymentTerms                string      `json:"PaymentTerms"`
    PurchasingOrganization      string      `json:"PurchasingOrganization"`
    PurchasingGroup             string      `json:"PurchasingGroup"`
    PurchaseOrderDate           string      `json:"PurchaseOrderDate"`
    DocumentCurrency            string      `json:"DocumentCurrency"`
    ExchangeRate                string      `json:"ExchangeRate"`
    ValidityStartDate           string      `json:"ValidityStartDate"`
    ValidityEndDate             string      `json:"ValidityEndDate"`
    SupplierRespSalesPersonName string      `json:"SupplierRespSalesPersonName"`
    SupplierPhoneNumber         string      `json:"SupplierPhoneNumber"`
    SupplyingPlant              string      `json:"SupplyingPlant"`
    IncotermsClassification     string      `json:"IncotermsClassification"`
    ManualSupplierAddressID     string      `json:"ManualSupplierAddressID"`
    AddressName                 string      `json:"AddressName"` 
    AddressCityName             string      `json:"AddressCityName"`
    AddressFaxNumber            string      `json:"AddressFaxNumber"`
    AddressPostalCode           string      `json:"AddressPostalCode"`
    AddressStreetName           string      `json:"AddressStreetName"`
    AddressPhoneNumber          string      `json:"AddressPhoneNumber"`
    AddressRegion               string      `json:"AddressRegion"`
    AddressCountry              string      `json:"AddressCountry"`
    ToItem                      string      `json:"to_PurchaseOrderItem"`
}

type Item struct {
	PurchaseOrder                  string `json:"PurchaseOrder"`
    PurchaseOrderItem              string `json:"PurchaseOrderItem"`
    PurchaseOrderItemText          string `json:"PurchaseOrderItemText"`
    Plant                          string `json:"Plant"`
    StorageLocation                string `json:"StorageLocation"`
    MaterialGroup                  string `json:"MaterialGroup"`
    PurchasingInfoRecord           string `json:"PurchasingInfoRecord"`
    SupplierMaterialNumber         string `json:"SupplierMaterialNumber"`
    OrderQuantity                  string `json:"OrderQuantity"`
    PurchaseOrderQuantityUnit      string `json:"PurchaseOrderQuantityUnit"`
    OrderPriceUnit                 string `json:"OrderPriceUnit"`
    DocumentCurrency               string `json:"DocumentCurrency"`
    NetPriceAmount                 string `json:"NetPriceAmount"`
    NetPriceQuantity               string `json:"NetPriceQuantity"`
    TaxCode                        string `json:"TaxCode"`
    OverdelivTolrtdLmtRatioInPct   string `json:"OverdelivTolrtdLmtRatioInPct"`
    UnlimitedOverdeliveryIsAllowed bool   `json:"UnlimitedOverdeliveryIsAllowed"`
    UnderdelivTolrtdLmtRatioInPct  string `json:"UnderdelivTolrtdLmtRatioInPct"`
    IsCompletelyDelivered          bool   `json:"IsCompletelyDelivered"`
    IsFinallyInvoiced              bool   `json:"IsFinallyInvoiced"`
    PurchaseOrderItemCategory      string `json:"PurchaseOrderItemCategory"`
    AccountAssignmentCategory      string `json:"AccountAssignmentCategory"`
    GoodsReceiptIsExpected         bool   `json:"GoodsReceiptIsExpected"`
    GoodsReceiptIsNonValuated      bool   `json:"GoodsReceiptIsNonValuated"`
    InvoiceIsExpected              bool   `json:"InvoiceIsExpected"`
    InvoiceIsGoodsReceiptBased     bool   `json:"InvoiceIsGoodsReceiptBased"`
    Customer                       string `json:"Customer"`
    SupplierIsSubcontractor        bool   `json:"SupplierIsSubcontractor"`
    ItemNetWeight                  string `json:"ItemNetWeight"`
    ItemWeightUnit                 string `json:"ItemWeightUnit"`
    IncotermsClassification        string `json:"IncotermsClassification"`
    PurchaseRequisition            string `json:"PurchaseRequisition"`
    PurchaseRequisitionItem        string `json:"PurchaseRequisitionItem"`
    RequisitionerName              string `json:"RequisitionerName"`
    Material                       string `json:"Material"`
    InternationalArticleNumber     string `json:"InternationalArticleNumber"`
    DeliveryAddressID              string `json:"DeliveryAddressID"`
    DeliveryAddressName            string `json:"DeliveryAddressName"`
    DeliveryAddressPostalCode      string `json:"DeliveryAddressPostalCode"`
    DeliveryAddressStreetName      string `json:"DeliveryAddressStreetName"`
    DeliveryAddressCityName        string `json:"DeliveryAddressCityName"`
    DeliveryAddressRegion          string `json:"DeliveryAddressRegion"`
    DeliveryAddressCountry         string `json:"DeliveryAddressCountry"`
    PurchasingDocumentDeletionCode string `json:"PurchasingDocumentDeletionCode"`
    ToItemScheduleLine             string `json:"to_ScheduleLine"`
    ToItemPricingElement           string `json:"to_PurchaseOrderPricingElement"`
    ToItemAccount                  string `json:"to_AccountAssignment"`
}
