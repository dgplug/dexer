function keyuphandler(){
    $(".search-field").keyup(search);
}

function search(){
    var searchValue = $(".search-field").val();
    if(searchValue === ""){
        $(".search-result-header").html("Type in the box above");
        $(".search-results").html("");
    }
    $.ajax({url: "/search/"+searchValue, success: receiver});
}

function receiver(result){
    var json = jQuery.parseJSON(result);
    if(json.length === 0){
        $(".search-result-header").html("Sorry no results found.");
        $(".search-results").html("");
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
