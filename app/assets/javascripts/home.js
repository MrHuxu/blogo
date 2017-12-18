$(() => {
  $('.home-item-content').hover(function () {
    $(this).animateCss('pulse');
    $(this).animate({ backgroundColor: "#EFEFEF" }, 200);
  }, function () {
    $(this).animate({ backgroundColor: "#FFFFFF" }, 100);
  });

  $('.snippet-arrow').click(() => {
    $('body').animate({ scrollTop: window.innerHeight });
  });
});