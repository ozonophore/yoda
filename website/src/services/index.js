import { parseISO } from "date-fns";
import { NextRun } from "../context/actions";

function ping(ws) {
  console.log("#Ping");
  ws.send("__ping__");
}

let ws = null;

const createWSService = (dispatch) => {
  if (ws) {
    console.log("#Already Connected");
    return;
  }
  console.log("#New Connecting");
  let interval = null;
  const host = window.location.hostname.match("([\\w\\.]+)*(:[0-9]+)?")[1];
  ws = new WebSocket(`ws://${host}:88/ws`);

  ws.onopen = () => {
    console.log("#Connected");
    interval = setInterval(() => {
      ping(ws);
    }, 30000);
  };

  ws.onmessage = (evt) => {
    if (evt.data === "__pong__") {
      console.log("#Pong");
      return;
    }
    console.log("#Received: ", evt.data);
    const message = JSON.parse(evt.data);
    dispatch(NextRun(parseISO(message.nextRun)));
  };

  ws.onclose = () => {
    console.log("#Disconnected");
    clearInterval(interval);
  };
};

export default createWSService;
