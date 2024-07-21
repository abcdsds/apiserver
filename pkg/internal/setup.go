package internal

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"xlog-apiserver-dscho/pkg/internal/handler"
)

const root = "/root"

type Server struct {
	logger  logr.Logger
	handler handler.XLog
}

func (s *Server) Run() {
	ws := new(restful.WebService)
	ws.Path(root).Produces(restful.MIME_JSON, "text/plain")

	ws.Route(
		ws.GET("/xlogs").
			Param(restful.PathParameter("a", "b")).
			Param(restful.QueryParameter("a", "b")).
			Param(restful.QueryParameter("a", "b")).
			Operation("GetXLogs").
			To(HandleError(s.handler.GetXLogList)))
	ws.Route(
		ws.GET("/xlogs/{txid}").
			Param(restful.PathParameter("a", "b")).
			Param(restful.QueryParameter("a", "b")).
			Param(restful.QueryParameter("a", "b")).
			Operation("GetXLogByTxid").
			To(HandleError(s.handler.GetXLog)))

	ws.Route(
		ws.GET("/activeServices").
			Param(restful.PathParameter("a", "b")).
			Param(restful.QueryParameter("a", "b")).
			Param(restful.QueryParameter("a", "b")).
			Operation("GetActiveServices").
			To(HandleError(s.handler.GetActiveService)))

	ws.Route(
		ws.GET("/profiles").
			Param(restful.PathParameter("a", "b")).
			Param(restful.QueryParameter("a", "b")).
			Param(restful.QueryParameter("a", "b")).
			Operation("GetProfiles").
			To(HandleError(s.handler.GetProfiles)))

	s.logger.Info("run")
	s.Run()
}

func HandleError(handler func(*restful.Request, *restful.Response) error) restful.RouteFunction {
	return func(request *restful.Request, response *restful.Response) {
		err := handler(request, response)
		if err != nil {
			status := apierrors.NewInternalError(err).Status()
			response.WriteHeaderAndJson(int(status.Code), status, restful.MIME_JSON)
		}
	}
}
