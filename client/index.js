import React from 'react';
import { renderToString } from 'react-dom/server';

import { Provider } from 'react-redux';
import { createStaticStore } from './store';

import { Switch, Route, StaticRouter } from 'react-router';
import routes from './routes';

import { ServerStyleSheet, StyleSheetManager } from 'styled-components';
const sheet = new ServerStyleSheet();

export const renderHtmlString = (url, data) => {
  const html = renderToString(
    <StyleSheetManager sheet={ sheet.instance }>
      <Provider store={ createStaticStore(data) }>
        <StaticRouter location={ url }>
          <Switch>
            { routes.map(route => (
              <Route { ...route } />
            )) }
          </Switch>
        </StaticRouter>
      </Provider>
    </StyleSheetManager>
  );
  const styles = sheet.getStyleTags();
  return html + styles;
};
