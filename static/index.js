layui.extend({admin: 'lay/modules/admin'}).define(['admin', 'conf'], function (exports) {
    //IE8 não suporta console
    window.console = window.console || (function () {
        var c = {};
        c.log = c.warn = c.debug = c.info = c.error = c.time = c.dir = c.profile
            = c.clear = c.exception = c.trace = c.assert = function () {
        };
        return c;
    })();
    layui.admin.initPage();
    exports('index', {});
});