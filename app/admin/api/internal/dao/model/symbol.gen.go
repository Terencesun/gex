// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "gorm.io/plugin/soft_delete"

const TableNameSymbol = "symbol"

// Symbol mapped from table <symbol>
type Symbol struct {
	ID            uint32                `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	SymbolName    string                `gorm:"column:symbol_name;not null" json:"symbol_name"`
	BaseCoinID    uint32                `gorm:"column:base_coin_id;not null;comment:基础币ID" json:"base_coin_id"`
	BaseCoinName  string                `gorm:"column:base_coin_name;not null;comment:基础币名称" json:"base_coin_name"`
	BaseCoinPrec  int32                 `gorm:"column:base_coin_prec;not null;comment:基础币精度" json:"base_coin_prec"`
	QuoteCoinID   uint32                `gorm:"column:quote_coin_id;not null;comment:计价币ID" json:"quote_coin_id"`
	QuoteCoinName string                `gorm:"column:quote_coin_name;not null;comment:计价币名称" json:"quote_coin_name"`
	QuoteCoinPrec int32                 `gorm:"column:quote_coin_prec;not null;comment:计价币精度" json:"quote_coin_prec"`
	CreatedAt     uint32                `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`
	UpdatedAt     uint32                `gorm:"column:updated_at;not null;comment:修改时间" json:"updated_at"`
	DeletedAt     soft_delete.DeletedAt `gorm:"softDelete:unix" json:"deleted_at"`
}

// TableName Symbol's table name
func (*Symbol) TableName() string {
	return TableNameSymbol
}