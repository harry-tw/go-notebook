package article

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go-notebook/demo-go-kit/errors"
	"go-notebook/demo-go-kit/service"
)

type Request struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type Response struct {
	Message string `json:"message,omitempty"`
}

func MakeCreateEndpoint(svc service.ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*Request)
		if !ok {
			return nil, errors.EndpointReqTypeErr
		}
		err = svc.Create(ctx, req.Title, req.Content)
		if err != nil {
			return nil, err
		}
		rsp := Response{Message: "success"}
		return rsp, nil
	}
}
