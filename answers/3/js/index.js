
var app = angular.module('myApp', ['ui.bootstrap']);
app.controller('myCtrl', function($scope, $http, $timeout, $interval, $q) {
        
        console.log("index.js loaded")
        var url = location.protocol+"//localhost/validate_json";
        $scope.textModel = ""


        $scope.get_json_from_textbox = function()
        {
            var request = {}
            request["request"] = "validate_json";
            request["json-data"] = $scope.textModel;

            $http({
                url:url,
                method:'POST',
                data : JSON.stringify(request)
            }).then(
                function(response) {
                    $("#error-msg").html(response.data);
                    console.log(response.data)
                },
                function(response) {
                    console.error("mod_WSGI or APACHE not configured or installed OR server side error");
                }) 

        }
});
