import React, { Component } from 'react';
import { WS_ROOT } from '../constants/api-url';

const convertToArrayBuffer = (data) => (
  new Promise((resolve) => {
    const fileReader = new FileReader();
    fileReader.onload = function () {
      resolve(this.result);
    };
    fileReader.readAsArrayBuffer(data);
  })
);

class WsTest extends Component {
  constructor(props) {
    super(props);

    const channelId = '7Y7LrlDA5gbEL';
    this.socket = new WebSocket(`${WS_ROOT}/api/channels/${channelId}/stream`, 'GET');

    this.pause = this.pause.bind(this);
    this.resume = this.resume.bind(this);
  }

  componentWillMount() {
    document.cookie = 'user_id=1; path=/;'; // Fake cookie
    document.cookie = 'user_nickname=Harry; path=/;'; // Fake cookie
    console.log("running component will mount");

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
    
    // play audio on beginning
    audio.src = URL.createObjectURL(mediaSource);
    audio.play();
  }

  pause() {
    this.socket.send('pause');
  }

  resume() {
    this.socket.send('resume');
  }
  
  render() {
    return (
      <div>
        wsTest
        <button onClick={this.pause}>Pause</button>
        <button onClick={this.resume}>Resume</button>
      </div>
    );
  }
}

export default WsTest;
