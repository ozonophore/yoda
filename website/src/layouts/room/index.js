import { useEffect } from "react";
import MDBox from "../../components/MDBox";
import MDSnackbar from "../../components/MDSnackbar";
import { useMaterialUIController } from "../../context";
import { CloseError, CreateRoom, RefreshOrganisations, RefreshRooms } from "../../context/actions";
import Footer from "../../examples/Footer";
import DashboardLayout from "../../examples/LayoutContainers/DashboardLayout";
import DashboardNavbar from "../../examples/Navbars/DashboardNavbar";
import Header from "./components/Header";
import RoomGrid from "./components/RoomGrid";

function Room() {
  const [controller, dispatch] = useMaterialUIController();
  const { error } = controller;
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
    dispatch(RefreshOrganisations());
  }, []);
  const handleOnSubmit = (room) => {
    dispatch(CreateRoom(room));
  };
  const handleOnRefresh = (e) => {
    dispatch(RefreshRooms());
    e.preventDefault();
  };

  return (
    <DashboardLayout>
      <DashboardNavbar onRefresh={handleOnRefresh} />
      <Header onSubmit={handleOnSubmit} />
      {renderError}
      <MDBox>
        <RoomGrid />
      </MDBox>
      <Footer />
    </DashboardLayout>
  );
}

export default Room;
