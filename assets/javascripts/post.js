const songs = {
  562456: '强く儚い者たち - 真喜志智子',
  27571802: '我的钢琴很简单 (钢琴曲) - 雷诺儿',
  28427271: 'My Goodbye - Brian',
  28737747: 'Lost Stars - Adam Levine',
  34916769: 'Legacy - Shire Music/Really Slow Motion',
  5256469: '大哥 - 卫兰',
  4975665: 'Butter-Fly - V.A.',
  27646199: '斑马，斑马 - 宋冬野',
  27533158: 'Fingerprints - Kari Kimmel',
  108983: '会有那么一天 - 林俊杰',
  496549: 'Chiru (Saisei no Uta) - Robert de Boron',
  276294: '相思 - 毛阿敏',
  211258: '乱红(笛) - 陈悦',
  306664: '身骑白马 - 徐佳莹',
  29816798: 'HUMAN - Ken Arai',
  26093260: '千本桜 - まらしぃ',
  26664325: 'Main Title - Ramin Djawadi',
  190473: '秋意浓 - 张学友',
  3560431: 'Stay - Tonya Mitchell'
};
const songIds = Object.keys(songs);

$(() => {
  $('p>img').each((index, img) => {
    $(img.parentElement).css({ textAlign: 'center' });
  });

  setTimeout(() => {
    var songId = songIds[parseInt(Math.random() * 100) % songIds.length];
    var playerDom = $(`<iframe id="wangyi-player" frameborder="no" border="0" marginwidth="0" marginheight="0" width=280 height=86 src="http://music.163.com/outchain/player?type=2&id=${songId}&auto=1&height=66"></iframe>`)
    playerDom.appendTo($('.post-container'));
  }, 1400);
});