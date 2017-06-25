import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';

import { WS_ROOT } from '../constants/api-url';
import { getRehydrationStatus } from '../reducers/reducer-persistence';
import { fetchChannel } from '../actions/action-channel';
import { fetchQueue } from '../actions/action-queue';
import { getData, getChannelIsFetching } from '../reducers/reducer-channel';
import { getQueue, getQueueIsFetching } from '../reducers/reducer-queue';

import Nav from '../components/Nav';
import Player from '../components/Player';
import Queue from '../components/Queue';
import '../stylesheets/channel.scss';

const convertToArrayBuffer = (data) => (
    new Promise((resolve) => {
      const fileReader = new FileReader();
      fileReader.onload = function () {
        resolve(this.result);
      };
      fileReader.readAsArrayBuffer(data);
    })
  );

class Channel extends Component {
  constructor(props) {
    super(props);

    this.state = { rehydrated: false };

    this.initWS = this.initWS.bind(this);
    this.pause = this.pause.bind(this);
    this.resume = this.resume.bind(this);
  }

  componentWillMount() {
    if (this.props.rehydrated) {
      this.initWS();
      this.setState({ rehydrated: true });
      console.log('compoinentwillmount');
    }
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.rehydrated && !this.state.rehydrated) {
      this.initWS();
      this.setState({ rehydrated: true });
      console.log('componentWillreceiveipoprs');
    }
  }

  initWS() {
    document.cookie = 'user_id=1; path=/;'; // Fake cookie
    document.cookie = 'user_nickname=Harry; path=/;'; // Fake cookie

    const channelId = this.props.response.id;
    this.socket = new WebSocket(`${WS_ROOT}/api/channels/${channelId}/stream`, 'GET');
    console.log('Initializing WS connection');

    const audio = document.createElement('audio');
    const mediaSource = new MediaSource();

    mediaSource.addEventListener('sourceopen', () => {
      const sourceBuffer = mediaSource.addSourceBuffer('audio/mpeg');

      this.socket.onmessage = (event) => {
        if (typeof event.data !== 'string') {
          convertToArrayBuffer(event.data).then((arrayBuffer) => {
            sourceBuffer.appendBuffer(arrayBuffer);
          });
        } else {
          const wsJson = JSON.parse(event.data);
          switch (wsJson.command) {
            case 'pause':
              audio.pause();
              console.log('Received pause from server');
              break;
            case 'resume':
              audio.play();
              console.log('Received resume from server');
              break;
            // case 'ping':
            //   console.log('Received ping from server');
            //   break;
            default:
              break;
          }
        }
      };
    });

    // play audio once mounted
    audio.src = URL.createObjectURL(mediaSource);
    audio.play();
  }

  pause() {
    this.socket.send('pause');
  }

  resume() {
    this.socket.send('resume');
  }

  componentWillMount() {
    const { data } = this.props;
    if (data.status === 'ok') {
      this.props.fetchQueue(this.props.data.channel.id);
    } else {
      this.props.fetchChannel(this.props.match.params.hash);
    }
  }

  render() {
    const { data, channelIsFetching, queue } = this.props;

    if (channelIsFetching || data.status !== 'ok') {
      return <div>Loading...</div>;
    }

    return (
      <div className="channel">
        <Nav title={this.props.response.name} />
        <Player pause={this.pause} resume={this.resume} />
        <Queue queue={this.props.queue} />
      </div>
    );
  }
}

Channel.defaultProps = {
  data: {}
};

Channel.propTypes = {
  match: PropTypes.object.isRequired,

  data: PropTypes.object,
  channelIsFetching: PropTypes.bool.isRequired,
  queue: PropTypes.array.isRequired,
  queueIsFetching: PropTypes.bool.isRequired,

  fetchChannel: PropTypes.func.isRequired,
  fetchQueue: PropTypes.func.isRequired
};

const mapStateToProps = (state) => ({
  response: getResponse(state),
  rehydrated: getRehydrationStatus(state),
  data: getData(state),
  channelIsFetching: getChannelIsFetching(state),
  queue: getQueue(state),
  queueIsFetching: getQueueIsFetching(state)
});

const mapDispatchToProps = (dispatch) => (
  bindActionCreators({
    fetchChannel,
    fetchQueue
  }, dispatch)
);

export default connect(mapStateToProps, mapDispatchToProps)(Channel);
