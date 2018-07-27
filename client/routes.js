import Home from './components/home';
import Test from './components/test';
import NoMatch from './components/404';

export default [
  { exact: true, path: '/', component: Home },
  { path: '/test', component: Test },

  { path: '/404', component: NoMatch }
];
