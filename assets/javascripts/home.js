$(() => {
  $('.snippet-container').hover(function () {
    $(this).animate({
      backgroundColor: "#EFEFEF",
      width: '49%'
    }, 200);
  }, function () {
    $(this).animate({
      backgroundColor: "#FFFFFF",
      width: '48%'
    }, 100);
  });
});