package sap_api_input_reader

type EC_MC struct {
	ConnectionKey  string `json:"connection_key"`
	Result         bool   `json:"result"`
	RedisKey       string `json:"redis_key"`
	Filepath       string `json:"filepath"`
	SalesQuotation struct {
		SalesQuotation    string `json:"document_no"`
		Plant             string `json:"deliver_to"`
		QuotationQuantity string `json:"quantity"`
		PickedQuantity    string `json:"picked_quantity"`
		NetPriceAmount    string `json:"price"`
		Batch             string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema          string `json:"api_schema"`
	MaterialCode       string `json:"material_code"`
	Supplier           string `json:"plant/supplier"`
	Stock              string `json:"stock"`
	SalesQuotationType string `json:"document_type"`
	SQNumber           string `json:"document_no"`
	PlannedDate        string `json:"planned_date"`
	ValidatedDate      string `json:"validated_date"`
	Deleted            bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey  string `json:"connection_key"`
	Result         bool   `json:"result"`
	RedisKey       string `json:"redis_key"`
	Filepath       string `json:"filepath"`
	SalesQuotation struct {
		SalesQuotation                 string  `json:"SalesQuotation"`
		SalesQuotationType             *string `json:"SalesQuotationType"`
		SalesOrganization              *string `json:"SalesOrganization"`
		DistributionChannel            *string `json:"DistributionChannel"`
		OrganizationDivision           *string `json:"OrganizationDivision"`
		SalesGroup                     *string `json:"SalesGroup"`
		SalesOffice                    *string `json:"SalesOffice"`
		SalesDistrict                  *string `json:"SalesDistrict"`
		SoldToParty                    *string `json:"SoldToParty"`
		CreationDate                   *string `json:"CreationDate"`
		LastChangeDate                 *string `json:"LastChangeDate"`
		PurchaseOrderByCustomer        *string `json:"PurchaseOrderByCustomer"`
		CustomerPurchaseOrderDate      *string `json:"CustomerPurchaseOrderDate"`
		SalesQuotationDate             *string `json:"SalesQuotationDate"`
		TotalNetAmount                 *string `json:"TotalNetAmount"`
		TransactionCurrency            *string `json:"TransactionCurrency"`
		SDDocumentReason               *string `json:"SDDocumentReason"`
		PricingDate                    *string `json:"PricingDate"`
		RequestedDeliveryDate          *string `json:"RequestedDeliveryDate"`
		ShippingCondition              *string `json:"ShippingCondition"`
		CompleteDeliveryIsDefined      *bool   `json:"CompleteDeliveryIsDefined"`
		ShippingType                   *string `json:"ShippingType"`
		HeaderBillingBlockReason       *string `json:"HeaderBillingBlockReason"`
		DeliveryBlockReason            *string `json:"DeliveryBlockReason"`
		BindingPeriodValidityStartDate *string `json:"BindingPeriodValidityStartDate"`
		BindingPeriodValidityEndDate   *string `json:"BindingPeriodValidityEndDate"`
		HdrOrderProbabilityInPercent   *string `json:"HdrOrderProbabilityInPercent"`
		ExpectedOrderNetAmount         *string `json:"ExpectedOrderNetAmount"`
		IncotermsClassification        *string `json:"IncotermsClassification"`
		CustomerPaymentTerms           *string `json:"CustomerPaymentTerms"`
		CustomerPriceGroup             *string `json:"CustomerPriceGroup"`
		PriceListType                  *string `json:"PriceListType"`
		PaymentMethod                  *string `json:"PaymentMethod"`
		CustomerTaxClassification1     *string `json:"CustomerTaxClassification1"`
		ReferenceSDDocument            *string `json:"ReferenceSDDocument"`
		ReferenceSDDocumentCategory    *string `json:"ReferenceSDDocumentCategory"`
		SalesQuotationApprovalReason   *string `json:"SalesQuotationApprovalReason"`
		SalesDocApprovalStatus         *string `json:"SalesDocApprovalStatus"`
		OverallSDProcessStatus         *string `json:"OverallSDProcessStatus"`
		TotalCreditCheckStatus         *string `json:"TotalCreditCheckStatus"`
		OverallSDDocumentRejectionSts  *string `json:"OverallSDDocumentRejectionSts"`
		SalesQuotationItem             []struct {
			SalesQuotationItem            string  `json:"SalesQuotationItem"`
			SalesQuotationItemCategory    *string `json:"SalesQuotationItemCategory"`
			SalesQuotationItemText        *string `json:"SalesQuotationItemText"`
			PurchaseOrderByCustomer       *string `json:"PurchaseOrderByCustomer"`
			Material                      *string `json:"Material"`
			MaterialByCustomer            *string `json:"MaterialByCustomer"`
			RequestedQuantity             *string `json:"RequestedQuantity"`
			RequestedQuantityUnit         *string `json:"RequestedQuantityUnit"`
			ItemOrderProbabilityInPercent *string `json:"ItemOrderProbabilityInPercent"`
			ItemGrossWeight               *string `json:"ItemGrossWeight"`
			ItemNetWeight                 *string `json:"ItemNetWeight"`
			ItemWeightUnit                *string `json:"ItemWeightUnit"`
			ItemVolume                    *string `json:"ItemVolume"`
			ItemVolumeUnit                *string `json:"ItemVolumeUnit"`
			TransactionCurrency           *string `json:"TransactionCurrency"`
			NetAmount                     *string `json:"NetAmount"`
			MaterialGroup                 *string `json:"MaterialGroup"`
			MaterialPricingGroup          *string `json:"MaterialPricingGroup"`
			Batch                         *string `json:"Batch"`
			Plant                         *string `json:"Plant"`
			IncotermsClassification       *string `json:"IncotermsClassification"`
			CustomerPaymentTerms          *string `json:"CustomerPaymentTerms"`
			ProductTaxClassification1     *string `json:"ProductTaxClassification1"`
			SalesDocumentRjcnReason       *string `json:"SalesDocumentRjcnReason"`
			WBSElement                    *string `json:"WBSElement"`
			ProfitCenter                  *string `json:"ProfitCenter"`
			ReferenceSDDocument           *string `json:"ReferenceSDDocument"`
			ReferenceSDDocumentItem       *string `json:"ReferenceSDDocumentItem"`
			SDProcessStatus               *string `json:"SDProcessStatus"`
		} `json:"SalesQuotationItem"`
	} `json:"SalesQuotation"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	SQNumber  string   `json:"sales_quotation"`
	Deleted   bool     `json:"deleted"`
}
