package handler

import (
	"compress/gzip"
	"github.com/emicklei/go-restful/v3"
	"net/http"
	"xlog-apiserver-dscho/pkg/internal/repository"
)

type XLog struct {
	scouterClient   repository.Scouter
	timescaleClient repository.TimeScaleDB
	kubeClient      repository.Kubernetes
}

func (receiver XLog) GetXLog(req *restful.Request, resp *restful.Response) error {
	req.QueryParameter("a")
	req.QueryParameter("b")
	req.QueryParameter("c")
	return nil
}

func (receiver XLog) GetXLogList(req *restful.Request, resp *restful.Response) error {
	//limit := req.QueryParameter("limit")
	req.QueryParameter("b")
	req.QueryParameter("c")

	rw := resp.ResponseWriter
	w := gzip.NewWriter(rw)
	defer w.Close()

	defaultOffset := 5000
	limit := 20000
	count := 0
	for {
		// TODO limit이 offset보다 작은 경우와 큰 경우 체크
		all := receiver.timescaleClient.FindAll(defaultOffset, limit)
		count += len(all)
		if _, err := w.Write([]byte(all.ConvertedToString())); err != nil {
			return err
		}

		if f, ok := rw.(http.Flusher); ok {
			f.Flush()
		}

		// TODO 종료 개선
		if count < limit {
			return nil
		}
	}
}

func (receiver XLog) GetProfiles(req *restful.Request, resp *restful.Response) error {
	req.QueryParameter("a")
	req.QueryParameter("b")
	req.QueryParameter("c")

	return resp.WriteAsJson(nil)
}

func (receiver XLog) GetActiveService(req *restful.Request, resp *restful.Response) error {
	req.QueryParameter("a")
	req.QueryParameter("b")
	req.QueryParameter("c")

	return resp.WriteAsJson(nil)
}
