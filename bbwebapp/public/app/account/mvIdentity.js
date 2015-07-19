angular.module('app').factory('mvIdentity', function($window) {
    var currentUser = undefined;
    var locallyStoredUser = localStorage.getItem('currentUser');
    if(!!locallyStoredUser) {
        currentUser = JSON.parse(locallyStoredUser);
    }
    return {
        currentUser: currentUser,
        isAuthenticated: function () {
            return !!this.currentUser;
        },
        isAuthorized: function () {
            return !!this.currentUser;
        }
    }
});
