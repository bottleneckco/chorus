import { createStore, applyMiddleware, compose } from 'redux';
import { createLogger } from 'redux-logger';
import thunkMiddleware from 'redux-thunk';
import { persistStore, autoRehydrate } from 'redux-persist';

import hydrationCompleted from './actions/action-persistence';
import rootReducer from './reducers/reducer-index';

const loggerMiddleware = createLogger();

// eslint-disable-next-line no-underscore-dangle
const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

/* Use
  autoRehydrate({
    log: true,
  }),

  for redux-persist logging
*/

const store = createStore(
  rootReducer,
  composeEnhancers(
    applyMiddleware(
      thunkMiddleware,
      loggerMiddleware
    ),
    autoRehydrate()
  )
);

export const persistor = persistStore(store, {}, () => {
  store.dispatch(hydrationCompleted());
});

export default store;
