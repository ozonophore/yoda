import { parse } from "date-fns";
import { DefaultService } from "../../generated";
import { setDate } from "../index";

const RefreshRooms = () => (dispatch) => {
  dispatch({ type: "LOADING", value: true });
  DefaultService.getRooms()
    .then((data) => {
      dispatch({ type: "REFRESH_ROOMS", value: data });
    })
    .catch((error) => {
      dispatch({ type: "SHOW_ERROR", value: error.statusText });
    });
};

const RefreshOrganisations = () => (dispatch) => {
  dispatch({ type: "LOADING", value: true });
  DefaultService.getOrganisations()
    .then((data) => {
      dispatch({ type: "REFRESH_ORGANISATIONS", value: data });
    })
    .catch((error) => {
      dispatch({ type: "SHOW_ERROR", value: error.statusText });
    });
};

const CloseError = () => (dispatch) => {
  dispatch({ type: "CLOSE_ERROR", value: false });
};

const CreateRoom = (room) => (dispatch) => {
  dispatch({ type: "LOADING", value: true });
  DefaultService.createRoom(room)
    .then(() => {
      RefreshRooms()(dispatch);
    })
    .catch((error) => {
      dispatch({ type: "SHOW_ERROR", value: error.statusText });
    });
};

const RefreshJobs = () => (dispatch) => {
  dispatch({ type: "LOADING", value: true });
  DefaultService.getJobs()
    .then((data) => {
      dispatch({ type: "REFRESH_JOBS", value: data });
    })
    .catch((error) => {
      dispatch({ type: "SHOW_ERROR", value: error.statusText });
    });
};

const RoomAddToggle = (value) => (dispatch) => {
  dispatch({ type: "ROOM_ADD_TOGGLE", value });
};

const RoomGridToggle = (value) => (dispatch) => {
  dispatch({ type: "ROOM_GRID_TOGGLE", value });
};

const Ping = () => (dispatch) => {
  DefaultService.ping().then((data) => {
    const currentDate = parse(data.date, "yyyy-MM-dd HH:mm:ss", new Date());
    setDate(dispatch, currentDate);
  });
};

const NextRun = (value) => (dispatch) => {
  dispatch({ type: "NEXT_RUN", value });
};

export {
  RefreshRooms,
  CloseError,
  CreateRoom,
  RefreshJobs,
  RoomAddToggle,
  RoomGridToggle,
  Ping,
  NextRun,
  RefreshOrganisations,
};
