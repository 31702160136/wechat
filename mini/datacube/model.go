package datacube

type Datacube struct {
	appID  string
	secret string
}
type NameValueV2 struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
type NameValueV3 struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}
type KeyValueV2 struct {
	Key   int `json:"key"`
	Value int `json:"value"`
}
