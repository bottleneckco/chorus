import 'whatwg-fetch';

import * as types from '../constants/action-types';
import API_ROOT from '../constants/api-url';

const createChannelRequest = () => ({
  type: types.CREATE_CHANNEL_REQUEST
});

const createChannelSuccess = (res) => ({
  type: types.CREATE_CHANNEL_SUCCESS,
  response: res
});

const createChannelFailure = (error) => ({
  type: types.CREATE_CHANNEL_FAILURE,
  error
});

export const createChannel = (data) => (dispatch) => {
  dispatch(createChannelRequest());
  return fetch(`${API_ROOT}/api/channels`, {
    method: 'POST',
    body: JSON.stringify(data),
    headers: {
      'Content-Type': 'application/json'
    }
  })
    .then((res) => (
      res.json()
    ))
    .then((res) => {
      dispatch(createChannelSuccess(res));
    })
    .catch((err) => {
      dispatch(createChannelFailure(err));
    });
};
