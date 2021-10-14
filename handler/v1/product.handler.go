package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_heroku/common/obj"
	"github.com/ydhnwb/golang_heroku/common/response"
	"github.com/ydhnwb/golang_heroku/dto"
	"github.com/ydhnwb/golang_heroku/service"
)

type ProductHandler interface {
	All(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	FindOneProductByID(ctx *gin.Context)
}

type productHandler struct {
	productService service.ProductService
	jwtService     service.JWTService
}

func NewProductHandler(productService service.ProductService, jwtService service.JWTService) ProductHandler {
	return &productHandler{
		productService: productService,
		jwtService:     jwtService,
	}
}


// @Summary Get All Product
// @Description Get All Product
// @ID Get All Product
// @Param Authorization header string true "Token"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/product [get]
func (c *productHandler) All(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	products, err := c.productService.All(userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", products)
	ctx.JSON(http.StatusOK, response)
}

// @Summary Create Product
// @Description Create Product
// @ID Create Product
// @Param Authorization header string true "Token"
// @Param body body dto.CreateProductRequest true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/product [post]
func (c *productHandler) CreateProduct(ctx *gin.Context) {
	var createProductReq dto.CreateProductRequest
	err := ctx.ShouldBind(&createProductReq)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	res, err := c.productService.CreateProduct(createProductReq, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)

}

// @Summary Find Product By ID
// @Description Find Product By ID
// @ID Find Product By ID
// @Param Authorization header string true "Token"
// @Param id query string false "name search by id"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/product/:id [get]
func (c *productHandler) FindOneProductByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.productService.FindOneProductByID(id)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

// @Summary Find Product By ID
// @Description Find Product By ID
// @ID Find Product By ID
// @Param Authorization header string true "Token"
// @Param id query string false "name search by id"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/product/:id [delete]
func (c *productHandler) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	err := c.productService.DeleteProduct(id, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := response.BuildResponse(true, "OK!", obj.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}

// @Summary Find Product By ID
// @Description Find Product By ID
// @ID Find Product By ID
// @Param Authorization header string true "Token"
// @Param id query string false "name search by id"
// @Param body body dto.UpdateProductRequest true "request body"
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /api/product/:id [put]
func (c *productHandler) UpdateProduct(ctx *gin.Context) {
	updateProductRequest := dto.UpdateProductRequest{}
	err := ctx.ShouldBind(&updateProductRequest)

	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)
	updateProductRequest.ID = id
	product, err := c.productService.UpdateProduct(updateProductRequest, userID)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := response.BuildResponse(true, "OK!", product)
	ctx.JSON(http.StatusOK, response)

}
