package requests

type HeaderItem struct {
	Header
	To_PurchaseOrderItem `json:"to_PurchaseOrderItem"`
}

type To_PurchaseOrderItem struct {
	Results []Item `json:"results"`
}
