import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import Landing from './containers/Landing';
import Onboard from './containers/Onboard';
<<<<<<< HEAD
import Channel from './containers/Channel';
=======
import WsTest from './containers/WsTest';
>>>>>>> Add ws to react

import './stylesheets/normalize.css';
import './stylesheets/base.scss';
import './stylesheets/controls.scss';

const App = () => (
  <BrowserRouter>
    <div className="container">
<<<<<<< HEAD
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
=======
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
        path="/ws-test"
        render={(props) => <WsTest {...props} />}
      />
>>>>>>> Add ws to react
    </div>
  </BrowserRouter>
);

export default App;
