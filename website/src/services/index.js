import { parseISO } from "date-fns";
import { NextRun } from "../context/actions";

const createWSService = (dispatch) => {
  console.log("#location: ", window.location);
  const host = window.location.hostname.match("([\\w\\.]+)*(:[0-9]+)?")[1];
  console.log("#host: ", host);
  const ws = new WebSocket(`ws://${host}:88/ws`);

  ws.onmessage = (evt) => {
    console.log("#Received: ", evt.data);
    const message = JSON.parse(evt.data);
    dispatch(NextRun(parseISO(message.nextRun)));
  };
};

export default createWSService;
