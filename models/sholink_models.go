package models

import "time"

type ShortLink struct {
  ID               uint           `gorm:"primary_key"`
  ShortCode        string         `gorm:"size:100;uniqueIndex"`
  URL              string         `gorm:"size:255"`
  CreatedAt        time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
  UpdatedAt        time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
}