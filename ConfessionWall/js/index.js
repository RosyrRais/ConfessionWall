$(function(){
	$("#top a").mouseover(function(){
		$(this).stop(true, false).animate({
			"background-color":"#c14000",
			"border-color":"#f37126",
			"color":"#fcd2d1"
			}, 200);
	})
	$("#top a").mouseout(function(){
		$(this).stop(true, false).animate({
			"background-color":"#e5a65e",
			"border-color":"#c14000",
			"color":"#000"
			}, 200);
	})
	$(".content").mouseover(function(){
		$(this).stop(true, false).animate({
			"opacity":"0.8"
			}, 100);
	})
	$(".content").mouseout(function(){
		$(this).stop(true, false).animate({
			"opacity":"0.6"
			}, 100);
	})
})