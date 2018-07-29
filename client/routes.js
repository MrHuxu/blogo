import Home from './components/home';
import Post from './components/post';
import Tags from './components/tags';
import NoMatch from './components/404';

export default [
  { path: '/page/:page', component: Home },
  { path: '/post/:title', component: Post },

  { path: '/tags', component: Tags },

  { path: '/404', component: NoMatch }
];
