angular.module('app').factory('mvAuth', function($http, mvIdentity, $q) {
    return {
        authenticateUser: function(username, password) {
            var dfd = $q.defer();

            $http.post('/api/account/login', { username:username, password:password }).then(function(response){
                if(response.data && response.data.success) {
                    mvIdentity.currentUser = { username:username, token:response.data.token }
                    localStorage.setItem('currentUser', JSON.stringify(mvIdentity.currentUser));
                    dfd.resolve(true);
                } else {
                    dfd.resolve(false);
                }
            })
            return dfd.promise;
        },
        signOut: function() {
            var dfd = $q.defer();
            $http.post('/api/account/logout', {}).then(function(response){
                mvIdentity.currentUser = undefined;
                localStorage.removeItem('currentUser');
                dfd.resolve(true);
            });
            return dfd.promise;
        },
        refreshLogin: function() {
            var dfd = $q.defer();
            
            $http.post('/api/account/refresh', {}).then(function(response){
                if(response.data.success) {
                    mvIdentity.currentUser.token = response.data.token
                    localStorage.setItem('currentUser', JSON.stringify(mvIdentity.currentUser));
                    dfd.resolve(true);
                } else {
                    dfd.resolve(false);
                }
            })
            return dfd.promise;
        }
    }
});
