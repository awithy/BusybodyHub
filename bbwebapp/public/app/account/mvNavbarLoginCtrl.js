//angular.module('app').controller('mvNavbarLoginCtrl', function($scope, $http, $location, mvIdentity, mvAuth, mvNotifier){
angular.module('app').controller('mvNavbarLoginCtrl', function($scope, $http, mvIdentity, mvAuth, mvNotifier){
    $scope.identity = mvIdentity;
    $scope.user = {};
    $scope.user.username = "";
    $scope.user.password = "";

    $scope.update = function(user) {
        mvAuth.authenticateUser(user.username, user.password).then(function(success){
            if(success) {
                mvNotifier.notify('Logged in!')
                $scope.my=true;
            }
            else {
                mvNotifier.error('Not logged in');
            }
        })
    };

    $scope.signOut = function() {
        mvAuth.signOut().then(function () {
            $scope.user.username = "";
            $scope.user.password = "";
            mvNotifier.notify('You have been logged out!');
            //$location.path('/');
        })
    }
});
