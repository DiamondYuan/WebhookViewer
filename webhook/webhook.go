package webhook

import (
	"github.com/go-macaron/macaron"
)

type request struct{
	body string
	header map[string][]string
}
var cacheMap map[string]request

func GetWebhook(ctx *macaron.Context)  {

	if _, ok := cacheMap[ctx.Req.URL.Path]; ok {
		ctx.Resp.Write([]byte(cacheMap[ctx.Req.URL.Path].body))
	}else {
		ctx.Resp.Write([]byte("not found"))
	}
}

func SetWebhook(ctx *macaron.Context)  {
	requestBody,err:=ctx.Req.Body().String()
	if err != nil {
		ctx.Resp.Write([]byte("failed"))
		return
	}
	if cacheMap == nil {
		cacheMap = make(map[string]request)
	}
	cacheMap[ctx.Req.URL.Path] = request{
		requestBody,
		ctx.Req.Header,
	}
	ctx.Resp.Write([]byte("success"))
}





