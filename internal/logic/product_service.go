package logic

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/muskong/GoAdmin/internal/entity"
	"github.com/muskong/GoCore/config"
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

func (l *_productService) Create(data entity.ProductService) error {
	err := entity.ProductServiceEntity.Insert(&data)
	if data.Id <= 0 || err != nil {
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

func (l *_productService) Install() ([]map[string]string, error) {
	dir := config.App.GetString("plugins.path")

	plugins := []map[string]string{}
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
		// log.Println(path)
		info := Plugin(path).Info()

		plugins = append(plugins, info)
		return nil
	})
	return plugins, err
}
