
var app = angular.module('myApp', []);
app.controller('userList', function($scope, $http) {

//http://localhost:51911/Home/AngularData

var d = new Date();
var start = d.getTime();
$scope.startTime = start;

    $scope.initTable = function(){
        $http.get("http://localhost:8035/getAllUser")
        .then(function(response) {
            
            $scope.users = response.data;
            var e = new Date();
            $scope.endTime = e.getTime();
            $scope.fark = e.getTime() - start;
        });
    }



    $scope.initTable();



        $http.defaults.headers.post["Content-Type"] = "text/plain; charset=utf-8";

    $scope.formData = {};

    $scope.processForm = function() {
        $http.post("http://localhost:8035/insertUser", $scope.formData)
        .then(function(data){
           console.log(data);
           $scope.initTable();
            $scope.formData = {};
        });
    };
});
