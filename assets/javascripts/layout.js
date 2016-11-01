$(() => {
  var $toTop = $('#back-to-top');

  $toTop.click(() => {
    $('body').animate({ scrollTop: 0 });
  });

  $(window).scroll(() => {
    if ($('body').scrollTop() > 10) {
      $toTop.show(400);
    } else {
      $toTop.hide(400);
    }
  });
});