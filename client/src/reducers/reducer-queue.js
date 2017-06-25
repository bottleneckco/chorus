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
        return action.data.queue;
      default:
        return state;
    }
  };

  // Handle errors from fetchQueue, addToQueue, SkipQueue
  const fetchError = (state = '', action) => {
    switch (action.type) {
      case types.FETCH_CHANNEL_QUEUE_REQUEST:
      case types.FETCH_CHANNEL_QUEUE_SUCCESS:
        return '';
      case types.FETCH_CHANNEL_QUEUE_FAILURE:
        return action.error;
      default:
        return state;
    }
  };

  const addError = (state = '', action) => {
    switch (action.type) {
      case types.ADD_TO_CHANNEL_QUEUE_REQUEST:
      case types.ADD_TO_CHANNEL_QUEUE_SUCCESS:
        return '';
      case types.ADD_TO_CHANNEL_QUEUE_FAILURE:
        return action.error;
      default:
        return state;
    }
  };

  const skipError = (state = '', action) => {
    switch (action.type) {
      case types.SKIP_QUEUE_ITEM_REQUEST:
      case types.SKIP_QUEUE_ITEM_SUCCESS:
        return '';
      case types.SKIP_QUEUE_ITEM_FAILURE:
        return action.error;
      default:
        return state;
    }
  };

  return combineReducers({
    isFetching,
    data,
    fetchError,
    addError,
    skipError
  });
};

export default queue;

export const getQueueIsFetching = (state) => state.queue.isFetching;
export const getQueue = (state) => state.queue.data;
