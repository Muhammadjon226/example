$(document).on("click", ".show-more-button", function() {
    if ($(this).text() == "Show more") {
      $(this).text("Show less");
      $(this).closest('table').children(".show-more").css('display', 'block');
    } else {
      $(this).text("Show more");
      $(this).closest('table').children(".show-more").css('display', 'none');
    } 
  });