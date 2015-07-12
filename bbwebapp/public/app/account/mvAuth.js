angular.module('app').factory('mvAuth', function($http, mvIdentity, $q) {
    return {
        authenticateUser: function(username, password) {
            var dfd = $q.defer();

            //TODO: Temp
            mvIdentity.currentUser = { username:username };
            dfd.resolve(true);
            return dfd.promise;

            $http.post('/api/login', { username:username, password:password }).then(function(response){
                if(response.data.success) {
                    mvIdentity.currentUser = response.data;
                    dfd.resolve(true);
                } else {
                    dfd.resolve(false);
                }
            })
            return dfd.promise;
        },
        signOut: function() {
            var dfd = $q.defer();

            //TODO: Temp
            mvIdentity.currentUser = undefined;
            dfd.resolve(true);
            return dfd.promise;

            $http.post('/api/logout', {}).then(function(response){
                if(response.data.success) {
                    mvIdentity.currentUser = undefined;
                    dfd.resolve(true);
                } else {
                    dfd.resolve(false);
                }
            })
            return dfd.promise;
        },
        refreshLogin: function() {
            var dfd = $q.defer();
            
            //TODO: Temp
            mvIdentity.currentUser = { username:username };
            dfd.resolve(true);
            return dfd.promise;

            $http.post('/api/refreshLogin', {}).then(function(response){
                if(response.data.success) {
                    mvIdentity.currentUser = response.data;
                    dfd.resolve(true);
                } else {
                    dfd.resolve(false);
                }
            })
            return dfd.promise;
        }
    }
});
