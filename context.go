package ancontext

/*
*一个基础是context，提供代理功能
 */
import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

type anCtx struct {
	client *http.Client
}

//NewContext 创建一个新的上下文
func NewContext(transeport ...string) (context.Context, error) {
	ctx := &anCtx{}
	if len(transeport) != 0 && len(transeport[0]) != 0 {
		//设置代理
		transport, err := url.Parse(transeport[0])
		if err == nil {
			ctx.client = &http.Client{
				Transport: &http.Transport{
					Proxy: http.ProxyURL(transport),
				},
			}
		} else {
			log.Println(fmt.Errorf("error on url.Parse err:%w", err))
		}
	}
	return ctx, nil
}

func (*anCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*anCtx) Done() <-chan struct{} {
	return nil
}

func (*anCtx) Err() error {
	return nil
}

func (ctx *anCtx) Value(key interface{}) interface{} {

	// switch key.(type) {
	// case internal.HTTPClient:
	if ctx.client == nil {
		return http.DefaultClient
	}
	return ctx.client
	// default:
	// 	return nil
	// }
	// return nil
}

//String
func (*anCtx) String() string {
	return "an  custom context"
}
