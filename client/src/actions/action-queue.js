import * as types from '../constants/action-types';
import API_ROOT from '../constants/api-url';

const fetchQueueRequest = () => ({
  type: types.FETCH_CHANNEL_QUEUE_REQUEST
});

const fetchQueueSuccess = (data) => ({
  data,
  type: types.FETCH_CHANNEL_QUEUE_SUCCESS
});

const fetchQueueFailure = () => ({
  type: types.FETCH_CHANNEL_QUEUE_FAILURE
});

export const fetchQueue = (queueId) => (dispatch) => {
  dispatch(fetchQueueRequest());
  return fetch(`${API_ROOT}/api/channels/${queueId}/queue`, {
    method: 'GET'
  }).then((response) => {
    return response.json();
  }).then((json) => {
    dispatch(fetchQueueSuccess(json));
  }).catch((err) => {
    dispatch(fetchQueueFailure(err));
  });
}