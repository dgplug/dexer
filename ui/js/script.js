$(document).ready(function(){
    $(".search-field").keyup(function(){
        $.ajax({url: "/search/"+$(".search-field").val(), success: function(result){
            $(".search-results").html(result);
        }});
    });
});