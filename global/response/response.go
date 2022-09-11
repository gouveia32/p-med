package response

import (
	"net/http"
	"p-med/global"

	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/server/web/context"
)

const (
	ERROR   = 0
	SUCCESS = 1
)

// Response Estrutura do parâmetro de resposta
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Url  string      `json:"url"`
	Wait int         `json:"wait"`
}

// Result Retornar função auxiliar de resultado
func Result(code int, msg string, data interface{}, url string, wait int, header map[string]string, ctx *context.Context) {
	if ctx.Input.IsPost() {
		result := Response{
			Code: code,
			Msg:  msg,
			Data: data,
			Url:  url,
			Wait: wait,
		}

		if len(header) > 0 {
			for k, v := range header {
				ctx.Output.Header(k, v)
			}
		}

		ctx.Output.JSON(result, false, false)

		//Controller Uso de this.StopRun() em
		panic(beego.ErrAbort)
	}

	if url == "" {
		url = ctx.Request.Referer()
		if url == "" {
			url = "/admin/index/index"
		}
	}

	ctx.Redirect(http.StatusFound, url)
}

// Success Sucesso, Retornar Normal
func Success(ctx *context.Context) {
	Result(SUCCESS, "Operação realizada com sucesso.", "", global.URL_BACK, 0, map[string]string{}, ctx)
}

// SuccessWithMessage Sucesso, Retornar mensagem personalizada
func SuccessWithMessage(msg string, ctx *context.Context) {
	Result(SUCCESS, msg, "", global.URL_BACK, 0, map[string]string{}, ctx)
}

// SuccessWithMessageAndUrl Sucesso, Retornar informações personalizadas e url
func SuccessWithMessageAndUrl(msg string, url string, ctx *context.Context) {
	Result(SUCCESS, msg, "", url, 0, map[string]string{}, ctx)
}

// SuccessWithDetailedSucesso, Retornar todas as informações personalizadas
func SuccessWithDetailed(msg string, url string, data interface{}, wait int, header map[string]string, ctx *context.Context) {
	Result(SUCCESS, msg, data, url, wait, header, ctx)
}

// Error Falha, Retornar Normal
func Error(ctx *context.Context) {
	Result(ERROR, "Oeração Falhou", "", global.URL_CURRENT, 0, map[string]string{}, ctx)
}

// ErrorWithMessage Falha, Retornar mensagem personalizada
func ErrorWithMessage(msg string, ctx *context.Context) {
	Result(ERROR, msg, "", global.URL_CURRENT, 0, map[string]string{}, ctx)
}

// ErrorWithMessageAndUrl Falha, Retornar informações personalizadas e url
func ErrorWithMessageAndUrl(msg string, url string, ctx *context.Context) {
	Result(ERROR, msg, "", url, 0, map[string]string{}, ctx)
}

// ErrorWithDetailed Falha, Retornar todas as informações personalizadas
func ErrorWithDetailed(msg string, url string, data interface{}, wait int, header map[string]string, ctx *context.Context) {
	Result(ERROR, msg, data, url, wait, header, ctx)
}
