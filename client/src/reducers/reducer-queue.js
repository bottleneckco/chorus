import { combineReducers } from 'redux';
import * as types from '../constants/action-types';

const queue = () => {
  const isFetching = (state = false, action) => {
    switch (action.type) {
      case types.FETCH_CHANNEL_QUEUE_REQUEST:
        return true;
      case types.FETCH_CHANNEL_QUEUE_SUCCESS:
      case types.FETCH_CHANNEL_QUEUE_FAILURE:
        return false;
      default:
        return state;
    }
  };

  const data = (state = [], action) => {
    switch (action.type) {
      case types.FETCH_CHANNEL_QUEUE_SUCCESS:
        return action.data;
      default:
        return state;
    }
  };

  return combineReducers({
    isFetching,
    data
  });
};

export default queue;
