import { parseISO } from "date-fns";
import { NextRun } from "../context/actions";

const createWSService = (dispatch) => {
  console.log("#location: ", window.location);
  const host = window.location.hostname.match("([\\w\\.]+)*(:[0-9]+)?")[1];
  console.log("#host: ", host);
  const ws = new WebSocket(`ws://${host}:88/ws`);

  ws.onopen = () => {
    console.log("#Connected");
  };

  ws.onmessage = (evt) => {
    console.log("#Received: ", evt.data);
    const message = JSON.parse(evt.data);
    dispatch(NextRun(parseISO(message.nextRun)));
  };

  ws.onclose = () => {
    console.log("#Disconnected");
  };
};

export default createWSService;
