package handler

import (
	"github.com/autorunners/meerkat/core/request"
)

// 如果没有相关配置，则使用global中的相关配置
func mergeGlobal(req request.Request, gReq request.Request) request.Request {
	if req.Host == "" {
		req.Host = gReq.Host
	}
	if req.Timeout == 0 {
		req.Timeout = gReq.Timeout
	}
	for hk, hv := range gReq.Headers {
		if req.Headers[hk] == "" {
			req.Headers[hk] = hv
		}
	}
	if req.FullUri == "" {
		req.FullUri = req.Host + req.Uri
	}
	if req.Timeout == 0 {
		req.Timeout = 50 // ms
	}
	for ck, cv := range gReq.Cookies {
		if req.Cookies[ck] == "" {
			req.Cookies[ck] = cv
		}
	}
	return req
}
