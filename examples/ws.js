// Basic WS example code for client
var ws = new WebSocket("ws://localhost:8000/api/ws");

ws.onclose = ({ wasClean, code, reason }) => {
  console.log(`onclose:   ${JSON.stringify({ wasClean, code, reason })}`);
};

ws.onerror = (error) => {
  console.log(error);
  console.log("onerror:   An error has occurred. See console for details.");
};

ws.onmessage = ({ data }) => {
  console.log(`onmessage: ${data}`);
};

ws.onopen = () => {
  console.log("onopen:    Connected successfully.");
  // We don't accept anything from the frontend side but here is an example of how it could be done.
  // ws.send("hi from the client");
};
