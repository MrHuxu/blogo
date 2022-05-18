import React from 'react';

import { Switch, Route } from 'react-router';
import routes from '../routes';

import { Container, App as AppContainer, Title, TitleArrow, BackToTopContainer, BackToTop, Footer } from './layout/elements';

const App = () => (
  <Container>
    <AppContainer>
      <Title>
        Life of xhu
        <TitleArrow href="/"><i className="icon-double-right-arrow" style={{ fontSize: 14, color: 'gray' }} /></TitleArrow>
      </Title>

      <Switch>
        {routes.map(route => (
          <Route {...route} />
        ))}
      </Switch>
    </AppContainer>

    <BackToTopContainer>
      <BackToTop id="back-to-top" />
    </BackToTopContainer>

    <Footer>
      <p>
        Copyright © 2022 -&nbsp;
        <a href="mailto:hxtheone@gmail.com">xhu</a>
        &nbsp;- Powered by&nbsp;
        <a target="_blank" href="https://github.com/gin-gonic/gin">Gin</a>,&nbsp;
        <a target="_blank" href="http://jquery.com/">jQuery</a>,&nbsp;
        <a target="_blank" href="https://daneden.github.io/animate.css/">Animate.css</a>
      </p>
    </Footer>
  </Container>
);

export default App;
