function keyuphandler(){
    $(".search-field").keyup(search);
}

function search(){
    $.ajax({url: "/search/"+$(".search-field").val(), success: receiver});
}

function receiver(result){
    var json = jQuery.parseJSON(result);
    if(json.length === 0){
        $(".search-result-header").html("Sorry no results found.");
        return;
    }
    $(".search-result-header").html("Results found: " + json.length);
    var temp = "";
    for(var i in json){
        if(i === 1){
            temp = '<div class="entry">' + JSON.stringify(json[i].id) + "</div>";
        }
        temp += '<div class="entry">' + JSON.stringify(json[i].id) + "</div>";
    }
    $(".search-results").html(temp);
}

$(document).ready(keyuphandler);