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

    this.state = {
      playing: true,
      showSearchBar: true
    };

    this.initWS = this.initWS.bind(this);
    this.pause = this.pause.bind(this);
    this.resume = this.resume.bind(this);
    this.skip = this.skip.bind(this);
    this.adjustAudioVol = this.adjustAudioVol.bind(this);
    this.skipQueueItem = this.skipQueueItem.bind(this);
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

    const queue = [];

    mediaSource.addEventListener('sourceopen', () => {
      const sourceBuffer = mediaSource.addSourceBuffer('audio/mpeg');
      sourceBuffer.mode = 'sequence';

      setInterval(() => {
        if (queue.length === 0) {
          return;
        }
        const data = queue.pop();
        convertToArrayBuffer(data).then((arrayBuffer) => {
          try {
            sourceBuffer.appendBuffer(arrayBuffer);
          } catch (err) {
            console.log(err);
          }
        });
      }, 600);

      this.socket.onmessage = (event) => {
        if (typeof event.data !== 'string') {
          queue.unshift(event.data);
        } else {
          const wsJson = JSON.parse(event.data);
          if (wsJson.command !== 'ping') {
            console.log(`Received ${wsJson.command} from server`);
          }
          
          switch (wsJson.command) {
            case 'pause':
              audio.pause();
              this.setState({ playing: false });
              break;
            case 'resume':
              audio.play();
              this.setState({ playing: true });
              break;
            case 'skipCurrent':
              sourceBuffer.remove(sourceBuffer.buffered.start(0), sourceBuffer.buffered.end(0) + 20);
              break;
            case 'updateQueue':
              // console.log("FUCK")
              this.props.fetchQueue();
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

  skip() {
    this.socket.send('skipCurrent');
  }

  adjustAudioVol(newVol) {
    this.audio.volume = newVol / 100.0;
  }

  skipQueueItem(index) {
    this.props.skipQueueItem(this.props.data.id, index);
  }

  render() {
    const currentSong = this.props.queue[0] || {};

    return (
      <div className="channel">
        <Nav title={this.props.data.name} />
        <Player
          pause={this.pause}
          resume={this.resume}
          skip={this.skip}
          playing={this.state.playing}
          currentSong={currentSong}
          adjustAudioVol={this.adjustAudioVol}
        />
        <Queue queue={this.props.queue} skipQueueItem={this.skipQueueItem} />
      </div>
    );
  }
}

WSPlayer.propTypes = {
  data: PropTypes.object.isRequired,
  queue: PropTypes.array.isRequired,
  skipQueueItem: PropTypes.func.isRequired,
  fetchQueue: PropTypes.func.isRequired
};

export default WSPlayer;
