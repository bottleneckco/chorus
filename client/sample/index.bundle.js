const socket = new WebSocket("ws://localhost:8080/stream", "GET");

window.AudioContext = window.AudioContext || window.webkitAudioContext;
const context = new AudioContext();

const convertToArrayBuffer = (data) => {
  return new Promise((resolve) => {
    const fileReader = new FileReader();
    fileReader.onload = function () {
      resolve(this.result);
    };
    fileReader.readAsArrayBuffer(data);
  });
}

socket.onmessage = (event) => {
  const source = context.createBufferSource();

  convertToArrayBuffer(event.data).then((arrayBuffer) => {
    context.decodeAudioData(arrayBuffer, (buffer) => {
      source.buffer = buffer;
      source.connect(context.destination);
      source.start(context.currentTime);
    });
  });
};

socket.onopen = (event) => {
  socket.send('Hello, this is JS!');
};

