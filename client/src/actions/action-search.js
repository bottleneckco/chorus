import * as types from '../constants/action-types';
import API_ROOT from '../constants/api-url';

const searchMusicRequest = () => ({
  type: types.SEARCH_MUSIC_REQUEST
});

const searchMusicSuccess = (data) => ({
  data,
  type: types.SEARCH_MUSIC_SUCCESS
});

const searchMusicFailure = (error) => ({
  error,
  type: types.SEARCH_MUSIC_FAILURE
});

// searchMusic must be called before adding to queue
export const searchMusic = (channelId, term) => (dispatch) => {
  dispatch(searchMusicRequest());
  return fetch(`${API_ROOT}/api/channels/${channelId}/search?term=${term}`, {
    method: 'GET',
    credentials: 'same-origin'
  }).then((response) => (
    response.json()
  )).then((json) => {
    if (json.status === 'ok') {
      dispatch(searchMusicSuccess(json.results));
    } else {
      dispatch(searchMusicFailure(JSON.stringify(json)));
    }
  }).catch((err) => {
    dispatch(searchMusicFailure(err));
  });
};
