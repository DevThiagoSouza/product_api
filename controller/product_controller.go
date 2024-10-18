package controller

import (
	"game-shop-api/model"
	"game-shop-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) ProductController {
	return ProductController{
		productUsecase: usecase,
	}
}

func (pc *ProductController) GetProducts(ctx *gin.Context) {

	products, err := pc.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)

}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var product model.Procuct

	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertProduct, err := pc.productUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, insertProduct)
}

func (pc *ProductController) GetProductsById(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := model.Response{
			Messagem: "ID nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	idProduct, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Messagem: "ID tem que ser um numero",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := pc.productUsecase.GetProductsById(idProduct)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Messagem: "Nemhum dado encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (pc *ProductController) DeleteProducts(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := model.Response{
			Messagem: "ID nao pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	idProduct, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Messagem: "ID tem que ser um numero",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := pc.productUsecase.GetProductsById(idProduct)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Messagem: "Nemhum dado encontrado na base de dados",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
