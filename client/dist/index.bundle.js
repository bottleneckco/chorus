const socket = new WebSocket("ws://localhost:8080/stream", "GET");

socket.onmessage = (event) => {
  console.log(event.data);
};

socket.onopen = (event) => {
  socket.send('Hello, this is JS!');
};

