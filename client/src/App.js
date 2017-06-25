import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import Landing from './containers/Landing';
import Onboard from './containers/Onboard';
import Channel from './containers/Channel';

import './stylesheets/normalize.css';
import './stylesheets/base.scss';
import './stylesheets/controls.scss';

const App = () => (
  <BrowserRouter>
    <div className="container">
      <Switch>
        <Route
          exact
          path="/"
          render={(props) => <Landing {...props} />}
        />
        <Route
          path="/onboard"
          render={(props) => <Onboard {...props} />}
        />
        <Route
          path="/:hash"
          render={(props) => <Channel {...props} />}
        />
      </Switch>
    </div>
  </BrowserRouter>
);

export default App;
