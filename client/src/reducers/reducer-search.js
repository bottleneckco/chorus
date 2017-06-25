import { combineReducers } from 'redux';
import * as types from '../constants/action-types';

const search = () => {
  const isFetching = (state = false, action) => {
    switch (action.type) {
      case types.SEARCH_MUSIC_REQUEST:
        return true;
      case types.SEARCH_MUSIC_SUCCESS:
      case types.SEARCH_MUSIC_FAILURE:
        return false;
      default:
        return state;
    }
  };

  const results = (state = [], action) => {
    switch (action.type) {
      case types.SEARCH_MUSIC_SUCCESS:
        return action.data;
      default:
        return state;
    }
  };

  return combineReducers({
    isFetching,
    results
  });
};

export default search;

export const getSearchIsFetching = (state) => state.search.isFetching;
export const getResults = (state) => state.search.results;
