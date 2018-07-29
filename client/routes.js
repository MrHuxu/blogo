import Home from './components/home';
import Tags from './components/tags';
import NoMatch from './components/404';

export default [
  { path: '/page/:page', component: Home },
  { path: '/tags', component: Tags },

  { path: '/404', component: NoMatch }
];
