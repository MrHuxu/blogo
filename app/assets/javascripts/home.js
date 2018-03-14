$(() => {
  const enableHomeItemHoverAnimation = () => {
    $('.home-item-content').hover(function () {
      $(this).animateCss('pulse');
      $(this).animate({ backgroundColor: "#EFEFEF" }, 200);
    }, function () {
      $(this).animate({ backgroundColor: "#FFFFFF" }, 100);
    });
  };
  enableHomeItemHoverAnimation();

  const homeHeader = $('.home-header');
  const homeHeaderBgPic = $('.home-header > .bg-pic');
  const { innerWidth, innerHeight } = window;
  const bgImg = $('.bg-pic > img');
  const imgW = bgImg.width(), imgH = bgImg.height();
  const subW = Math.min(imgW - innerWidth, innerHeight / 18), subH = Math.min(imgH - innerHeight, innerHeight / 18);
  homeHeader.on('mousemove', ({ pageX, pageY }) => {
    bgImg.css('transform', `translate3d(-${pageX / innerWidth * subW}px, -${pageY / innerHeight * subH}px, 0px)`);
  });

  const snippetArrow = $('.snippet-arrow');
  snippetArrow.click(() => {
    $('body').animate({ scrollTop: window.innerHeight });
  });

  const homeContainer = $('.home-container');
  const homeAppend = $('.home-append');
  const loadMore = $('.load-more');
  let $currentPage = 0;

  const convertPostToElement = post => {
    const img = `
      <img class="home-item-pic" src="https://raw.githubusercontent.com/MrHuxu/img-repo/master/blog-title/${post.seq < 13 ? ($maxPostSeq - post.seq) : post.seq}.jpg" />
    `;
    const tags = post.tags.map(tag => (
      `<a href="/archives?tag=${tag}">${tag}</a>`
    )).join('&nbsp;<div class="tag-divider"></div>&nbsp;');
    return `
      <div class="home-item">
        <div class="home-item-content">
          <div class="date">
          ${post.showDate}
          </div>
          <a class="link" href="/post/${post.title}">${post.title}</a>
        </div>
      </div>
    `;
  };

  const loadNextPage = () => {
    loadMore.text('Loading ...');

    $.get(`/page/${$currentPage + 1}`, (data, status) => {
      if ('success' === status) {
        const { canBeAppend, currentPage, titles, posts } = data;

        $currentPage = currentPage;

        homeAppend.remove();
        homeContainer.append(...titles.map(title => convertPostToElement(posts[title])));
        enableHomeItemHoverAnimation();
        homeAppend.appendTo(homeContainer);

        if (canBeAppend) {
          loadMore.on('click', loadNextPage);
          loadMore.text('More');
        } else {
          loadMore.text('All Loaded');
          loadMore.unbind('click');
          loadMore.css({ cursor: 'not-allowed' });
        }
      }
    });
  };

  loadMore.on('click', loadNextPage);
});