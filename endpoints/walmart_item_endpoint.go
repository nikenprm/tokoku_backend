package endpoints

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/tokoku_backend/repository"
	"github.com/tokoku_backend/services"
	"github.com/valyala/fasthttp"
)

type ItemResponseJSON struct {
	Meta ItemMetaResponse `json:"meta"`
	Data ItemDataResponse `json:"data"`
}

type ItemMetaResponse struct {
	RecordsCount int `json:"records_count"`
}

type ItemDataResponse struct {
	Items []repository.WalmartItem `json:"items"`
}

type JSONResponse struct {
	Message string `json:"message"`
}

func GetListOfItemByCategory(ctx *fasthttp.RequestCtx) {
	categoryID := ctx.QueryArgs().Peek("categoryID")
	page := ctx.QueryArgs().Peek("page")
	perPage := ctx.QueryArgs().Peek("perpage")

	pageInt := 0
	perPageInt := 0
	var errPage error
	if string(categoryID) == "" {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if string(page) == "" {
		pageInt = 1
	} else {
		pageString := string(page)
		pageInt, errPage = strconv.Atoi(pageString)
		if errPage != nil {
			fmt.Println(errPage)
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}
	}

	if string(perPage) == "" {
		perPageInt = 10
	} else {
		perPageString := string(perPage)
		perPageInt, errPage = strconv.Atoi(perPageString)
		if errPage != nil {
			fmt.Println(errPage)
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}
	}

	fmt.Println(string(categoryID))
	fmt.Println("perPage:", string(perPage))
	fmt.Println("perPageInt: ", perPageInt)

	listOfItem, err := repository.Mgr.GetListOfWalmartItemByCategory(string(categoryID), pageInt, perPageInt)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		return
	}

	itemCount := len(listOfItem)
	var itemMetaResponse = ItemMetaResponse{
		RecordsCount: itemCount,
	}

	var itemDataResponse = ItemDataResponse{
		Items: listOfItem,
	}

	var itemResponseJSON = ItemResponseJSON{
		Meta: itemMetaResponse,
		Data: itemDataResponse,
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	json.NewEncoder(ctx).Encode(itemResponseJSON)
}

func RetrieveItems(ctx *fasthttp.RequestCtx) {

	services.RetrieveAllWalmartItems()

	var jsonResponse = JSONResponse{
		Message: "Successfully retrieve all walmart items.",
	}
	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	json.NewEncoder(ctx).Encode(jsonResponse)

}
