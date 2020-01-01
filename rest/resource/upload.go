package resource

import (
	"github.com/cisordeng/beego/xenon"

	bResource "kylin/business/resource"
)

type Upload struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(Upload))
}

func (this *Upload) Resource() string {
	return "resource.upload"
}

func (this *Upload) Params() map[string][]string {
	return map[string][]string{
		"PUT": []string{
			"file",
		},
	}
}

func (this *Upload) Put() {
	file, fileHeader, _ := this.GetFile("file")

	resource := bResource.NewResource(file, fileHeader)
	data := bResource.EncodeResource(resource)
	this.ReturnJSON(data)
}
