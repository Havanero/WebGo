//http://classify.oclc.org/classify2/Classify?&summary=true&title=
//	templates := template.Must(template.ParseFiles("templates/base.html", "templates/web-go/index.html"))
//= template "templates/base.html"

function deleteBook(pk){
  console.log("Deteling ", pk);
  $.ajax({
    method:"GET",
    url: "/books/delete?pk="+pk,
    success: function(){
      $("#book-row-"+ pk).remove();
    }
  })
}
function showSearchPage(){
  $("#search-page").show();
  $("#view-page").hide();
}
function showViewPage(){
  $("#search-page").hide();
  $("#view-page").show();
}
function submitSearch(){
  $.ajax({
    url:"/search",
    method: "POST",
    data: $("#search-form").serialize(),
    success: function(rawData){
      var parsed = JSON.parse(rawData);
      if (!parsed)return
      var SearchResults = $("#search-results");
      SearchResults.empty();
      parsed.forEach(function(result){
        var row = $("<tr><td>"+ result.Title + "</td><td>" + result.Author + "</td><td>" + result.Year +
        "</td><td>"+ result.ID + "</td></tr>");
        SearchResults.append(row);
        row.on("click", function(){
          $.ajax({
            url:"books/add?id="+result.ID,
            method: "GET",
            success: function(data){
              var book = JSON.parse(data)
              if (!book)return;
              $("#view-results").append("<tr id='book-row-"+ book.PK +"'><td>"+ book.Title +
              "</td><td>" + book.Author + "</td><td>" +
               book.Classification +
               "</td><td><button class='delete-btn' onclick='deleteBook(" + book.PK + ")'>Delete</button></td></tr>");
            }
          })
        })
      });
    }
  });
  return false;
}
