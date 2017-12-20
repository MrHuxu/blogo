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
  homeHeader.on('mousemove', ({ pageX, pageY }) => {
    homeHeader.css('background-position', `${pageX / innerWidth * 14}% ${pageY / innerHeight * 8}%`);
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
    const content = `
      <div class="home-item-content" style="background-color: rgb(255, 255, 255);">
        <div class="snippet-container">
          <a class="snippet-header" href="/post/${post.title}"> ${post.title} </a>

          <div class="ui divider snippet-divider"></div>

          <div class="snippet-content">
            ${post.content}
          </div>

          <div class="tags">
            ${tags}
          </div>
          <div class="date">
            Dec 26, 2016
            ${post.showDate}
          </div>
        </div>
      </div>
    `;
    return `
      <div class="home-item">
        <div class="home-item-container-left">
          ${($maxPostSeq - post.seq) % 2 ? img : content}
        </div>
        <div class="home-item-container-right">
          ${($maxPostSeq - post.seq) % 2 ? content : img}
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