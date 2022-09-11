layui.define(function (exports) {
    exports('conf', {
        layTplConfig: {
            open: '{{',
            close: '}}'
        },
        //容器ID
        container: 'app',
        //容器内容ID
        containerBody: 'app-body',
        //número da versão
        v: layui.cache.version,
        //记录nepadmin文件夹所在路径
        base: layui.cache.base,
        css: layui.cache.base + 'css/',
        //O diretório onde a vista está localizada
        // views: layui.cache.base + 'views/',
        views: '',
        //Se deve abrir a guia
        viewTabs: true,
        //Mostrar barra de carregamento da página
        viewLoadBar: true,
        //Estilos comumente carregados
        style: [
            //layui.cache.base + "css/admin.css"
        ],
        //Para habilitar o modo de depuração, se a exceção da interface estiver habilitada, uma exceção será lançada nas informações de URL
        debug: layui.cache.debug,
        //Nome do site
        name: 'CCB',
        //Nome do arquivo de visualização padrão
        entry: '/main',
        //Ver extensão de arquivo
        engine: '.html',
        eventName: 'nepadmin-event',
        //本地存储表名
        tableName: 'nepadmin',
        //全局设置 headers 信息
        requestHeaders: {
            'Test-User-Agent': 'os=pc;ver=0.0.1;imei=asdasdas'
        },
        //request URL base
        requestUrl: './',
        //独立页面路由，可随意添加（无需写参数）
        indPage: [
            '/login', //登入页
            '/reg', //注册页
            '/user/forget' //找回密码
        ],
        //登录页面，当未登录或登录失效时进入
        loginPage: '/login',
        //登录 token 名称，request 请求的时候会带上此参数到 header
        tokenName: 'token',
        //是否要强制检查登录状态， 使用tokenName进行登录验证，不通过的话会返回 loginPage 页面
        loginCheck: true,
        //根据服务器返回的 HTTP 状态码检查登录过期，设置为false不通过http返回码检查
        logoutHttpCode: '401',
        //全局自定义响应字段
        response: {
            //数据状态的字段名称
            statusName: 'code',
            statusCode: {
                //数据状态一切正常的状态码
                ok: 200,
                //数据状态一切错误的状态码
                error: 500,
                //通过接口返回的登录过期状态码
                logout: 401
            },
            msgName: 'msg', //状态信息的字段名称
            dataName: 'data', //数据详情的字段名称
            countName: 'count' //数据条数的字段名称，用于 table
        },
        //全局 table 配置
        //参数请参照 https://www.layui.com/doc/modules/table.html
        table: {
            page: true,
            size: 'lg',
            skin: 'line',
            //每页显示的条数
            limit: 20,
            //是否显示加载条
            loading: true,
            //用于对分页请求的参数：page、limit重新设定名称
            request: {
                pageName: 'page', //O nome do parâmetro do número da página, padrão: página
                limitName: 'size' //O nome do parâmetro do volume de dados por página, padrão: limite
            }
        },
        //Extensão de terceiros
        extend: {
            //Método de expansão nos bastidores de acordo com as necessidades do negócio
            helper: 'lay/extends/helper',
            //Gerar código QR
            qrcode: 'lay/extends/qrcode',
            // Gerar criptografia MD5
            md5: 'lay/extends/md5',
            //Gerar gráfico
            echarts: 'lay/extends/echarts',
            echartsTheme: 'lay/extends/echartsTheme',
            //Copiar conteúdo para a área de transferência
            clipboard: 'lay/extends/clipboard',
            //Funções do menu árvore
            treeGrid: 'lay/extends/treeGrid'
        }
    })
})
