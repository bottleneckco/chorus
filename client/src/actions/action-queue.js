import * as types from '../constants/action-types';
import API_ROOT from '../constants/api-url';

const fetchQueueRequest = () => ({
  type: types.FETCH_CHANNEL_QUEUE_REQUEST
});

const fetchQueueSuccess = (data) => ({
  data,
  type: types.FETCH_CHANNEL_QUEUE_SUCCESS
});

const fetchQueueFailure = (error) => ({
  error,
  type: types.FETCH_CHANNEL_QUEUE_FAILURE
});

export const fetchQueue = (queueId) => (dispatch) => {
  dispatch(fetchQueueRequest());
  return fetch(`${API_ROOT}/api/channels/${queueId}/queue`, {
    method: 'GET'
  }).then((response) => (
    response.json()
  )).then((json) => {
    dispatch(fetchQueueSuccess(json));
  }).catch((err) => {
    dispatch(fetchQueueFailure(err));
  });
};

const addToQueueRequest = () => ({
  type: types.ADD_TO_CHANNEL_QUEUE_REQUEST
});

const addToQueueSuccess = () => ({
  type: types.ADD_TO_CHANNEL_QUEUE_SUCCESS
});

const addToQueueFailure = (error) => ({
  error,
  type: types.ADD_TO_CHANNEL_QUEUE_FAILURE
});

export const addToQueue = (channelId, data) => (dispatch) => {
  dispatch(addToQueueRequest());
  return fetch(`${API_ROOT}/api/channels/${channelId}/queue`, {
    method: 'POST',
    body: JSON.stringify(data),
    headers: {
      'Content-Type': 'application/json'
    }
  }).then((response) => (
    response.json()
  )).then((json) => {
    if (json.status === 'ok') {
      dispatch(addToQueueSuccess());
    } else {
      dispatch(addToQueueFailure(json));
    }
  }).catch((err) => {
    dispatch(addToQueueFailure(err));
  });
};

const skipQueueItemRequest = () => ({
  type: types.SKIP_QUEUE_ITEM_REQUEST
});

const skipQueueItemSuccess = () => ({
  type: types.SKIP_QUEUE_ITEM_SUCCESS
});

const skipQueueItemFailure = (error) => ({
  error,
  type: types.SKIP_QUEUE_ITEM_FAILURE
});

// index starts from 0
export const skipQueue = (channelId, index) => (dispatch) => {
  dispatch(skipQueueItemRequest());
  return fetch(`${API_ROOT}/api/channels/${channelId}/queue/${index}`, {
    method: 'DELETE'
  }).then((response) => (
    response.json()
  )).then((json) => {
    if (json.status === 'ok') {
      dispatch(skipQueueItemSuccess());
    } else {
      dispatch(skipQueueItemFailure(json));
    }
  }).catch((err) => {
    dispatch(skipQueueItemFailure(err));
  });
};
