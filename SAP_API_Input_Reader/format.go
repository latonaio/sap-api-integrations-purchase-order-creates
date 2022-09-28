package sap_api_input_reader

import (
	"sap-api-integrations-purchase-order-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeaderItem() *requests.HeaderItem {
	data := sdc.PurchaseOrder
	header := sdc.ConvertToHeader()

	itemResults := make([]requests.Item, 0, len(data.PurchaseOrderItem))
	for i := range data.PurchaseOrderItem {
		itemResults = append(itemResults, *sdc.ConvertToItem(i))
	}

	return &requests.HeaderItem{
		Header: *header,
		ToItem: requests.ToItem{
			Results: itemResults,
		},
	}
}

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.PurchaseOrder
	return &requests.Header{
		PurchaseOrder:               data.PurchaseOrder,
		CompanyCode:                 data.CompanyCode,
		PurchaseOrderType:           data.PurchaseOrderType,
		PurchasingProcessingStatus:  data.PurchasingProcessingStatus,
		CreationDate:                data.CreationDate,
		Supplier:                    data.Supplier,
		Language:                    data.Language,
		PaymentTerms:                data.PaymentTerms,
		PurchasingOrganization:      data.PurchasingOrganization,
		PurchasingGroup:             data.PurchasingGroup,
		PurchaseOrderDate:           data.PurchaseOrderDate,
		DocumentCurrency:            data.DocumentCurrency,
		SupplierRespSalesPersonName: data.SupplierRespSalesPersonName,
		SupplierPhoneNumber:         data.SupplierPhoneNumber,
		SupplyingPlant:              data.SupplyingPlant,
		IncotermsClassification:     data.IncotermsClassification,
		ManualSupplierAddressID:     data.ManualSupplierAddressID,
		AddressName:                 data.AddressName,
		AddressCityName:             data.AddressCityName,
		AddressFaxNumber:            data.AddressFaxNumber,
		AddressPostalCode:           data.AddressPostalCode,
		AddressStreetName:           data.AddressStreetName,
		AddressPhoneNumber:          data.AddressPhoneNumber,
		AddressRegion:               data.AddressRegion,
	}
}

func (sdc *SDC) ConvertToItem(item int) *requests.Item {
	dataPurchaseOrder := sdc.PurchaseOrder
	data := sdc.PurchaseOrder.PurchaseOrderItem[item]
	itemPricingElementResults := make([]*requests.ItemPricingElement, 0, len(data.ItemPricingElement))
	for i := range data.ItemPricingElement {
		itemPricingElementResults = append(itemPricingElementResults, sdc.ConvertToItemPricingElement(item, i))
	}

	return &requests.Item{
		PurchaseOrder:                  dataPurchaseOrder.PurchaseOrder,
		PurchaseOrderItem:              data.PurchaseOrderItem,
		Plant:                          data.Plant,
		StorageLocation:                data.StorageLocation,
		MaterialGroup:                  data.MaterialGroup,
		PurchasingInfoRecord:           data.PurchasingInfoRecord,
		SupplierMaterialNumber:         data.SupplierMaterialNumber,
		OrderQuantity:                  data.OrderQuantity,
		DocumentCurrency:               data.DocumentCurrency,
		TaxCode:                        data.TaxCode,
		UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
		IsCompletelyDelivered:          data.IsCompletelyDelivered,
		IsFinallyInvoiced:              data.IsFinallyInvoiced,
		PurchaseOrderItemCategory:      data.PurchaseOrderItemCategory,
		AccountAssignmentCategory:      data.AccountAssignmentCategory,
		GoodsReceiptIsExpected:         data.GoodsReceiptIsExpected,
		GoodsReceiptIsNonValuated:      data.GoodsReceiptIsNonValuated,
		InvoiceIsExpected:              data.InvoiceIsExpected,
		InvoiceIsGoodsReceiptBased:     data.InvoiceIsGoodsReceiptBased,
		Customer:                       data.Customer,
		SupplierIsSubcontractor:        data.SupplierIsSubcontractor,
		IncotermsClassification:        data.IncotermsClassification,
		PurchaseRequisition:            data.PurchaseRequisition,
		PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
		RequisitionerName:              data.RequisitionerName,
		Material:                       data.Material,
		InternationalArticleNumber:     data.InternationalArticleNumber,
		PurchasingDocumentDeletionCode: data.PurchasingDocumentDeletionCode,
		ToItemPricingElement: requests.ToItemPricingElement{
			Results: itemPricingElementResults,
		},
	}
}

func (sdc *SDC) ConvertToItemPricingElement(item, itemPricingElement int) *requests.ItemPricingElement {
	dataPurchaseOrder := sdc.PurchaseOrder
	dataPurchaseOrderItem := sdc.PurchaseOrder.PurchaseOrderItem[item]
	data := sdc.PurchaseOrder.PurchaseOrderItem[item].ItemPricingElement[itemPricingElement]
	return &requests.ItemPricingElement{
		PurchaseOrder:               dataPurchaseOrder.PurchaseOrder,
		PurchaseOrderItem:           dataPurchaseOrderItem.PurchaseOrderItem,
		PricingDocument:             data.PricingDocument,
		PricingDocumentItem:         data.PricingDocumentItem,
		PricingProcedureStep:        data.PricingProcedureStep,
		PricingProcedureCounter:     data.PricingProcedureCounter,
		ConditionType:               data.ConditionType,
		ConditionRateValue:          data.ConditionRateValue,
		ConditionCurrency:           data.ConditionCurrency,
		PriceDetnExchangeRate:       data.PriceDetnExchangeRate,
		TransactionCurrency:         data.TransactionCurrency,
		ConditionAmount:             data.ConditionAmount,
		ConditionQuantityUnit:       data.ConditionQuantityUnit,
		ConditionQuantity:           data.ConditionQuantity,
		ConditionApplication:        data.ConditionApplication,
		PricingDateTime:             data.PricingDateTime,
		ConditionCalculationType:    data.ConditionCalculationType,
		ConditionBaseValue:          data.ConditionBaseValue,
		ConditionToBaseQtyNmrtr:     data.ConditionToBaseQtyNmrtr,
		ConditionToBaseQtyDnmntr:    data.ConditionToBaseQtyDnmntr,
		ConditionCategory:           data.ConditionCategory,
		PricingScaleType:            data.PricingScaleType,
		ConditionOrigin:             data.ConditionOrigin,
		IsGroupCondition:            data.IsGroupCondition,
		ConditionSequentialNumber:   data.ConditionSequentialNumber,
		ConditionInactiveReason:     data.ConditionInactiveReason,
		PricingScaleBasis:           data.PricingScaleBasis,
		ConditionScaleBasisValue:    data.ConditionScaleBasisValue,
		ConditionScaleBasisCurrency: data.ConditionScaleBasisCurrency,
		ConditionScaleBasisUnit:     data.ConditionScaleBasisUnit,
		ConditionIsManuallyChanged:  data.ConditionIsManuallyChanged,
		ConditionRecord:             data.ConditionRecord,
	}
}
