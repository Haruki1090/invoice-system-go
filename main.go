package main

import "fmt"

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
	Customer     Customer
	Items        []InvoiceItem
	InvoiceTotal float64
}

// 合計金額を計算するメソッド
func (inv *Invoice) CalculateTotal() {
	var total float64
	for _, item := range inv.Items {
		total += item.Total()
	}
	inv.InvoiceTotal = total
}

func main() {
	// 顧客情報を作成
	customer := Customer{
		Name:    "Haruki",
		Address: "TOKYO",
		Email:   "zenn@sample.com",
	}

	// 請求項目を作成
	items := []InvoiceItem{
		{Description: "Web開発サービス", UnitPrice: 150000, Quantity: 1},
		{Description: "保守サポート", UnitPrice: 30000, Quantity: 3},
	}

	// 請求書を作成
	invoice := Invoice{
		Customer: customer,
		Items:    items,
	}

	// 合計金額を計算
	invoice.CalculateTotal()

	// 請求書を表示
	fmt.Printf("請求書\n顧客: %s\n住所: %s\nメール: %s\n\n", invoice.Customer.Name, invoice.Customer.Address, invoice.Customer.Email)
	fmt.Println("請求項目:")
	for _, item := range invoice.Items {
		fmt.Printf("- %s: ¥%.2f x %d = ¥%.2f\n", item.Description, item.UnitPrice, item.Quantity, item.Total())
	}
	fmt.Printf("\n合計金額: ¥%.2f\n", invoice.InvoiceTotal)
}
