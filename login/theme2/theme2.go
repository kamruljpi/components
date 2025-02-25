package theme2

import (
	"github.com/kamruljpi/components/login"
	"github.com/kamruljpi/go-admin/template"
)

type Theme2 struct{
	*template.BaseComponent
}

func (*Theme2) GetAssetList() []string {
	return AssetsList
}

func (*Theme2) GetAsset(name string) ([]byte, error) {
	return Asset(name)
}

func (*Theme2) GetHTML() string {
	return List["login"]
}

func init() {
	login.Register("theme2", new(Theme2))
}
