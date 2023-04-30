import { parseISO } from "date-fns";
import { NextRun } from "../context/actions";

function ping(ws) {
  ws.send("__ping__");
}

let ws = null;

const createWSService = (dispatch) => {
  if (ws) {
    return;
  }
  let interval = null;
  const host = window.location.hostname.match("([\\w\\.]+)*(:[0-9]+)?")[1];
  ws = new WebSocket(`ws://${host}:88/ws`);

  ws.onopen = () => {
    interval = setInterval(() => {
      ping(ws);
    }, 30000);
  };

  ws.onmessage = (evt) => {
    if (evt.data === "__pong__") {
      return;
    }
    const message = JSON.parse(evt.data);
    dispatch(NextRun(parseISO(message.nextRun)));
  };

  ws.onclose = () => {
    clearInterval(interval);
  };
};

export default createWSService;
