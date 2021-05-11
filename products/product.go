package products

type Product struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	StockAmount  int64   `json:"stock_amount"`
	SupplierCost float64 `json:"supplier_cost"`
	SellPrice    float64 `json:"sell_price"`
	CompanyID    uint    `json:"company_id"`
}
