package service

import (
	"github.com/DavidKorochik/go-sharding/components/shard"
	"github.com/DavidKorochik/go-sharding/pkg/models"
	"gorm.io/gorm"
)

type ShardService struct {
	Shard  *models.Shard
	Shards []*models.Shard
}

func NewShardService() shard.Service {
	return &ShardService{}
}

func (svc *ShardService) RegisterShards(tables []*gorm.TableType) {
}

func (svc *ShardService) GetShard(key string) *models.Shard {
	// We could perform a binary search for more optimization, but the Shards array is not sorted
	// Also, it's not worth to sort the array of shards, because we'll only use it sorted for the purpose of the binary search
	for _, singleShard := range svc.Shards {
		if singleShard.ShardKey == key {
			return singleShard
		}
	}

	return nil
}

func (svc *ShardService) GetShardsList() []*models.Shard {
	return svc.Shards
}

// GetNumberOfShards - Monitor the number of the shards we have
func (svc *ShardService) GetNumberOfShards() int {
	return len(svc.Shards)
}
