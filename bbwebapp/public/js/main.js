angular.module('app', ['ngRoute'])

angular.module('app').config(['$httpProvider', function($httpProvider) {
    $httpProvider.interceptors.push(['$q', 'mvIdentity', function ($q, mvIdentity) {
        return {
            'request': function(config) {
                config.headers = config.headers || {};
                if(mvIdentity.currentUser) {
                   config.headers.Authorization = 'Bearer ' + mvIdentity.currentUser.token;
                }
                return config;
            },
            'responseError': function(response) {
                return $q.reject(response);
            }
        };
    }]);
}]);

angular.module('app').config(function($routeProvider) {
    $routeProvider.when('/sites', {templateUrl: 'app/sites/sites.html', controller: 'mvSitesCtrl'});
});
