package models

import (
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"time"
)

type Shard struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	ShardKey  string            `json:"shard_key"`
	Nodes     []*snowflake.Node `json:"nodes"`
	CreatedAt time.Time         `json:"created_at"'`
	UpdatedAt time.Time         `json:"updated_at"'`

	*gorm.DB
}
