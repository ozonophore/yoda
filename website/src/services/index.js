import { parseISO } from "date-fns";
import { NextRun } from "../context/actions";

const createWSService = (dispatch) => {
  const ws = new WebSocket("ws://localhost:88/ws");

  ws.onmessage = (evt) => {
    const message = JSON.parse(evt.data);
    dispatch(NextRun(parseISO(message.nextRun)));
  };
};

export default createWSService;
