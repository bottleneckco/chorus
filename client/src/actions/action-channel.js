import 'whatwg-fetch';

import * as types from '../constants/action-types';
import API_ROOT from '../constants/api-url';

const createChannelRequest = () => ({
  type: types.CREATE_CHANNEL_REQUEST
});

const createChannelSuccess = () => ({
  type: types.CREATE_CHANNEL_SUCCESS
});

const createChannelFailure = () => ({
  type: types.CREATE_CHANNEL_FAILURE
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
    .then((response) => {
      dispatch(createChannelSuccess());
      console.log(response);
    })
    .catch((error) => {
      dispatch(createChannelFailure());
      console.log(error);
    });
};
