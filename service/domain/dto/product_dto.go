package dto

import "time"

type ProductResponse struct {
	Id            int                    `json:"id"`
	Name          string                 `json:"name"`
	Description   string                 `json:"description"`
	Weight        float64                `json:"weight"`
	Price         float64                `json:"price"`
	OrderMinQty   int                    `json:"order_min_qty"`
	OrderMaxQty   int                    `json:"order_max_qty"`
	Status        int                    `json:"status"`
	Stock         int                    `json:"stock"`
	Discount      float64                `json:"discount"`
	Category      CategoryResponse       `json:"category,omitempty"`
	ProductImages []ProductImageResponse `json:"product_images,omitempty"`
	CreatedAt     time.Time              `json:"created_at,omitempty"`
	UpdatedAt     time.Time              `json:"updated_at,omitempty"`
}

type ProductImageResponse struct {
	Id        int       `json:"id"`
	ProductId int       `json:"product_id,omitempty"`
	ImageUrl  string    `json:"image_url,omitempty"`
	MainImage int       `json:"main_image,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
