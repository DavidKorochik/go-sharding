package models

type Shards struct {
	ShardsList []*Shards `json:"shards_list"`
	NumShards  int       `json:"num_shards"`
}
