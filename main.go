package main

// 顧客情報の構造体
type Customer struct {
	Name    string
	Address string
	Email   string
}

// 請求項目の構造体
type InvoiceItem struct {
	Description string
	UnitPrice   float64
	Quantity    int
}

// 小計を計算するメソッド
func (item InvoiceItem) Total() float64 {
	return item.UnitPrice * float64(item.Quantity) // 単価 * 数量
}

// 請求書の構造体
type Invoice struct {
	Customer    Customer
	Item        []InvoiceItem
	IvoiceTotal float64
}

// 合計金額を計算するメソッド
func (inv *Invoice) CalculateTotal() {
	var total float64
	for _, item := range inv.Item {
		total += item.Total()
	}
	inv.IvoiceTotal = total
}
