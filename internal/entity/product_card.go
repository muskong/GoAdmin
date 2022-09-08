package entity

import (
	"time"

	"github.com/muskong/GoPkg/gorm"
	"github.com/muskong/GoPkg/zaplog"
	gdb "gorm.io/gorm"
)

type (
	ProductCard struct {
		gorm.Model
		Title      string `json:"Title,omitempty" db:"title"`
		IconUrl    string `json:"IconUrl,omitempty" db:"icon_url"`
		Batch      string `json:"Batch,omitempty" db:"batch"`
		Single     string `json:"Single,omitempty" db:"single"`
		Status     string `json:"Status,omitempty" db:"status"`
		Regularity string `json:"Regularity,omitempty" db:"regularity"`
		Note       string `json:"Note,omitempty" db:"note"`
		Example    string `json:"Example,omitempty" db:"example"`
	}
	_productCard struct{}
)

var ProductCardEntity = new(_productCard)

func (*_productCard) db() *gdb.DB {
	return gorm.NewModel(&ProductCard{}).Where("deleted_at IS NULL")
}

func (e *_productCard) BatchAllow() string {
	return "allow"
}
func (e *_productCard) BatchDeny() string {
	return "deny"
}

func (e *_productCard) SingleAllow() string {
	return "allow"
}
func (e *_productCard) SingleDeny() string {
	return "deny"
}

func (e *_productCard) StatusAllow() string {
	return "allow"
}
func (e *_productCard) StatusDeny() string {
	return "deny"
}

func (e *_productCard) GetProductCard(uuid string) (*ProductCard, error) {
	var data ProductCard
	err := e.db().Where("uuid = ?", uuid).First(&data).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return &data, err
}

func (e *_productCard) GetProductCards(offset, limit int) (products []*ProductCard, count int64, err error) {
	err = e.db().Count(&count).Order("id desc").Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_productCard) Insert(card *ProductCard) (err error) {
	err = e.db().Create(card).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}

func (e *_productCard) Delete(cardId int) error {
	pc := ProductCard{}
	pc.DeletedAt = gorm.NullString(time.Now().Format("2006-01-02 15:04:05"))
	err := e.db().Where("id = ?", cardId).Updates(pc).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return err
}

func (e *_productCard) GetAll() (products []*ProductCard, err error) {
	err = e.db().Order("id desc").Find(&products).Error
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return
}
