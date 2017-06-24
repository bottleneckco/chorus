import * as types from '../constants/action-types';
import API_ROOT from '../constants/api-url';

const fetchQueueRequest = () => ({
  type: types.FETCH_CHANNEL_QUEUE_REQUEST
});

const fetchQueueSuccess = () => ({
  type: types.FETCH_CHANNEL_QUEUE_SUCCESS
});

const fetchQueueFailure = () => ({
  type: types.FETCH_CHANNEL_QUEUE_FAILURE
});

export const fetchQueue = () => (dispatch) => {
  
}