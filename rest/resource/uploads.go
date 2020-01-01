package resource

import (
	"github.com/cisordeng/beego/xenon"
	bResource "kylin/business/resource"
)

type Uploads struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(Uploads))
}

func (this *Uploads) Resource() string {
	return "resource.uploads"
}

func (this *Uploads) Params() map[string][]string {
	return map[string][]string{
		"PUT": []string{
			"files",
		},
	}
}

func (this *Uploads) Put() {
	fileHeaders, _ := this.GetFiles("files")

	resources := bResource.NewResources(fileHeaders)
	data := bResource.EncodeManyResource(resources)
	this.ReturnJSON(xenon.Map{
		"resources": data,
	})
}
