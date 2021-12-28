package core

type PushConfig struct {
	Mongo struct {
		Name            string `default:"Mongo"`
		Url             string `required:"true"`
		DataBase        string `required:"true"`
		AuthCollection  string `required:"true"`
		StatsCollection string `required:"true"`
		PoolSize        int    `default:"100"`
	}
	GeTui struct {
		AgentId      string `required:"true"`
		BundleId     string `required:"true"`
		AppId        string `required:"true"`
		AppKey       string `required:"true"`
		MasterSecret string `required:"true"`
		TimeoutSec   int    `default:"5"`
	}
	Apns struct {
		AgentId    string `required:"true"`
		BundleId   string `required:"true"`
		Key        []byte `required:"true"`
		KeyId      string `required:"true"`
		TeamId     string `required:"true"`
		TimeoutSec int    `default:"10"`
	}
}
