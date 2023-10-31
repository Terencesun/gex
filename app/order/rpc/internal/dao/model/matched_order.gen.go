// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMatchedOrder = "matched_order"

// MatchedOrder mapped from table <matched_order>
type MatchedOrder struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:雪花算法id" json:"id"`
	MatchID      string `gorm:"column:match_id;not null;comment:撮合id" json:"match_id"`
	SymbolID     int32  `gorm:"column:symbol_id;not null;comment:交易对id" json:"symbol_id"`
	SymbolName   string `gorm:"column:symbol_name;not null;comment:交易对名称" json:"symbol_name"`
	TakerOrderID string `gorm:"column:taker_order_id;not null;comment:taker订单id" json:"taker_order_id"`
	MakerOrderID string `gorm:"column:maker_order_id;not null;comment:maker订单id" json:"maker_order_id"`
	Price        string `gorm:"column:price;not null;comment:价格" json:"price"`
	Qty          string `gorm:"column:qty;not null;comment:数量(基础币)" json:"qty"`
	Amount       string `gorm:"column:amount;not null;comment:金额（计价币）" json:"amount"`
	MatchTime    int64  `gorm:"column:match_time;comment:撮合时间" json:"match_time"`
	CreatedAt    int64  `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`
	UpdatedAt    int64  `gorm:"column:updated_at;not null;comment:修改时间" json:"updated_at"`
}

// TableName MatchedOrder's table name
func (*MatchedOrder) TableName() string {
	return TableNameMatchedOrder
}
