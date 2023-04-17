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
export { RefreshRooms, CloseError, CreateRoom, RefreshJobs, RoomAddToggle, RoomGridToggle };
