import { REHYDRATE } from 'redux-persist/constants';
import * as types from '../constants/action-types';

const rehydration = (state = { rehydrated: false }, action) => {
  switch (action.type) {
    case REHYDRATE:
      return {
        ...state,
        rehydrated: false
      };

    case types.HYDRATION_COMPLETED:
      return {
        ...state,
        rehydrated: true
      };

    default:
      return state;
  }
};

export default rehydration;

export const getRehydrationStatus = (state) => state.persistence.rehydrated;
