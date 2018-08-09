import React from 'react';

import { Switch, Route } from 'react-router';
import routes from '../routes';

import { Container, App as AppContainer, Footer } from './layout/elements';

const App = () => (
  <Container>
    <AppContainer>
      <div className="title">
          Life of xhu
        <a href="/"><i className="angle double right icon snippet-arrow" /></a>
      </div>

      <Switch>
        { routes.map(route => (
          <Route { ...route } />
        )) }
      </Switch>
    </AppContainer>

    <Footer>
      <p>
          Copyright © 2018 -&nbsp;
        <a href="mailto:hxtheone@gmail.com">xhu</a>
          &nbsp;- Powered by&nbsp;
        <a target="_blank" href="https://github.com/gin-gonic/gin">Gin</a>,
        <a target="_blank" href="http://jquery.com/">jQuery</a>,
        <a target="_blank" href="https://daneden.github.io/animate.css/">Animate.css</a>,
        <a target="_blank" href="http://semantic-ui.com/">Semantic UI</a>
      </p>
    </Footer>
  </Container>
);

export default App;
