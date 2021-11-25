package backend

type JSON = map[string]interface{}
type Blob = map[string]JSON

type Repo interface {
	GetBlob() (Blob, error)
	GetApps() ([]string, error)
	GetConfig(app string) (JSON, error)
	SaveConfig(app string, config JSON) (JSON, error)
	DeleteConfig(app string) (JSON, error)
}
