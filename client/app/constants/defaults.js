angular.module('golang')
	.constant('DEFAULTS', function() {
  		return {
			"task1": {
				"width": 8,
				"height": 4,
				"symbol": "#"
			}, 
			"task2": [{
				"width": 8,
				"height": 5
			}, {
				"width": 6,
				"height": 9
			}]
		};
	}());