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

  const isFetching = (state = false, action) => {
    switch (action.type) {
      case types.FETCH_CHANNEL_REQUEST:
        return true;
      case types.FETCH_CHANNEL_SUCCESS:
      case types.FETCH_CHANNEL_FAILURE:
        return false;
      default:
        return state;
    }
  };

  const data = (state = {}, action) => {
    switch (action.type) {
      case types.CREATE_CHANNEL_SUCCESS:
        return action.data.channel;
      case types.FETCH_CHANNEL_SUCCESS:
        return action.data.channel;
      case types.ADD_USER_TO_CHANNEL_SUCCESS:
        return {
          ...state,
          users: [
            ...state.users,
            action.data
          ]
        };
      default:
        return state;
    }
  };

  return combineReducers({
    isCreating,
    isFetching,
    data
  });
};

export default channel;

export const getChannelIsCreating = (state) => state.channel.isCreating;
export const getChannelIsFetching = (state) => state.channel.isFetching;
export const getChannelData = (state) => state.channel.data;
