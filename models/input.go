package models

type UrlStore struct {
	ID        uint64 `gorm:"column:id;type:bigint;not null;primaryKey" json:"id"`
	LongUrl   string `gorm:"column:long_url;type:varchar(200);not null" json:"long_url"`
	ShortUrl  string `gorm:"column:short_url;type:varchar(200;not null" json:"short_url"`
	CreatedAt uint64 `gorm:"column:created_at;type:timestamp;" json:"createdAt"`
}

type ConverterInput struct {
	Url string `json:"url"`
}
