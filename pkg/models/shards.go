package models

type ShardsList struct {
	Shards    []*Shard `json:"shards"`
	NumShards int      `json:"num_shards"`
}
