package entity

type Order struct {
	ID            string
	Investor      *Investor
	Asset         *Asset
	Shares        int
	PendingShares int
	Price         float64
	OrderType     OrderType
	Status        StatusOrder
	Transactions  []*Transaction
}

type StatusOrder uint8

const (
	Open StatusOrder = iota
	Closed
)

type OrderType uint8

const (
	Buy OrderType = iota
	Sell
)

func NewOrder(orderID string, investor *Investor, asset *Asset, shares int, price float64, orderType OrderType) *Order {
	return &Order{
		ID:            orderID,
		Investor:      investor,
		Asset:         asset,
		Shares:        shares,
		PendingShares: shares,
		Price:         price,
		OrderType:     orderType,
		Status:        Open,
		Transactions:  []*Transaction{},
	}
}
