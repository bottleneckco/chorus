const socket = new WebSocket("ws://localhost:8080/stream", "GET");

window.AudioContext = window.AudioContext || window.webkitAudioContext;
const context = new AudioContext();

socket.onmessage = (event) => {
  const source = context.createBufferSource();

  const fileReader = new FileReader();
  fileReader.onload = function () {
    context.decodeAudioData(this.result, (buffer) => {
      source.buffer = buffer;
      source.connect(context.destination);
      source.start(context.currentTime);
    });
  };
  fileReader.readAsArrayBuffer(event.data);
};

socket.onopen = (event) => {
  socket.send('Hello, this is JS!');
};

