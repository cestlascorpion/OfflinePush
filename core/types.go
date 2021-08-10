package core

type UniqueId struct {
	PushAgent string `json:"push_agent"`
	BundleId  string `json:"bundle_id"`
}

type AuthToken struct {
	Token    string `json:"token" bson:"token"`
	ExpireAt int64  `json:"expire_at" bson:"expire_at"`
}
