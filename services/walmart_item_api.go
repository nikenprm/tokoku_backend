package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"bitbucket.org/icehousecorp/inuit_service_backend/logger"

	"strings"

	"github.com/tokoku_backend/config"
	"github.com/tokoku_backend/repository"
)

type ItemListResponseJSON struct {
	ItemList []ItemResponse `json:"items"`
}

type ItemResponse struct {
	ItemId                    int64   `json:"itemId"`
	ParentItemId              int64   `json:"parentItemId"`
	Name                      string  `json:"name`
	SalePrice                 float32 `json:"salePrice"`
	Upc                       string  `json:"upc"`
	CategoryPath              string  `json:"categoryPath"`
	LongDescription           string  `json:"longDescription"`
	BrandName                 string  `json:"brandName"`
	ThumbnailImage            string  `json:"thumbnailImage"`
	MediumImage               string  `json:"mediumImage"`
	LargeImage                string  `json:"largeImage"`
	ProductTrackingUrl        string  `json:"productTrackingUrl"`
	NinetySevenCentShipping   bool    `json:"ninetySevenCentShipping"`
	StandardShipRate          float32 `json:"standardShipRate"`
	Size                      string  `json:"size"`
	MarketPlace               bool    `json:"marketPlace"`
	ShipToStore               bool    `json:"shipToStore"`
	FreeShipToStore           bool    `json:"freeShipToStore"`
	ProductUrl                string  `json:"productUrl"`
	CategoryNode              string  `json:"categoryNode"`
	Bundle                    bool    `json:"bundle"`
	Clearance                 bool    `json:"clearance"`
	PreOrder                  bool    `json:"preOrder"`
	Stock                     string  `json:"stock"`
	AddToCartUrl              string  `json:"addToCartUrl"`
	AffiliateAddToCartUrl     string  `json:"affiliateAddToCartUrl"`
	FreeShippingOver50Dollars bool    `json:"freeShippingOver50Dollars"`
	AvailableOnline           bool    `json:"availableOnline"`
}

func RetrieveWalmartItemsByCategory(categoryId string) {
	client := &http.Client{}
	categoryNode := strings.Split(categoryId, "_")
	fmt.Println(categoryId)
	fmt.Println("Node terakhir: ", categoryNode[2])
	req, err := http.NewRequest("GET", config.Config.Walmart.URL+"/paginated/items?category="+categoryId+"&apiKey="+config.Config.Walmart.APIKey, nil)
	if err != nil {
		logger.Printf("NewRequest:", err)
		return
	}
	//req.Header.Add("Authorization", config.Config.Jira.Auth)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Response:", err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	var itemListResponse ItemListResponseJSON

	err = json.Unmarshal([]byte(data), &itemListResponse)
	if err != nil {
		fmt.Println(err)
	}

	itemsInDB, err := repository.Mgr.GetListOfWalmartItemByCategory(categoryId, 0, -1)
	if err != nil {
		fmt.Println(err)
	}
	rowCount := len(itemsInDB)
	fmt.Println("di db: ", rowCount)

	jsonCount := len(itemListResponse.ItemList)
	fmt.Println("di API: ", jsonCount)
	//jsonCount := len(projects)

	if rowCount < jsonCount {
		if rowCount > 0 {
			//clean up data
			fmt.Println("Deleting & re-inserting data to DB...")
			err = repository.Mgr.DeleteWalmartItem(categoryId)
			if err != nil {
				fmt.Println(err)
			}
		}
		for i := 0; i < jsonCount; i++ {
			if strings.Contains(itemListResponse.ItemList[i].CategoryNode, categoryNode[2]) {
				item := &repository.WalmartItem{
					ItemId:          itemListResponse.ItemList[i].ItemId,
					ParentItemId:    itemListResponse.ItemList[i].ParentItemId,
					Name:            itemListResponse.ItemList[i].Name,
					SalePrice:       itemListResponse.ItemList[i].SalePrice,
					Upc:             itemListResponse.ItemList[i].Upc,
					CategoryPath:    itemListResponse.ItemList[i].CategoryPath,
					LongDescription: itemListResponse.ItemList[i].LongDescription,
					BrandName:       itemListResponse.ItemList[i].BrandName,
					ThumbnailImage:  itemListResponse.ItemList[i].ThumbnailImage,
					MediumImage:     itemListResponse.ItemList[i].MediumImage,
					LargeImage:      itemListResponse.ItemList[i].LargeImage,
					Size:            itemListResponse.ItemList[i].Size,
					ProductUrl:      itemListResponse.ItemList[i].ProductUrl,
					CategoryNode:    itemListResponse.ItemList[i].CategoryNode,
					Stock:           itemListResponse.ItemList[i].Stock,
				}
				if err := repository.Mgr.AddWalmartItem(item); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func RetrieveAllWalmartItems() {
	listOfWalmartCategory, err := repository.Mgr.GetListOfWalmartCategory()
	if err != nil {
		fmt.Println(err)
	}

	for _, walmartCategory := range listOfWalmartCategory {
		fmt.Println("Retrieving items for category: ", walmartCategory.Name)
		RetrieveWalmartItemsByCategory(walmartCategory.Id)
	}
}
