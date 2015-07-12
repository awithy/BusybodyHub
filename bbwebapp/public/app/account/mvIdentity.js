angular.module('app').factory('mvIdentity', function($window) {
    var currentUser = undefined;
    if(!!$window.bootstrappedUserObject){
        currentUser = {};
        angular.extend(currentUser, $window.bootstrappedUserObject);
    }
    return {
        currentUser: currentUser,
        isAuthenticated: function () {
            console.log("auth:" + !!this.currentUser);
            return !!this.currentUser;
        },
        isAuthorized: function () {
            return !!this.currentUser;
        }
    }
});
