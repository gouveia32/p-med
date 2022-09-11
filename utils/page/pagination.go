package page

import (
	"fmt"
	"math"
	"net/url"
	"sort"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
)

// Pagination struct
type Pagination struct {
	//当前页
	CurrentPage int
	//最后一页
	LastPage int
	//数据总数
	Total int
	//每页数量
	ListRows int
	//SimNão有下一页
	HasMore bool
	//渲染标签
	BootStrapRenderLink string
	//url中参数
	Parameters url.Values
}

//每页数量
//config配置参数
//  page:当前页,
//  path:url路径,
//  query:url额外参数,
//  fragment:url锚点,
//  var_page:分页变量,
//  list_rows:每页数量
func (pagination *Pagination) Paginate(seter orm.QuerySeter, listRows int, parameters url.Values) orm.QuerySeter {

	pagination.Parameters = parameters

	//当前页
	var page int
	pageStr := pagination.Parameters.Get("page")
	if pageStr == "" {
		page = 1
	} else {
		page, _ = strconv.Atoi(pageStr)
	}

	if page < 1 {
		page = 1
	}

	pagination.CurrentPage = page
	pagination.ListRows = listRows
	total, err := seter.Count()
	if err != nil {
		pagination.Total = 1
	}

	pagination.Total = int(total)
	pagination.LastPage = int(math.Ceil(float64(pagination.Total / pagination.ListRows)))
	if (pagination.Total % pagination.ListRows) > 0 {
		pagination.LastPage += 1
	}

	pagination.HasMore = pagination.CurrentPage < pagination.LastPage

	//Coloque na última execução, você precisa atribuir um valor na frente
	pagination.BootStrapRenderLink = pagination.render()

	return seter.Limit(pagination.ListRows, (pagination.CurrentPage-1)*pagination.ListRows)
}

//renderizar paginação html
func (pagination *Pagination) render() string {
	if pagination.hasPages() {
		return fmt.Sprintf("<ul class=\"pagination pagination-sm no-margin pull-right\">%s %s %s</ul>",
			pagination.getPreviousButton(),
			pagination.getLinks(),
			pagination.getNextButton(),
		)
	} else {
		return ""
	}
}

//Botão página anterior
func (pagination *Pagination) getPreviousButton() string {
	text := "&laquo;"
	if pagination.CurrentPage <= 1 {
		return pagination.getDisabledTextWrapper(text)
	}
	url := pagination.url(pagination.CurrentPage - 1)
	return pagination.getPageLinkWrapper(url, text)
}

//botão de próxima página
func (pagination *Pagination) getNextButton() string {
	text := "&raquo;"
	if !pagination.HasMore {
		return pagination.getDisabledTextWrapper(text)
	}
	url := pagination.url(pagination.CurrentPage + 1)
	return pagination.getPageLinkWrapper(url, text)
}

//Botão de página
func (pagination *Pagination) getLinks() string {
	block := map[string]map[int]string{
		"first":  nil,
		"slider": nil,
		"last":   nil,
	}

	side := 3
	window := side * 2

	if pagination.LastPage < window+6 {
		block["first"] = pagination.getUrlRange(1, pagination.LastPage)
	} else if pagination.CurrentPage <= window {
		block["first"] = pagination.getUrlRange(1, window+2)
		block["last"] = pagination.getUrlRange(pagination.LastPage-1, pagination.LastPage)
	} else if pagination.CurrentPage > (pagination.LastPage - window) {
		block["first"] = pagination.getUrlRange(1, 2)
		block["last"] = pagination.getUrlRange(pagination.LastPage-(window+2), pagination.LastPage)
	} else {
		block["first"] = pagination.getUrlRange(1, 2)
		block["slider"] = pagination.getUrlRange(pagination.CurrentPage-side, pagination.CurrentPage+side)
		block["last"] = pagination.getUrlRange(pagination.LastPage-1, pagination.LastPage)
	}

	html := ""
	if len(block["first"]) > 0 {
		html += pagination.getUrlLinks(block["first"])
	}

	if len(block["slider"]) > 0 {
		html += pagination.getDots()
		html += pagination.getUrlLinks(block["slider"])
	}

	if len(block["last"]) > 0 {
		html += pagination.getDots()
		html += pagination.getUrlLinks(block["last"])
	}

	return html

}

//Criar um conjunto de links de página
func (pagination *Pagination) getUrlRange(start, end int) map[int]string {
	urls := map[int]string{}
	for page := start; page <= end; page++ {
		urls[page] = pagination.url(page)
	}
	return urls
}

//Obtenha o link correspondente ao número da página
func (pagination *Pagination) url(page int) string {
	parameters := pagination.Parameters

	urlValue := parameters.Get("queryParamUrl")
	parameters.Del("queryParamUrl")
	parameters.Del("_pjax")

	if len(parameters) > 0 {
		//valor de cópia
		parameters.Set("page", strconv.Itoa(page))
		urlStr := parameters.Encode()
		return urlValue + "?" + urlStr
	}
	return urlValue + "?page=" + strconv.Itoa(page)
}

//Gerar um botão clicável
func (pagination *Pagination) getAvailablePageWrapper(url string, page string) string {
	return `<li><a href="` + url + `">` + page + `</a></li>`
}

//Gerar um botão Desativado
func (pagination *Pagination) getDisabledTextWrapper(text string) string {
	return `<li class="disabled"><span>` + text + `</span></li>`
}

//Gerar um botão ativo
func (pagination *Pagination) getActivePageWrapper(text string) string {
	return `<li class="active"><span>` + text + `</span></li>`
}

//Gerar botão de reticências
func (pagination *Pagination) getDots() string {
	return pagination.getDisabledTextWrapper("...")
}

//Botão gerar lote de número de página
func (pagination *Pagination) getUrlLinks(urls map[int]string) string {
	html := ""
	var sortKeys []int
	for page, _ := range urls {
		sortKeys = append(sortKeys, page)
	}
	sort.Ints(sortKeys)
	for _, page := range sortKeys {
		html += pagination.getPageLinkWrapper(urls[page], page)
	}
	return html
}

//Botão Gerar número de página normal
func (pagination *Pagination) getPageLinkWrapper(url string, page interface{}) string {
	pageInt, ok := page.(int)
	if ok {
		if pagination.CurrentPage == pageInt {
			return pagination.getActivePageWrapper(strconv.Itoa(pageInt))
		}
		return pagination.getAvailablePageWrapper(url, strconv.Itoa(pageInt))
	}
	pageStr := page.(string)
	return pagination.getAvailablePageWrapper(url, pageStr)
}

//Paginação sim/não
func (pagination *Pagination) hasPages() bool {
	return !(1 == pagination.CurrentPage && !pagination.HasMore)
}
