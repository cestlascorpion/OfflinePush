package core

type UniqueId struct {
	PushAgent string `json:"push_agent"`
	BundleId  string `json:"bundle_id"`
}

type AuthToken struct {
	Token    string `json:"token"`
	ExpireAt int64  `json:"expire_at"`
}

type AuthDoc struct {
	PushAgent string `bson:"push_agent"`
	BundleId  string `bson:"bundle_id"`
	Token     string `bson:"token"`
	ExpireAt  int64  `bson:"expire_at"`
}
