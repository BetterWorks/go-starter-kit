package routes

import "github.com/gin-gonic/gin"

// CreateController
type CreateController interface {
	Create(*gin.Context)
}

// DeleteController
type DeleteController interface {
	Delete(*gin.Context)
}

// DetailController
type DetailController interface {
	Detail(*gin.Context)
}

// ListController
type ListController interface {
	List(*gin.Context)
}

// UpdateController
type UpdateController interface {
	Update(*gin.Context)
}

// Controller
type Controller interface {
	CreateController
	DeleteController
	DetailController
	ListController
	UpdateController
}
