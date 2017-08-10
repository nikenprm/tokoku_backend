package repository

import "fmt"

type WalmartItem struct {
	ItemId          int64   `json:"itemId"`
	ParentItemId    int64   `json:"parentItem"`
	Name            string  `json:"name`
	SalePrice       float32 `json:"salePrice"`
	Upc             string  `json:"upc"`
	CategoryPath    string  `json:"categoryPath"`
	LongDescription string  `json:"longDescription"`
	BrandName       string  `json:"brandName"`
	ThumbnailImage  string  `json:"thumbnailImage"`
	MediumImage     string  `json:"mediumImage"`
	LargeImage      string  `json:"largeImage"`
	Size            string  `json:"size"`
	ProductUrl      string  `json:"productUrl"`
	CategoryNode    string  `json:"categoryNode"`
	Stock           string  `json:"stock"`
}

func (mgr *manager) AddWalmartItem(walmartItem *WalmartItem) (err error) {
	mgr.db.Create(&walmartItem)

	if errs := mgr.db.GetErrors(); len(errs) > 0 {
		err = errs[0]
	}
	return
}

func (mgr *manager) GetListOfWalmartItemByCategory(categoryID string, page int, perPage int) (listOfWalmartItem []WalmartItem, err error) {
	offset := (page - 1) * perPage
	fmt.Println("offset: ", offset)
	mgr.db.Where("category_node = ?", categoryID).Offset(offset).Limit(perPage).Find(&listOfWalmartItem)

	if errs := mgr.db.GetErrors(); len(errs) > 0 {
		err = errs[0]
	}
	return

}

func (mgr *manager) DeleteWalmartItem(categoryID string) (err error) {
	mgr.db.Where("category_node = ?", categoryID).Delete(WalmartItem{})

	if errs := mgr.db.GetErrors(); len(errs) > 0 {
		err = errs[0]
	}
	return

}
