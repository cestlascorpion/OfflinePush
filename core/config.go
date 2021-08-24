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
	TestApp struct {
		PushAgent    string `required:"true"`
		BundleId     string `required:"true"`
		AppId        string `required:"true"`
		AppKey       string `required:"true"`
		MasterSecret string `required:"true"`
		TimeoutSec   int    `default:"5"`
	}
}
