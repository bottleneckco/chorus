import { combineReducers } from 'redux';
import channel from './reducer-channel';
import search from './reducer-search';
import queue from './reducer-queue';

const rootReducer = combineReducers({
  channel: channel(),
  search: search(),
  queue: queue()
});

export default rootReducer;
