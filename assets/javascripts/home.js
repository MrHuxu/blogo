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

  $('.paginate-container').height($('.snippet-container').last().height());

  var containerHeight = $('.paginate-prev').height() || $('.paginate-next').height();
  var selfHeight = $('.paginate-prev i').height() || $('.paginate-next i').height();
  var containerWidth = $('.paginate-prev').width() || $('.paginate-next').width();
  var selfWidth = $('.paginate-prev i').width() || $('.paginate-next i').width();
  $('.paginate-prev i, .paginate-next i').css({
    top: containerHeight / 2 - selfHeight / 2,
    left: containerWidth / 2 - selfWidth / 2
  });

  $('.paginate-prev, .paginate-next').hover(function () {
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