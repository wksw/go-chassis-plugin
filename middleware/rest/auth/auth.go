/*
	rest认证鉴权
*/
package rest_auth_middleware

import (
	"github.com/go-chassis/go-chassis/v2/core/handler"
	"github.com/go-chassis/go-chassis/v2/core/invocation"
)

const (
	HANDLER_NAME  = "auth_rest"
	PROTOCOL_NAME = "rest"
)

type Handler struct{}

func init() {
	handler.RegisterHandler(HANDLER_NAME, New)
}

func New() handler.Handler {
	return &Handler{}
}

func (m *Handler) Name() string {
	return HANDLER_NAME
}

// rest请求认证鉴权
func (m *Handler) Handle(chain *handler.Chain, inv *invocation.Invocation, cb invocation.ResponseCallBack) {
	if inv.Protocol == PROTOCOL_NAME {
		chain.Next(inv, cb)
		return
	} else {
		chain.Next(inv, cb)
	}
}
