package handler

import (
	"hmdp/src/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigRouter(r *gin.Engine) {

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	blogController := r.Group("/blog", middleware.JWTAuth())

	{
		blogController.POST("", blogHandler.SaveBlog)
		blogController.PUT("/like/:id", blogHandler.LikeBlog)
		blogController.GET("/of/me", blogHandler.QueryMyBlog)
		blogController.GET("/:id", blogHandler.GetBlogById)
		blogController.GET("/likes/:id", blogHandler.QueryUserLiked)
		blogController.GET("/of/follow", blogHandler.QueryBlogOfFollow)
	}

	blogControllerWithOutMid := r.Group("/blog")
	{
		blogControllerWithOutMid.GET("/hot", blogHandler.QueryHotBlog)
	}

	// blogCommentsController := r.Group("/blog-comments")

	// followController := r.Group("/follow")

	shopController := r.Group("/shop", middleware.JWTAuth())

	{
		shopController.GET("/:id", shopHandler.QueryShopById)
		shopController.POST("", shopHandler.SaveShop)
		shopController.PUT("", shopHandler.UpdateShop)
		shopController.GET("/of/type", shopHandler.QueryShopByType)
		shopController.GET("/of/name", shopHandler.QueryShopByName)
	}

	shopTypeController := r.Group("/shop-type")

	{
		shopTypeController.GET("/list", shopTypeHandler.QueryShopTypeList)
	}

	uploadController := r.Group("/upload", middleware.JWTAuth())

	{
		uploadController.POST("/blog", uploadHandler.UploadImage)
		uploadController.GET("/blog/delete", uploadHandler.DeleteBlogImg)
	}

	userController := r.Group("/user", middleware.JWTAuth())

	{
		userController.POST("/logout", userHandler.Logout)
		userController.GET("/me", userHandler.Me)
		userController.GET("/info/:id", userHandler.Info)
	}

	userControllerWithOutMid := r.Group("/user")

	{
		userControllerWithOutMid.POST("/code", userHandler.SendCode)
		userControllerWithOutMid.POST("/login", userHandler.Login)
	}

	voucherController := r.Group("/voucher", middleware.JWTAuth())

	{
		voucherController.POST("", voucherHandler.AddVoucher)
		voucherController.POST("/seckill", voucherHandler.AddSecKillVoucher)
		voucherController.GET("/list/:shopId", voucherHandler.QueryVoucherOfShop)
	}

	voucherOrderController := r.Group("/voucher-order", middleware.JWTAuth())

	{
		voucherOrderController.POST("/seckill/:id", voucherOrderHandler.SeckillVoucher)
	}

	followContoller := r.Group("/follow", middleware.JWTAuth())

	{
		followContoller.PUT("/:id/:isFollow", followHanlder.Follow)
		followContoller.GET("/common/:id", followHanlder.FollowCommons)
		followContoller.GET("/or/not/:id", followHanlder.IsFollow)
	}
}
