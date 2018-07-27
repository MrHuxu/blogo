const { resolve: resolvePath } = require('path');
const { readdirSync, open, close } = require('fs');
const { createInterface } = require('readline');
const { info } = require('better-console');

const rl = createInterface({
  input  : process.stdin,
  output : process.stdout
});

const getSequence = post => new Promise(resolve => {
  const nextSeq = readdirSync(resolvePath(__dirname, '../archives')).filter(
    file => file.endsWith('.md')
  ).map(file => parseInt(/\d+/.exec(file.split('*')[0])[0])).sort(
    (a, b) => a > b ? -1 : 1
  )[0] + 1;

  rl.question(`The sequence of the post:\n[Default: ${nextSeq}] `, seq => {
    seq = seq.trim();
    post.seq = seq && parseInt(seq) ? seq : nextSeq;
    resolve(post);
  });
});

const getTitle = post => new Promise(resolve => {
  rl.question('The title of the post:\n[Default: Placeholder] ', title => {
    title = title.trim();
    post.title = title.length ? title : 'Placeholder';
    resolve(post);
  });
});

const formatMonthDate = num => num >= 10 ? num : ('0' + num);

const getDate = post => new Promise(resolve => {
  const date = new Date();
  const nowDate = `${date.getUTCFullYear()}${formatMonthDate(date.getUTCMonth() + 1)}${formatMonthDate(date.getUTCDate())}`;
  const validateDate = str => (
    8 === str.length &&
    parseInt(str.slice(0, 4)) &&
    parseInt(str.slice(4, 6)) && parseInt(str.slice(4, 6)) <= 12 &&
    parseInt(str.slice(6)) && parseInt(str.slice(6)) <= 31
  );

  rl.question(`The date when you write the post:\n[Default: ${nowDate}] `, date => {
    date.trim();
    post.date = validateDate(date) ? date : nowDate;
    resolve(post);
  });
});

const getTags = post => new Promise(resolve => {
  rl.question('The tags of the post:\n[Default: Placeholder] ', tags => {
    tags = tags.trim();
    post.tags = tags.length ? tags.split(' ') : ['Placeholder'];
    resolve(post);
    rl.close();
  });
});

const touchFile = post => {
  const { seq, title, date, tags } = post;
  const fileName = 'WIP: ' + [
    seq, title, date, tags.join('-')
  ].join('*') + '.md';
  open(resolvePath(__dirname, '../archives/', fileName), 'w', (_, file) => close(file, () => {
    info(`\n[ ${fileName} ] successfully created!`);
  }));
};

getSequence({}).then(
  post => getTitle(post)
).then(
  post => getDate(post)
).then(
  post => getTags(post)
).then(
  post => touchFile(post)
);
