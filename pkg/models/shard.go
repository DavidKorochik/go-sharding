package models

type Shard struct {
	ShardId   string            `json:"shard_id"`
	ShardName string            `json:"shard_name"`
	Metadata  map[string]string `json:"metadata"`
}
