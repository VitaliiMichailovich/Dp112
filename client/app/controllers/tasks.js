angular.module('golang')
	.controller('TasksCtrl', function ($scope, tasks) {
		$scope.tasks = tasks.tasks;
		$scope.data = '';
		$scope.error = '';

		$scope.runTask = function (id) {
			console.log('id');
			tasks.runTask(id).then(function (data) {
				console.log(data);
				$scope.data = data.data.resp;
				$scope.error = data.data.reason;
			});
		};

		$scope.runAll = function () {
			tasks.runAll().then(function (data) {
				console.log(data);
				$scope.data = '';
				$scope.error = '';
				data.data.forEach(function (response, i) {
					if (response.resp) {
						$scope.data += response.resp + '\n';
					}
					if (response.reason) {
						$scope.error += response.reason + '\n';
					}
				});
			});
		}
	});