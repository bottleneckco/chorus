import React from 'react';
import { BrowserRouter, Route } from 'react-router-dom';

import Landing from './containers/Landing';
import Onboard from './containers/Onboard';

import './stylesheets/normalize.css';
import './stylesheets/base.scss';

const App = () => (
  <BrowserRouter>
    <div>
      <Route
        exact
        path="/"
        render={(props) => <Landing {...props} />}
      />
      <Route
        path="/onboard"
        render={(props) => <Onboard {...props} />}
      />
    </div>
  </BrowserRouter>
);

export default App;
