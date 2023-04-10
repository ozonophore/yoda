import { DefaultService } from "../../generated";

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
  console.log("#RefreshJobs");
  dispatch({ type: "LOADING", value: true });
  DefaultService.getJobs()
    .then((data) => {
      dispatch({ type: "REFRESH_JOBS", value: data });
    })
    .catch((error) => {
      dispatch({ type: "SHOW_ERROR", value: error.statusText });
    });
};

export { RefreshRooms, CloseError, CreateRoom, RefreshJobs };
