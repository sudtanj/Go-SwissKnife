package rest_helper

import (
	"github.com/gin-gonic/gin"
	"github.com/sudtanj/Go-SwissKnife/responder"
	"net/http"
)

type IPostHelper[PARAMS any, QUERY any, BODY any] interface {
}

type PostHelper[PARAMS any, QUERY any, BODY any] struct {
	responder responder.IJSONResponder
}

func (p *PostHelper[PARAMS, QUERY, BODY]) GetParamsQueryBody(c *gin.Context) (params *PARAMS, query *QUERY, body *BODY, err error) {
	if err = c.BindJSON(body); err != nil {
		p.responder.JSON(http.StatusBadRequest, err)
		return nil, nil, nil, err
	}

	if err = c.BindQuery(query); err != nil {
		p.responder.JSON(http.StatusBadRequest, err)
		return nil, nil, nil, err
	}

	return params, query, body, nil
}
