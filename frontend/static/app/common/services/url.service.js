(function () {
  app.factory('URL', URL);
  URL.$inject = ['$window'];
  function URL($window) {
    var result = {
      'login': buildUrl('login'),
      'bots': buildUrl('bots'),
      'consts': buildUrl('consts'),
      'loginwechat': buildUrl('loginwechat'),
      'loginqq': buildUrl('loginqq')
    };
    return result;
    
    function buildUrl(url) {
      var host = '';
      var pathprefix = $window.location.pathname;

      host = $window.location.protocol + "//" + $window.location.host;
      /*
      if($window.location.port != "80" && $window.location.port != "443" && $window.location.port != "") {
	host += ":" + $window.location.port;
      }
      */
      //console.log(host +"," + pathprefix + "," + url);
      return host + pathprefix + url;
    }
  }
})();
