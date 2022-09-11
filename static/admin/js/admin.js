try {
    /** pjax相关 */
    //超时3秒(可选)
    $.pjax.defaults.timeout = 3000;
    $.pjax.defaults.type = 'GET';
    //存储容器id
    $.pjax.defaults.container = '#pjax-container';
    //目标id
    $.pjax.defaults.fragment = '#pjax-container';
    //最大缓存长度(可选)
    $.pjax.defaults.maxCacheLength = 0;
    $(document).pjax('a:not(a[target="_blank"])', {
        //存储容器id
        container: '#pjax-container',
        //目标id
        fragment: '#pjax-container'
    });
    //ajax请求开始时执行
    $(document).ajaxStart(function () {
        //启动进度条
        NProgress.start();
    }).ajaxStop(function () {
        //ajax请求结束后执行
        //关闭进度条
        NProgress.done();
    });
} catch (e) {
    if (adminDebug) {
        console.log(e.message);
    }
}

$(document).on('pjax:timeout', function (event) {
    event.preventDefault();
});
$(document).on('pjax:send', function (xhr) {
    NProgress.start();
});
$(document).on('pjax:complete', function (xhr) {
    $('[data-toggle="tooltip"]').tooltip();
    NProgress.done();
});
//列表页搜索pjax
$(document).on('submit', '.searchForm', function (event) {
    $.pjax.submit(event, '#pjax-container');
});


//菜单搜索
$(function () {
    $('#sidebar-form').on('submit', function (e) {
        e.preventDefault();
    });

    $('.sidebar-menu li.active').data('lte.pushmenu.active', true);

    $('#search-input').on('keyup', function () {
        var term = $('#search-input').val().trim();

        if (term.length === 0) {
            $('.sidebar-menu li').each(function () {
                $(this).show(0);
                $(this).removeClass('active');
                if ($(this).data('lte.pushmenu.active')) {
                    $(this).addClass('active');
                }
            });
            return;
        }

        $('.sidebar-menu li').each(function () {
            if ($(this).text().toLowerCase().indexOf(term.toLowerCase()) === -1) {
                $(this).hide(0);
                $(this).removeClass('pushmenu-search-found', false);

                if ($(this).is('.treeview')) {
                    $(this).removeClass('active');
                }
            } else {
                $(this).show(0);
                $(this).addClass('pushmenu-search-found');

                if ($(this).is('.treeview')) {
                    $(this).addClass('active');
                }

                var parent = $(this).parents('li').first();
                if (parent.is('.treeview')) {
                    parent.show(0);
                }
            }

            if ($(this).is('.header')) {
                $(this).show();
            }
        });

        $('.sidebar-menu li.pushmenu-search-found.treeview').each(function () {
            $(this).find('.pushmenu-search-found').show(0);
        });
    });
});

//点击菜单高亮
$(function () {
    $('.sidebar-menu li:not(.treeview) > a').on('click', function () {
        var $parent = $(this).parent().addClass('active');
        $parent.siblings('.treeview.active').find('> a').trigger('click');
        $parent.siblings().removeClass('active').find('li').removeClass('active');
    });

    $('[data-toggle="popover"]').popover();
});

//bootstrap提示
$(function () {
    $('[data-toggle="tooltip"]').tooltip();
});

/** 表单验证错误展示 */
$.validator.setDefaults({
    errorElement: "span",
    errorClass: "help-block error",
    submitHandler: function (form) {
        formSubmit(form);
        return false;
    }
});

/** 边栏菜单 */
try {
    $('.sidebar-menu').tree();
} catch (e) {
    if (adminDebug) {
        console.log(e.message);
    }
}


/* 清除搜索表单 */
function clearSearchForm() {
    let url_all = window.location.href;
    let arr = url_all.split('?');
    let url = arr[0];
    $.pjax({url: url, container: '#pjax-container'});
}


$(function () {
    /* Retornar按钮 */
    $('body').on('click', '.BackButton', function (event) {
        event.preventDefault();
        history.back(1);
    });

    /* Atualizar按钮 */
    $('body').on('click', '.ReloadButton', function (event) {
        event.preventDefault();
        $.pjax.reload();
    });

});


/*列表中单个选择和取消*/
function checkThis(obj) {
    var id = $(obj).attr('value');
    if ($(obj).is(':checked')) {
        if ($.inArray(id, dataSelectIds) < 0) {
            dataSelectIds.push(id);
        }
    } else {
        if ($.inArray(id, dataSelectIds) > -1) {
            dataSelectIds.splice($.inArray(id, dataSelectIds), 1);
        }
    }

    var all_length = $("input[name='data-checkbox']").length;
    var checked_length = $("input[name='data-checkbox']:checked").length;
    if (all_length === checked_length) {
        $("#dataCheckAll").prop("checked", true);
    } else {
        $("#dataCheckAll").prop("checked", false);
    }
    console.log(dataSelectIds);
}

/*全部选择/取消*/
function checkAll(obj) {
    dataSelectIds = [];
    var all_check = $("input[name='data-checkbox']");
    if ($(obj).is(':checked')) {
        all_check.prop("checked", true);
        $(all_check).each(function () {
            dataSelectIds.push(this.value);
        });
    } else {
        all_check.prop("checked", false);
    }
}


/* 表单提交 */
function formSubmit(form) {
    let loadT = layer.msg('Enviando, aguarde...', {icon: 16, time: 0, shade: [0.3, "#000"],scrollbar: false,});
    let action = $(form).attr('action');
    let method = $(form).attr('method');
    let data = new FormData($(form)[0]);

    //设置全局通用_xsrf验证token
    data.set("_xsrf",$('meta[name="_xsrf"]').attr('content'));

    $.ajax({
            url: action,
            dataType: 'json',
            type: method,
            data: data,
            contentType: false,
            processData: false,
            success: function (result) {
                layer.close(loadT);
                layer.msg(result.msg, {
                    icon: result.code ? 1 : 2,
                    scrollbar: false,
                });
                goUrl(result.url);
            },
            error: function (xhr, type, errorThrown) {
                //异常处理；
                layer.msg('falha no acesso, código: ' + xhr.status, {icon: 2,scrollbar: false,});
            }
        }
    );
    return false;
}


/** 跳转到指定url */
function goUrl(url = 1) {
    console.log(url);
    //清除列表页选择的ID
    if (url !== 'url://current' && url !== 1) {
        dataSelectIds = [];
    }
    if (url === 'url://current' || url === 1) {
        console.log('Stay current page.');
    } else if (url === 'url://reload' || url === 2) {
        console.log('Reload current page.');
        $.pjax.reload();
    } else if (url === 'url://back' || url === 3) {
        console.log('Return to the last page.');
        history.back(1);
    }else if (url === 4 || url === 'url://close-refresh') {
        console.log('Close this layer page and refresh parent page.');
        let indexWindow = parent.layer.getFrameIndex(window.name);
        //先atualizar父级页面
        parent.goUrl(2);
        //再关闭当前layer弹窗
        parent.layer.close(indexWindow);
    } else if (url === 5 || url === 'url://close-layer') {
        console.log('Close this layer page.');
        let indexWindow = parent.layer.getFrameIndex(window.name);
        parent.layer.close(indexWindow);
    } else {
        console.log('Go to ' + url);
        try {
            $.pjax({
                url: url,
                container: '#pjax-container'
            });
        } catch (e) {
            window.location.href = url;
        }
    }
}

/**
 * ajax访问按钮
 * 例如元素为<a class="AjaxButton" 
 * data-confirm="1" 
 * data-type="1" 
 * data-url="disable" 
 * data-id="2" 
 * data-go="" ></a>
 * data-confirm  por SimNão prompt pop-up，1为Sim，2为Não。比如删除某条数据，data-confirm="1"就会弹出来提示
 * data-type为访问方式，1为直接ajax访问，例如删除操作。2Sim为打开layer窗口展示数据，例如查看操作日志详情
 * data-url为要访问的url
 * data-id为要操作的数据ID，可以填写正常的数据ID，例如
 * data-id="2"，
 * 或者填写checked表示获取当前数据列表选择的ID，也就Sim取的变量dataSelectIds的值
 * data-go为操作完成后的跳转url，不设置此参数默认根据后台Retornar的url跳转
 * data-confirm-title为确认提示弹窗的标题 例如data-confirm-title="删除警告"
 * data-confirm-content: O conteúdo do prompt de confirmação, por exemplo
 * data-confirm-content="Tem certeza de que deseja excluir esses dados?"
 * data-title: título da janela
 *
 */
$(function () {
    $('body').on('click', '.AjaxButton', function (event) {
        event.preventDefault();

        if (adminDebug) {
            console.log('AjaxButton clicked.');
        }
        //Sim Não menu pop-up
        var layerConfirm = $(this).data("confirm") || 1;
        //Modo de acesso, 1 é acesso direto, 2 é exibição de janela de camada
        var layerType = parseInt($(this).data("type") || 1);
        // URL visitada
        var url = $(this).data("url");
        //Método de acesso, post padrão
        var layerMethod = $(this).data("method") || 'post';
        //A página a ser saltada após o acesso bem sucedido, se este parâmetro não estiver definido, o padrão é saltar de acordo com a url do background Retornar
        var go = $(this).data("go") || 'url://reload';

        //Largura e altura podem ser definidas quando exibidas para a janela
        var layerWith = $(this).data("width") || '80%';
        var layerHeight = $(this).data("height") || '60%';

        //título da janela
        var layerTitle = $(this).data('title');

        //ID dos dados da operação atual
        var dataId = $(this).data("id");

        //Se nenhum ID for definido, consulte a propriedade data-data
        if (dataId === undefined) {
            var dataData = $(this).data("data") || {};
        } else {
            if (dataId === 'checked') {
                if (dataSelectIds.length === 0) {
                    layer.msg('Por favor, selecione pelo menos um registro', {icon: 2,scrollbar: false,});
                    return false;
                }
                dataId = dataSelectIds;
            }

            dataData = {"id": dataId};
        }

        //ajax设置beego 的 xsrf token验证
        $.ajaxSetup({
            headers: {
                'X-XSRFToken': $('meta[name="_xsrf"]').attr('content')
            }
        });

        if (typeof (dataData) != 'object') {
            dataData = JSON.parse(dataData);
        }

        /*Ação requerida*/
        if (parseInt(layerConfirm) === 1) {
            //o título da janela de prompt
            var confirmTitle = $(this).data("confirmTitle") || 'Confirmação...';
            //Conteúdo da janela de prompt
            var confirmContent = $(this).data("confirmContent") || 'Certeza de que quer fazer isso?';
            layer.confirm(confirmContent, {title: confirmTitle, btn: ['Sim','Não'], icon: 3}, function () {
                //se acesso direto
                if (layerType === 1) {
                    ajaxRequest(url, layerMethod, dataData, go);
                } else if (layerType === 2) {
                    //Se a janela estiver aberta
                    //Consultar permissões primeiro
                    if (checkAuth(url)) {
                        layer.open({
                            type: 1,
                            area: [layerWith, layerHeight],
                            title: layerTitle,
                            closeBtn: 1,
                            shift: 0,
                            content: url + "?request_type=layer_open&" + parseParam(dataData),
                            scrollbar: false,
                        });
                    }
                }
            });
        } else {
            //A confirmação da operação não é necessária
            if (layerType === 1) {
                //pedido direto
                ajaxRequest(url, layerMethod, dataData, go);
            } else if (layerType === 2) {
                //janela de pop-up
                //Verificar permissões
                if (checkAuth(url)) {
                    //aberto com janela
                    layer.open({
                        type: 2,
                        area: [layerWith, layerHeight],
                        title: layerTitle,
                        closeBtn: 1,
                        shift: 0,
                        content: url + "?request_type=layer_open&" + parseParam(dataData),
                        scrollbar: false,
                    });
                }
            }
        }
    });
});

//ajax请求封装
/**
 *
 * @param url 访问的url
 * @param method  访问方式
 * @param data  data数据
 * @param go 要跳转的url
 */
function ajaxRequest(url, method, data, go) {
    var loadT = layer.msg('Enviando, aguarde...', {icon: 16, time: 0, shade: [0.3, '#000'],scrollbar: false,});
    $.ajax({
            url: url,
            dataType: 'json',
            type: method,
            data: data,
            success: function (result) {
                layer.close(loadT);
                layer.msg(result.msg, {
                    icon: result.code ? 1 : 2,
                    scrollbar: false,
                });

                goUrl(go);
            },
            error: function (xhr, type, errorThrown) {
                //异常处理；
                layer.msg('falha no acesso, código: ' + xhr.status, {icon: 2,scrollbar: false,});
            }
        }
    );
}

//改变每页数量
function changePerPage(obj) {
    Cookies.set(cookiePrefix + 'admin_per_page', obj.value, {expires:30});
    $.pjax.reload();
}


/**
 * 检查授权
 */
function checkAuth(url) {
    var hasAuth = false;
    var loadT = layer.msg('Solicitando, aguarde...', {icon: 16, time: 0, shade: [0.3, '#000'],scrollbar: false,});
    $.post({
        url: url,
        data: {"check_auth": 1},
        dataType: 'json',
        async: false,
        success: function (result) {
            layer.close(loadT);
            if (result.code === 1) {
                hasAuth = true;
            } else {
                layer.msg(result.msg, {
                    icon: 2,
                    scrollbar: false,
                });
            }
        },
        error: function (xhr, type, errorThrown) {
            layer.msg('falha no acesso, código: ' + xhr.status, {icon: 2,scrollbar: false,});
        }
    });
    return hasAuth;
}

/** 处理url参数 **/
function parseParam(param, key) {
    var paramStr = "";
    if (param instanceof String || param instanceof Number || param instanceof Boolean) {
        paramStr += "&" + key + "=" + encodeURIComponent(param);
    } else {
        $.each(param, function (i) {
            var k = key == null ? i : key + (param instanceof Array ? "[" + i + "]" : "." + i);
            paramStr += '&' + parseParam(this, k);
        });
    }
    return paramStr.substr(1);
}

/** 导出excel **/
function exportData(url) {
    var exportUrl = url || 'index';
    var openUrl = exportUrl + '?export_data=1&' + $("#searchForm").serialize();
    window.open(openUrl);

}


