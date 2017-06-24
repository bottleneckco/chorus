import { combineReducers } from 'redux';
import channel from './reducer-channel';

const rootReducer = combineReducers({
  channel: channel()
});

export default rootReducer;
