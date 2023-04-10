import Card from "@mui/material/Card";
import Divider from "@mui/material/Divider";
import { format } from "date-fns";
import PropTypes from "prop-types";
import { useTranslation } from "react-i18next";
import useDeepCompareEffect from "use-deep-compare-effect";
import MDBox from "../../../../components/MDBox";
import MDTypography from "../../../../components/MDTypography";
import { useMaterialUIController } from "../../../../context";
import { RefreshJobs } from "../../../../context/actions";

function RegularCard({ title }) {
  const [t] = useTranslation();
  const [controller, dispatch] = useMaterialUIController();
  const { weekDays, atTimes, nextRun, lastRun } = controller.jobs.loader;
  const dayOfWeekLabel = weekDays.map((item) => t(`dayOfWeek.short.${item}`)).join("; ");

  useDeepCompareEffect(() => {
    dispatch(RefreshJobs());
  }, [controller.jobs]);

  const renderAtTimes = (label, runDate) => (
    <MDBox pb={2} px={2} display="flex" flexDirection="row">
      <MDTypography
        component="span"
        variant="button"
        fontWeight="bold"
        display="flex"
        flexDirection="column"
      >
        {label}:
      </MDTypography>
      <MDTypography
        component="span"
        variant="button"
        color="text"
        display="flex"
        flexDirection="column"
      >
        &nbsp;{runDate ? format(runDate, "dd.MM.yyyy HH:mm") : ""}
      </MDTypography>
    </MDBox>
  );

  return (
    <Card>
      <MDBox pt={3} px={2}>
        <MDTypography variant="h6" fontWeight="medium" textTransform="uppercase">
          {title}
        </MDTypography>
      </MDBox>
      <MDBox pt={1} pb={0} px={2}>
        <MDBox component="ul" display="flex" flexDirection="column" p={0} m={0}>
          {renderAtTimes("Дата следующего запуска", nextRun)}
          {renderAtTimes("Дата последнего запуска", lastRun)}
          <Divider />
          <MDBox pb={2} px={2}>
            <MDTypography component="p" variant="button" color="text" display="flex">
              {t("dayOfWeek.title.long")}:&nbsp;{dayOfWeekLabel}
            </MDTypography>
            <MDTypography component="p" variant="button" color="text" display="flex">
              {t("time.title.long")}:&nbsp;{atTimes || atTimes.join("; ")}
            </MDTypography>
          </MDBox>
        </MDBox>
      </MDBox>
    </Card>
  );
}

RegularCard.propTypes = {
  title: PropTypes.string.isRequired,
};

export default RegularCard;
