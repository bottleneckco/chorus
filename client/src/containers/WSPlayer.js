import React, { Component } from 'react';
import PropTypes from 'prop-types';

import { WS_ROOT } from '../constants/api-url';
import Nav from '../components/Nav';
import Queue from '../components/Queue';
import Player from '../components/Player';

const convertToArrayBuffer = (data) => (
  new Promise((resolve) => {
    const fileReader = new FileReader();
    fileReader.onload = function () {
      resolve(this.result);
    };
    fileReader.readAsArrayBuffer(data);
  })
);


class WSPlayer extends Component {
  constructor(props) {
    super(props);

    this.state = { playing: true };

    this.initWS = this.initWS.bind(this);
    this.pause = this.pause.bind(this);
    this.resume = this.resume.bind(this);
    this.adjustAudioVol = this.adjustAudioVol.bind(this);
  }

  componentWillMount() {
    this.initWS();
  }

  initWS() {
    const channelId = this.props.data.id;
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
              this.setState({ playing: false });
              console.log('Received pause from server');
              break;
            case 'resume':
              audio.play();
              this.setState({ playing: true });
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
    audio.volume = 0.8;
    audio.play();

    this.audio = audio;
  }

  pause() {
    this.socket.send('pause');
  }

  resume() {
    this.socket.send('resume');
  }

  adjustAudioVol(newVol) {
    this.audio.volume = newVol / 100.0;
  }

  render() {
    return (
      <div className="channel">
        <Nav title={this.props.data.name} />
        <Player
          pause={this.pause}
          resume={this.resume}
          playing={this.state.playing}
          adjustAudioVol={this.adjustAudioVol}
        />
        <Queue queue={this.props.queue} />
      </div>
    );
  }
}

WSPlayer.propTypes = {
  data: PropTypes.object.isRequired,
  queue: PropTypes.array.isRequired
};

export default WSPlayer;
