import { useEffect, useState } from "react";
import Grid from "@mui/material/Grid";
import MDSnackbar from "../../components/MDSnackbar";
import { RefreshRooms, CloseError, CreateRoom } from "../../context/actions";
import DashboardLayout from "../../examples/LayoutContainers/DashboardLayout";
import DashboardNavbar from "../../examples/Navbars/DashboardNavbar";
import MDBox from "../../components/MDBox";
import { useMaterialUIController } from "../../context";
import RoomInfoCard from "../../examples/Cards/InfoCards/RoomInfoCard";
import Header from "./components/Header";
import RoomCard from "./components/RoomCard";

function Room() {
  const [isNewRoom, setIsNewRoom] = useState(false);
  const [editKey, setEditKey] = useState(false);
  const [controller, dispatch] = useMaterialUIController();
  const { rooms, error } = controller;
  const handleCloseError = () => dispatch(CloseError());

  const renderError = (
    <MDSnackbar
      color="error"
      icon="warning"
      title="Room"
      content={error?.message}
      dateTime="11 mins ago"
      open={Boolean(error?.message)}
      onClose={handleCloseError}
      close={handleCloseError}
      bgWhite
    />
  );

  useEffect(() => {
    dispatch(RefreshRooms());
  }, []);
  const handleOnSubmit = (room) => {
    dispatch(CreateRoom(room));
    setEditKey(null);
  };
  const handleOnNewRoom = (e) => {
    setIsNewRoom(true);
    setEditKey(null);
    e.preventDefault();
  };
  const handleOnCancel = (e) => {
    setEditKey(null);
    setIsNewRoom(false);
    e.preventDefault();
  };
  const handleOnEdit = (room) => {
    setEditKey(room.code);
    setIsNewRoom(false);
  };
  const handleOnRefresh = (e) => {
    dispatch(RefreshRooms());
    e.preventDefault();
  };
  return (
    <DashboardLayout>
      <DashboardNavbar onRefresh={handleOnRefresh} />
      <Header isShow={!isNewRoom} onClick={handleOnNewRoom} />
      {renderError}
      {Boolean(isNewRoom) && <RoomCard onCancel={handleOnCancel} onSubmit={handleOnSubmit} />}
      <MDBox mt={4.5}>
        <Grid container spacing={2}>
          {rooms.map((room) =>
            room.code === editKey ? (
              <Grid key={room.code} item xs={12} md={12} lg={12}>
                <MDBox mb={3}>
                  <RoomCard onSubmit={handleOnSubmit} onCancel={handleOnCancel} room={room} />
                </MDBox>
              </Grid>
            ) : (
              <Grid key={room.code} item xs={12} md={6} lg={6}>
                <MDBox mb={3}>
                  <RoomInfoCard
                    color="success"
                    icon="tv"
                    title={room.code}
                    name={room.name}
                    days={room.days}
                    time={room.times}
                    onEdit={() => handleOnEdit(room)}
                    percentage={{
                      color: "success",
                      amount: "+1%",
                      label: "than yesterday",
                    }}
                  />
                </MDBox>
              </Grid>
            )
          )}
        </Grid>
      </MDBox>
    </DashboardLayout>
  );
}

export default Room;
