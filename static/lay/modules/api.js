//Solicitar URL
layui.define([], function (exports) {
    exports('api', {
        login: 'login',
        //Interface da barra de menu
        getMenu: "menu/index",
        addMenu: "menu/save",
        delMenu: "menu/del",
        listsMenu: "menu/lists",

        getClient: "client/index",
        addClient: "client/save",
        delClient: "client/del",
        listsClient: "client/lists",
    
        getGoods: 'json/goods.js',
        saveConfig: "sys/config",
        UpImg: "sys/upimg"
    });
})