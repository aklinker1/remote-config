package backend

type JSON = map[string]interface{}
type Blob = map[string]JSON

type Repo interface {
	GetBlob() Blob
	GetApps() []string
	GetConfig(app string) JSON
	SaveConfig(app string, config JSON) error
}
