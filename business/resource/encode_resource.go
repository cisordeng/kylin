package resource

import (
	"github.com/cisordeng/beego/xenon"
)

func EncodeResource(resource *Resource) xenon.Map {
	if resource == nil {
		return nil
	}

	mapResource := xenon.Map{
		"name": resource.Name,
		"type": resource.Type,
		"size": resource.Size,
		"path": resource.Path,
		"url": resource.Url,
	}
	return mapResource
}


func EncodeManyResource(resources []*Resource) []xenon.Map {
	mapResources := make([]xenon.Map, 0)
	for _, resource := range resources {
		mapResources = append(mapResources, EncodeResource(resource))
	}
	return mapResources
}