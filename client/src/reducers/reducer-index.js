import { combineReducers } from 'redux';
import channel from './reducer-channel';
import search from './reducer-search';
import queue from './reducer-queue';
import rehydration from './reducer-persistence';

const rootReducer = combineReducers({
  channel: channel(),
  search: search(),
  queue: queue(),
  persistence: rehydration
});

export default rootReducer;
