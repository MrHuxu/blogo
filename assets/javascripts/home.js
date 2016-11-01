$(() => {
  $('.snippet-container').hover(function () {
    $(this).animate({
      backgroundColor: "#EFEFEF",
      width: '50%'
    }, 200);
  }, function () {
    $(this).animate({
      backgroundColor: "#FFFFFF",
      width: '49%'
    }, 100);
  });
});