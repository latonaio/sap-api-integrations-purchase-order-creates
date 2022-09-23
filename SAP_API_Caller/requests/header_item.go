package requests

type HeaderItem struct {
	Header
	To_Item `json:"to_PurchaseOrderItem"`
}

type To_Item struct {
	Results []Item `json:"results"`
}
