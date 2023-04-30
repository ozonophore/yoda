// prop-types is a library for typechecking of props
// @mui material components
import Card from "@mui/material/Card";
import Divider from "@mui/material/Divider";
import Icon from "@mui/material/Icon";
import Tooltip from "@mui/material/Tooltip";

// Material Dashboard 2 React components
import MDBox from "components/MDBox";
import MDTypography from "components/MDTypography";
import PropTypes from "prop-types";
import { useTranslation } from "react-i18next";

function RoomInfoCard({ color, title, name, organization, days, time, icon, onEdit }) {
  const [t] = useTranslation();

  const dayOfWeekLabel = days.map((item) => t(`dayOfWeek.short.${item}`)).join("; ");
  return (
    <Card>
      <MDBox display="flex" justifyContent="space-between" pt={1} px={2}>
        <MDBox
          variant="gradient"
          bgColor={color}
          color={color === "light" ? "dark" : "white"}
          coloredShadow={color}
          borderRadius="xl"
          display="flex"
          justifyContent="center"
          alignItems="center"
          width="4rem"
          height="4rem"
          mt={-3}
        >
          <Icon fontSize="medium" color="inherit">
            {icon}
          </Icon>
        </MDBox>
        <MDBox textAlign="right" lineHeight={1.25}>
          <MDTypography variant="button" fontWeight="light" color="text">
            {title}&nbsp;
            <MDBox component="span" ml="auto" lineHeight={0} color="info" onClick={onEdit}>
              <Tooltip title={t("room.card.edit")} placement="top">
                <Icon sx={{ cursor: "pointer" }} fontSize="small">
                  edit
                </Icon>
              </Tooltip>
            </MDBox>
            <MDBox component="span" ml="auto" pl={1} lineHeight={0} color="red" onClick={onEdit}>
              <Tooltip title={t("room.card.delete")} placement="top">
                <Icon sx={{ cursor: "pointer" }} fontSize="small">
                  delete
                </Icon>
              </Tooltip>
            </MDBox>
          </MDTypography>
          <MDTypography variant="h4">{name}</MDTypography>
        </MDBox>
      </MDBox>
      <MDBox px={2}>
        {organization && <MDTypography variant="button">{organization}</MDTypography>}
        {!organization && (
          <MDTypography variant="button" color="warning">
            {t("room.orgUndefined")}
          </MDTypography>
        )}
      </MDBox>
      <Divider />
      <MDBox pb={2} px={2}>
        <MDTypography component="p" variant="button" color="text" display="flex">
          {t("dayOfWeek.title.short")}:&nbsp;{dayOfWeekLabel}
        </MDTypography>
        <MDTypography component="p" variant="button" color="text" display="flex">
          {t("time.title.short")}:&nbsp;{time.join("; ")}
        </MDTypography>
      </MDBox>
    </Card>
  );
}

// Setting default values for the props of ComplexStatisticsCard
RoomInfoCard.defaultProps = {
  color: "info",
  days: [],
  time: [],
  organization: null,
  percentage: {
    color: "success",
    text: "",
    label: "",
  },
};

// Typechecking props for the ComplexStatisticsCard
RoomInfoCard.propTypes = {
  color: PropTypes.oneOf([
    "primary",
    "secondary",
    "info",
    "success",
    "warning",
    "error",
    "light",
    "dark",
  ]),
  title: PropTypes.string.isRequired,
  days: PropTypes.arrayOf(PropTypes.string),
  time: PropTypes.arrayOf(PropTypes.string),
  name: PropTypes.oneOfType([PropTypes.string]).isRequired,
  organization: PropTypes.oneOfType([PropTypes.string]),
  onEdit: PropTypes.func.isRequired,
  percentage: PropTypes.shape({
    color: PropTypes.oneOf([
      "primary",
      "secondary",
      "info",
      "success",
      "warning",
      "error",
      "dark",
      "white",
    ]),
    amount: PropTypes.oneOfType([PropTypes.string, PropTypes.number]),
    label: PropTypes.string,
  }),
  icon: PropTypes.node.isRequired,
};

export default RoomInfoCard;
