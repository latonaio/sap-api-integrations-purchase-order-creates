package sap_api_input_reader

import (
	"sap-api-integrations-purchase-order-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.PurchaseOrder
	return &requests.Header{
		PurchaseOrder:                  data.PurchaseOrder,
		CompanyCode:                    data.CompanyCode,
		PurchaseOrderType:              data.PurchaseOrderType,
		PurchasingProcessingStatus:     data.PurchasingProcessingStatus,
		CreationDate:                   data.CreationDate,
		LastChangeDateTime:             data.LastChangeDateTime,
		Supplier:                       data.Supplier,
		Language:                       data.Language,
		PaymentTerms:                   data.PaymentTerms,
		PurchasingOrganization:         data.PurchasingOrganization,
		PurchasingGroup:                data.PurchasingGroup,
		PurchaseOrderDate:              data.PurchaseOrderDate,
		DocumentCurrency:               data.DocumentCurrency,
		ExchangeRate:                   data.ExchangeRate,
		ValidityStartDate:              data.ValidityStartDate,
		ValidityEndDate:                data.ValidityEndDate,
		SupplierRespSalesPersonName:    data.SupplierRespSalesPersonName,
		SupplierPhoneNumber:            data.SupplierPhoneNumber,
		SupplyingPlant:                 data.SupplyingPlant,
		IncotermsClassification:        data.IncotermsClassification,
		ManualSupplierAddressID:        data.ManualSupplierAddressID,
		AddressName:                    data.AddressName,
		AddressCityName:                data.AddressCityName,
		AddressFaxNumber:               data.AddressFaxNumber,
		AddressPostalCode:              data.AddressPostalCode,
		AddressStreetName:              data.AddressStreetName,
		AddressPhoneNumber:             data.AddressPhoneNumber,
		AddressRegion:                  data.AddressRegion,
		AddressCountry:                 data.AddressCountry,
	}
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	dataPurchaseOrder := sdc.PurchaseOrder
	data := sdc.PurchaseOrder.PurchaseOrderItem
	return &requests.Item{
		PurchaseOrder:                  dataPurchaseOrder.PurchaseOrder,
        PurchaseOrderItem:              data.PurchaseOrderItem,
		PurchaseOrderItemText:          data.PurchaseOrderItemText,
		Plant:                          data.Plant,
		StorageLocation:                data.StorageLocation,
		MaterialGroup:                  data.MaterialGroup,
		PurchasingInfoRecord:           data.PurchasingInfoRecord,
		SupplierMaterialNumber:         data.SupplierMaterialNumber,
		OrderQuantity:                  data.OrderQuantity,
		PurchaseOrderQuantityUnit:      data.PurchaseOrderQuantityUnit,
		OrderPriceUnit:                 data.OrderPriceUnit,
		DocumentCurrency:               data.DocumentCurrency,
		NetPriceAmount:                 data.NetPriceAmount,
		NetPriceQuantity:               data.NetPriceQuantity,
		TaxCode:                        data.TaxCode,
		OverdelivTolrtdLmtRatioInPct:   data.OverdelivTolrtdLmtRatioInPct,
		UnlimitedOverdeliveryIsAllowed: data.UnlimitedOverdeliveryIsAllowed,
		UnderdelivTolrtdLmtRatioInPct:  data.UnderdelivTolrtdLmtRatioInPct,
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
		ItemNetWeight:                  data.ItemNetWeight,
		ItemWeightUnit:                 data.ItemWeightUnit,
		IncotermsClassification:        data.IncotermsClassification,
		PurchaseRequisition:            data.PurchaseRequisition,
		PurchaseRequisitionItem:        data.PurchaseRequisitionItem,
		RequisitionerName:              data.RequisitionerName,
		Material:                       data.Material,
		InternationalArticleNumber:     data.InternationalArticleNumber,
		DeliveryAddressID:              data.DeliveryAddressID,
		DeliveryAddressName:            data.DeliveryAddressName,
		DeliveryAddressPostalCode:      data.DeliveryAddressPostalCode,
		DeliveryAddressStreetName:      data.DeliveryAddressStreetName,
		DeliveryAddressCityName:        data.DeliveryAddressCityName,
		DeliveryAddressRegion:          data.DeliveryAddressRegion,
		DeliveryAddressCountry:         data.DeliveryAddressCountry,
		PurchasingDocumentDeletionCode: data.PurchasingDocumentDeletionCode,
	}
}
