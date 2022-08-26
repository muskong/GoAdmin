package logic

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoCore/config"
	"github.com/muskong/GoPkg/zaplog"
)

type _productService struct {
	Logic
}

var ProductService = &_productService{}

func (l *_productService) List(page Page) (result Result, err error) {

	productData, productCount, err := entity.ProductServiceEntity.GetProductServices(page.getOffset(), page.getLimit())
	if len(productData) <= 0 || err != nil {
		err = errors.New("产品服务商无数据")
		return
	}

	result.Data = productData
	result.Pagination.Page = page
	result.Pagination.Total = productCount

	return
}

func (l *_productService) Create(data map[string]string) error {
	entityData := new(entity.ProductService)
	entityData.Class = data["Class"]
	entityData.Title = data["Title"]
	entityData.Status = entity.ProductServiceEntity.StatusDeny()
	entityData.Type = entity.ProductServiceEntity.TypeApi()
	delete(data, "Class")
	delete(data, "Title")
	entityData.Content = data

	err := entity.ProductServiceEntity.Insert(entityData)
	if entityData.Id <= 0 || err != nil {
		return errors.New("新增产品服务商失败")
	}

	l.Log("新增产品服务商", data)
	return err
}
func (l *_productService) Update(data entity.ProductService) (err error) {
	err = entity.ProductServiceEntity.Delete(data.Id)
	if err != nil {
		return errors.New("更新产品服务商失败")
	}
	err = entity.ProductServiceEntity.Insert(&data)
	if data.Id <= 0 || err != nil {
		return errors.New("更新产品服务商失败")
	}

	l.Log("更新产品服务商", data)
	return err
}
func (l *_productService) Delete(productId int) error {
	err := entity.ProductServiceEntity.Delete(productId)
	if err != nil {
		return errors.New("删除产品服务商失败")
	}
	l.Log("删除产品服务商", productId)
	return err
}

func (l *_productService) PluginList() ([]map[string]any, error) {

	dir := config.App.GetString("plugins.path")

	plugins := []map[string]any{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return nil
		}
		if f.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".so") {
			return nil
		}

		plugins = append(plugins, map[string]any{
			"FileName":    f.Name(),
			"FileSize":    f.Size(),
			"FileModTime": f.ModTime().Format("2006-01-02 15:04:05"),
		})

		// PluginTest(path)
		// PluginGaoBei(path)
		// api := Plugin(path)
		// if api == nil {
		// 	zaplog.Sugar.Info("err")
		// 	return nil
		// }
		// var info []any
		// err = api.Info(&info)
		// // ([]ApiInfo)(info)
		// // zaplog.Sugar.Infof("--%+v--", info)
		// // zaplog.Sugar.Infof("--%#v--", info)
		// // for _, v := range info {
		// // 	zaplog.Sugar.Infof("--%+v--", v)
		// // 	zaplog.Sugar.Infof("--%#v--", v)
		// // 	apiinfo := v.(*ApiInfo)
		// // 	plugins[f.Name()] = append(plugins[f.Name()], *apiinfo)
		// // }
		// plugins[f.Name()] = append(plugins[path], info...)
		return err
	})
	return plugins, err
}

func (l *_productService) Plugin(fileName string) ([]any, error) {
	dir := config.App.GetString("plugins.path")
	api := Plugin(dir + fileName)
	if api == nil {
		zaplog.Sugar.Info("err")
		return nil, nil
	}
	var info []any
	err := api.Info(&info)

	return info, err
}
