package resource

import (
	"fmt"
	"github.com/cisordeng/beego/xenon"
	"mime/multipart"
	"os"
	"strings"

	"github.com/cisordeng/beego"
)

type Resource struct {
	Name string
	Type string
	Size int64
	Path string
	Url string
}

func NewResource(file multipart.File, fileHeader *multipart.FileHeader) *Resource {
	instance := new(Resource)
	instance.Name = fileHeader.Filename
	instance.Type = fileHeader.Header["Content-Type"][0]
	instance.Size = fileHeader.Size
	resourceRoot :=  beego.AppConfig.String("resource::RESOURCE_ROOT")
	instance.Path = fmt.Sprintf("%s/%s/%s/%s", resourceRoot, beego.BConfig.AppName, strings.Split(instance.Type, "/")[0], instance.Name)
	err := os.MkdirAll(fmt.Sprintf("%s/%s/%s/", resourceRoot, beego.BConfig.AppName, strings.Split(instance.Type, "/")[0]), os.ModePerm)
	xenon.PanicNotNilError(err, "rest:mkdir failed", "创建文件夹失败")
	xenon.SaveToFile(file, instance.Path)
	resourceUrl :=  beego.AppConfig.String("resource::RESOURCE_URL")
	instance.Url = fmt.Sprintf("%s/%s/%s/%s", resourceUrl, beego.BConfig.AppName, strings.Split(instance.Type, "/")[0], instance.Name)
	return instance
}

func NewResources(fileHeaders []*multipart.FileHeader) []*Resource {
	resources := make([]*Resource, 0)
	for _, fileHeader := range fileHeaders {
		file, err := fileHeader.Open()
		xenon.PanicNotNilError(err, "rest:upload failed", "上传失败")
		resource := NewResource(file, fileHeader)
		resources = append(resources, resource)
	}
	return resources
}