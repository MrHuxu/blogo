const { createInterface } = require('readline');

const rl = createInterface({
  input  : process.stdin,
  output : process.stdout
});

const getSequence = post => new Promise(res => {
  rl.question('The sequence of the post:\n[Default: 95]', seq => {
    console.log();
    post.seq = seq;
    res(post);
  });
});

const getTitle = post => new Promise(res => {
  rl.question('The title of the post:\n', title => {
    console.log();
    post.title = title;
    res(post);
  });
});

const getDate = post => new Promise(res => {
  rl.question('The date when you write the post:\n[Default: 20171220]', date => {
    console.log();
    post.date = date;
    res(post);
  });
});

const getTags = post => new Promise(res => {
  rl.question('The tags of the post:\n', tags => {
    console.log();
    post.tags = tags.split(' ');
    res(post);
    rl.close();
  });
});

getSequence({}).then(
  post => getTitle(post)
).then(
  post => getDate(post)
).then(
  post => getTags(post)
).then(
  (post) => console.log(post)
);