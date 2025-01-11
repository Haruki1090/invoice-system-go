package main

// 顧客情報の構造体
type Customer struct {
	Name    string
	Address string
	Email   string
}

// 請求項目の構造体
type Invoice struct {
	Description string
	UnitPrice   float64
	Quantity    int
}
