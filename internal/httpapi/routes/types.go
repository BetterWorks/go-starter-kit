package routes

import "github.com/gin-gonic/gin"

type CreateController interface {
	Create(*gin.Context)
}

type DeleteController interface {
	Delete(*gin.Context)
}

type DetailController interface {
	Detail(*gin.Context)
}

type ListController interface {
	List(*gin.Context)
}

type UpdateController interface {
	Update(*gin.Context)
}

type Controller interface {
	CreateController
	DeleteController
	DetailController
	ListController
	UpdateController
}
