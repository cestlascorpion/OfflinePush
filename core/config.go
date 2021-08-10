package core

type AuthConfig struct {
	Name  string `default:"push-auth"`
	Mongo struct {
		Name       string `default:"Mongo"`
		Url        string `required:"true"`
		DataBase   string `required:"true"`
		Collection string `required:"true"`
		PoolSize   int    `default:"100"`
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

type StatsConfig struct {
	Name  string `default:"push-stats"`
	Mongo struct {
		Name       string `default:"Mongo"`
		Url        string `required:"true"`
		DataBase   string `required:"true"`
		Collection string `required:"true"`
		PoolSize   int    `default:"100"`
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
