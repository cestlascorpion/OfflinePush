package core

type PushConfig struct {
	Mongo struct {
		Name            string `json:"name"`
		Url             string `json:"url"`
		DataBase        string `json:"database"`
		AuthCollection  string `json:"auth_collection"`
		StatsCollection string `json:"stats_collection"`
		PoolSize        int    `json:"pool_size"`
	} `json:"mongo"`
	GeTui struct {
		AgentId      string `json:"agent_id"`
		BundleId     string `json:"bundle_id"`
		AppId        string `json:"app_id"`
		AppKey       string `json:"app_key"`
		MasterSecret string `json:"master_secret"`
	} `json:"getui"`
	Apns struct {
		AgentId  string `json:"agent_id"`
		BundleId string `json:"bundle_id"`
		Env      string `json:"env"`
		Key      string `json:"key"`
		KeyId    string `json:"key_id"`
		TeamId   string `json:"team_id"`
	} `json:"apns"`
}
