import Icon from "@mui/material/Icon";
import PropTypes from "prop-types";
import { useTranslation } from "react-i18next";
import MDBox from "../../../../components/MDBox";
import MDButton from "../../../../components/MDButton";

function RoomHeader({ isShow, onClick }) {
  const [t] = useTranslation();
  return (
    <MDBox pt={2} px={2} pb={2} display="flex" justifyContent="end" alignItems="center">
      {Boolean(isShow) && (
        <MDButton variant="gradient" color="info" onClick={onClick}>
          <Icon sx={{ fontWeight: "bold" }}>add</Icon>
          &nbsp;{t("room.button.add")}
        </MDButton>
      )}
    </MDBox>
  );
}

RoomHeader.propTypes = {
  isShow: PropTypes.bool.isRequired,
  onClick: PropTypes.func.isRequired,
};

export default RoomHeader;
