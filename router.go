package main

import (
	"fmt"

	"bitbucket.org/icehousecorp/inuit_service_backend/logger"

	"github.com/buaazp/fasthttprouter"
	"github.com/tokoku_backend/endpoints"
	// "github.com/contactapp/config"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	// ctx.Response.Header.Set("Access-Control-Allow-Headers", config.Config.CORS.Headers)
	// ctx.Response.Header.Set("Access-Control-Allow-Methods", config.Config.CORS.Methods)
	// ctx.Response.Header.Set("Access-Control-Allow-Origin", config.Config.CORS.Origin)
	fmt.Fprintf(ctx, "TOKOKU INDEX PAGE!\n")
}

func CreateRouter() {
	router := fasthttprouter.New()

	router.GET("/", Index)

	router.GET("/retrieve/item", endpoints.RetrieveItems)

	router.GET("/category", endpoints.GetListOfCategory)
	router.GET("/item", endpoints.GetListOfItemByCategory)
	logger.Println(fasthttp.ListenAndServe(":8080", router.Handler))
}
