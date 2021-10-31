package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы)
type Money int64

// Currency представляет код валюты
type Currency string

// Category представляет собой категорию, в которой был совершен платеж
type Category string

// Status представляем собой статус платежа.
type Status string

// Предопределенные статусы платежей.
const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

// PAN представляет номер карты
type PAN string

// Payment представляет информацию о платеже.
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}

// Card представляет информацию о платежной карте
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

// PaymentSource источник оплаты
type PaymentSource struct {
	Type    string // "card"
	Number  string // номер вида "5058 xxxx xxxx 8888"
	Balance Money
}
