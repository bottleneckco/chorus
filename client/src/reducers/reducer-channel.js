import { combineReducers } from 'redux';
import * as types from '../constants/action-types';

const channel = () => {
  const isCreating = (state = false, action) => {
    switch (action.type) {
      case types.CREATE_CHANNEL_REQUEST:
        return true;
      case types.CREATE_CHANNEL_SUCCESS:
      case types.CREATE_CHANNEL_FAILURE:
        return false;
      default:
        return state;
    }
  };

  const response = (state = {}, action) => {
    switch (action.type) {
      case types.CREATE_CHANNEL_SUCCESS:
        return action.response.channel;
      default:
        return state;
    }
  };

  return combineReducers({
    isCreating,
    response
  });
};

export default channel;

export const getIsCreating = (state) => state.channel.isCreating;
export const getResponse = (state) => state.channel.response;
