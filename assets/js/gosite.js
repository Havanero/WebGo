28615701

function apppendBook(book){
  $("#view-results").append("<tr id='book-row-"+ book.PK +"'><td>"+ book.Title +
  "</td><td>" + book.Author + "</td><td>" +
   book.Classification +
   "</td><td><button class='delete-btn' onclick='deleteBook(" + book.PK + ")'>Delete</button></td></tr>");
}
function deleteBook(pk){
  console.log("Deteling ", pk);
  $.ajax({
    method:"DELETE",
    url: "/books/"+pk,
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
            url:"/books?id="+result.ID,
            method: "PUT",
            success: function(data){
              var book = JSON.parse(data)
              if (!book)return;
              apppendBook(book);
            }
          })
        })
      });
    }
  });
  return false;
}
function sortBooks(columName){
  console.log("js sort invoked on " + columName);
  $.ajax({
    method:"GET",
    url:"/books?sortBy="+ columName,
    success: function(result){
      var books = JSON.parse(result);
      if(!books)return
      $("#view-results").empty()
      books.forEach(function(book){
        apppendBook(book)
      });
    }
  })
}
