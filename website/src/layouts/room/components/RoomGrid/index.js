import Grid from "@mui/material/Grid";
import { useState } from "react";
import MDBox from "../../../../components/MDBox";
import { useMaterialUIController } from "../../../../context";
import { CreateRoom, RoomGridToggle } from "../../../../context/actions";
import RoomInfoCard from "../../../../examples/Cards/InfoCards/RoomInfoCard";
import RoomCard from "../RoomCard";

function RoomGrid() {
  const [controller, dispatch] = useMaterialUIController();
  const { rooms } = controller;
  const { gridToggle } = controller.room;

  const [selectedRoom, setSelectedRoom] = useState(null);

  const handleOnCancel = () => {
    setSelectedRoom(null);
    dispatch(RoomGridToggle(false));
  };
  const habdleOnEdit = (room) => {
    setSelectedRoom(room);
    dispatch(RoomGridToggle(true));
  };
  const handleOnSubmit = (room) => {
    dispatch(CreateRoom(room));
    dispatch(RoomGridToggle(false));
  };
  return (
    <MDBox mt={4.5}>
      <Grid container spacing={2}>
        {rooms.map((room) =>
          gridToggle && selectedRoom.code === room.code ? (
            <Grid key={room.code} item xs={8} md={8} lg={8}>
              <MDBox mb={3}>
                <RoomCard room={selectedRoom} onCancel={handleOnCancel} onSubmit={handleOnSubmit} />
              </MDBox>
            </Grid>
          ) : (
            <Grid key={room.code} item xs={12} md={4} lg={4} onClick={() => habdleOnEdit(room)}>
              <MDBox mb={3}>
                <RoomInfoCard
                  color="success"
                  icon="tv"
                  title={room.code}
                  name={room.name}
                  days={room.days}
                  time={room.times}
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
  );
}

RoomCard.defaultProps = {};

RoomCard.propTypes = {};

export default RoomGrid;
