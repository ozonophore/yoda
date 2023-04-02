import { useTranslation } from "react-i18next";
import { useState } from "react";
import Icon from "@mui/material/Icon";
import Grid from "@mui/material/Grid";
import RoomEditCard from "../../examples/Cards/InfoCards/RoomEditCard";
import RoomNewCard from "../../examples/Cards/InfoCards/RoomNewCard";
import DashboardLayout from "../../examples/LayoutContainers/DashboardLayout";
import DashboardNavbar from "../../examples/Navbars/DashboardNavbar";
import MDBox from "../../components/MDBox";
import MDButton from "../../components/MDButton";
import { createNewRoom, useMaterialUIController } from "../../context";
import RoomInfoCard from "../../examples/Cards/InfoCards/RoomInfoCard";

function Room() {
  const [isNewRoom, setIsNewRoom] = useState(false);
  const [editRoom, setEditRoom] = useState(false);
  const [t] = useTranslation();
  const [controller, dispatch] = useMaterialUIController();
  const { rooms } = controller;
  const handleOnSubmit = (e) => {
    createNewRoom(dispatch);
    e.preventDefault();
  };
  const handleOnNewRoom = (e) => {
    setIsNewRoom(true);
    setEditRoom(null);
    e.preventDefault();
  };
  const handleOnCancel = (e) => {
    setIsNewRoom(false);
    setEditRoom(null);
    e.preventDefault();
  };
  const handleOnEdit = (room) => {
    setIsNewRoom(false);
    setEditRoom(room);
  };
  return (
    <DashboardLayout>
      <DashboardNavbar />
      <MDBox pt={2} px={2} pb={4} display="flex" justifyContent="end" alignItems="center">
        {Boolean(isNewRoom) === false && Boolean(editRoom) === false && (
          <MDButton variant="gradient" color="info" onClick={handleOnNewRoom}>
            <Icon sx={{ fontWeight: "bold" }}>add</Icon>
            &nbsp;{t("room.button.add")}
          </MDButton>
        )}
      </MDBox>
      {Boolean(isNewRoom) && <RoomNewCard onCancel={handleOnCancel} onSubmit={handleOnSubmit} />}
      {Boolean(editRoom) && (
        <RoomEditCard
          code={editRoom.code}
          name={editRoom.name}
          ozon={editRoom.ozon}
          wb={editRoom.wb}
          onCancel={handleOnCancel}
        />
      )}
      <MDBox mt={4.5}>
        <Grid container spacing={2}>
          {rooms.map((room) => (
            <Grid key={room.code} item xs={12} md={6} lg={6}>
              <MDBox mb={3}>
                <RoomInfoCard
                  color="success"
                  icon="tv"
                  title={room.code}
                  name={room.name}
                  dayOfWeek={room.dayOfWeek}
                  time={room.time}
                  onEdit={() => handleOnEdit(room)}
                  percentage={{
                    color: "success",
                    amount: "+1%",
                    label: "than yesterday",
                  }}
                />
              </MDBox>
            </Grid>
          ))}
        </Grid>
      </MDBox>
    </DashboardLayout>
  );
}

export default Room;
