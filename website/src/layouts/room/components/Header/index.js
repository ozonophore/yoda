import Grid from "@mui/material/Grid";
import Icon from "@mui/material/Icon";
import PropTypes from "prop-types";
import { useTranslation } from "react-i18next";
import MDBox from "../../../../components/MDBox";
import MDButton from "../../../../components/MDButton";
import { useMaterialUIController } from "../../../../context";
import { RoomAddToggle } from "../../../../context/actions";
import RoomCard from "../RoomCard";

function RoomHeader({ onSubmit }) {
  const [controller, dispatch] = useMaterialUIController();
  const { addToggle } = controller.room;
  const [t] = useTranslation();
  const handleOnCancel = () => {
    dispatch(RoomAddToggle(false));
  };

  const handleOnClicked = () => {
    dispatch(RoomAddToggle(true));
  };

  return (
    <MDBox pt={0} px={0} pb={0} display="flex" justifyContent="end" alignItems="center">
      {!addToggle && (
        <MDButton variant="gradient" color="info" onClick={handleOnClicked}>
          <Icon sx={{ fontWeight: "bold" }}>add</Icon>
          &nbsp;{t("room.button.add")}
        </MDButton>
      )}
      {addToggle && (
        <Grid container alignItems="center" justifyContent="center">
          <Grid item xs={12} md={8} lg={8}>
            <RoomCard onCancel={handleOnCancel} onSubmit={onSubmit} />
          </Grid>
        </Grid>
      )}
    </MDBox>
  );
}

RoomHeader.propTypes = {
  onSubmit: PropTypes.func.isRequired,
};

export default RoomHeader;
