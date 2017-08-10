package endpoints

import (
	"encoding/json"

	"github.com/tokoku_backend/repository"
	"github.com/valyala/fasthttp"
)

type CategoryResponseJSON struct {
	Categories []repository.WalmartCategory `json:"categories"`
}

func GetListOfCategory(ctx *fasthttp.RequestCtx) {

	listOfCategory, err := repository.Mgr.GetListOfWalmartCategory()
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
		return
	}

	var categoryResponseJSON = CategoryResponseJSON{
		Categories: listOfCategory,
	}

	ctx.Response.Header.Set("Content-Type", "application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	json.NewEncoder(ctx).Encode(categoryResponseJSON)
}
