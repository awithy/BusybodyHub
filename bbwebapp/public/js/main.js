angular.module('app', [])

//app.config(['$httpProvider', function($httpProvider) {
//
//});

/*
module.config(['$httpProvider', function($httpProvider) {
    $httpProvider.interceptors.push(['$q', 'mvIdentity', function ($q, mvIdentity) {
        return {
            'request': function(config) {
                config.headers = config.headers || {};
                if(mvIdentity.currentUser) {
                    config.headers.Authorization = "Bearer " + mvIdentity.currentUser.token;
                }
                return config;
            }
            'responseError': function(response) {
                if(response.status === 401 || response.status === 403) {
                    
                }
                return $q.reject(response);
            }
        });
    }
});
*/
