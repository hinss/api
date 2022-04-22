package v1

import (
	metav1 "github.com/hinss/api/pkg/meta/v1"
)

type Goods struct {
	metav1.ObjectMeta

	CategoryId      int     `json:"category" gorm:"category_id;not null"`
	BrandId         int     `json:"brand" gorm:"brand_id;not null"`
	OnSale          uint    `json:"onSale" gorm:"on_sale;not null"`
	GoodsSn         string  `json:"goods_sn" gorm:"on_sale;not null"`
	Name            string  `json:"name" gorm:"name;not null"`
	ClickNum        int     `json:"clickNum" gorm:"click_num;not null"`
	SoldNum         int     `json:"soldNum" gorm:"sold_num;not null"`
	FavNum          int     `json:"favNum" gorm:"fav_num;not null"`
	Stocks          int     `json:"stocks" gorm:"stocks;not null"`
	MarketPrice     float64 `json:"marketPrice" gorm:"market_price;not null"`
	ShopPrice       float64 `json:"shopPrice" gorm:"shop_price;not null"`
	GoodsBrief      string  `json:"goodsBrief" gorm:"goods_brief;not null"`
	ShipFree        uint    `json:"shipFree" gorm:"ship_free;not null"`
	Images          string  `json:"images" gorm:"images;not null"`
	DescImages      string  `json:"descImages" gorm:"desc_images;not null"`
	GoodsFrontImage string  `json:"goodsFrontImage" gorm:"goods_front_image;not null"`
	IsNew           uint    `json:"isNew" gorm:"is_new;not null"`
	IsHot           uint    `json:"isHot" gorm:"is_hot;not null"`
}

func (g *Goods) TableName() string {
	return "goods"
}
