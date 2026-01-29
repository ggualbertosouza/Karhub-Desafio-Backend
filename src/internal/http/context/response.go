package HttpContext

import "github.com/gin-gonic/gin"

func ResponseOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, data)
}

func ResourceCreated(ctx *gin.Context, resource string) {
	ctx.JSON(201, resource+" created successfully.")
}

func ResourceUpdated(ctx *gin.Context, resource string) {
	ctx.JSON(200, resource+" updated successfully.")
}

func ResourceDeleted(ctx *gin.Context, resource string) {
	ctx.JSON(200, resource+" deleted successfully.")
}
