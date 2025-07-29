package httpserve

import (
	services "github.com/build-smile/CRUD_APIs_Arise/internal/core/services/member"
	handlers "github.com/build-smile/CRUD_APIs_Arise/internal/handlers/member"

	"github.com/build-smile/CRUD_APIs_Arise/internal/repositories"
	"github.com/gin-gonic/gin"
)

func BindGetMembers(r *gin.Engine) gin.IRoutes {
	repo := repositories.NewMemberRepo()
	svc := services.NewGetMembersSvc(repo)
	hdl := handlers.NewGetMembersHdl(svc)
	return r.GET("/members", hdl.Handle)
}

func BindGetMember(r *gin.Engine) gin.IRoutes {
	repo := repositories.NewMemberRepo()
	svc := services.NewGetMemberSvc(repo)
	hdl := handlers.NewGetMemberHdl(svc)
	return r.GET("/member/:id", hdl.Handle)
}

func BindCreateMember(r *gin.Engine) gin.IRoutes {
	repo := repositories.NewMemberRepo()
	svc := services.NewCreateMemberSvc(repo)
	hdl := handlers.NewCreateMemberHdl(svc)
	return r.POST("/member", hdl.Handle)
}

func BindUpdateMember(r *gin.Engine) gin.IRoutes {
	repo := repositories.NewMemberRepo()
	svc := services.NewupdateMemberSvc(repo)
	hdl := handlers.NewUpdateMemberHdl(svc)
	return r.PATCH("/member/:id", hdl.Handle)

}
func BindDeleteMember(r *gin.Engine) gin.IRoutes {
	repo := repositories.NewMemberRepo()
	svc := services.NewDeleteMemberSvc(repo)
	hdl := handlers.NewDeleteMemberHdl(svc)
	return r.DELETE("/member/:id", hdl.Handle)
}
