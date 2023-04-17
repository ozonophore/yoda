import Card from "@mui/material/Card";
import Divider from "@mui/material/Divider";
import { format, formatDistanceStrict } from "date-fns";
import { ru } from "date-fns/locale";
import PropTypes from "prop-types";
import { useTranslation } from "react-i18next";
import MDBox from "../../../../components/MDBox";
import MDTypography from "../../../../components/MDTypography";

function JobCard({ job }) {
  const [t] = useTranslation();

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
  const { name, description } = job;
  const { weekDays, atTimes, nextRun, lastRun, interval, maxRuns } = job.loader;
  const dayOfWeekLabel =
    weekDays && weekDays.length > 0
      ? weekDays.map((item) => t(`dayOfWeek.short.${item}`)).join("; ")
      : null;
  const atTimeLabel = atTimes && atTimes.length > 0 ? atTimes.join("; ") : null;
  return (
    <Card>
      <MDBox pt={3} px={2}>
        <MDTypography variant="h6" fontWeight="medium" textTransform="uppercase">
          {name}
        </MDTypography>
        {description}
      </MDBox>
      <MDBox pt={1} pb={0} px={2}>
        <MDBox component="ul" display="flex" flexDirection="column" p={0} m={0}>
          {renderAtTimes("Дата следующего запуска", nextRun)}
          {renderAtTimes("Дата последнего запуска", lastRun)}
          <Divider />
          <MDBox pb={2} px={2}>
            {weekDays && (
              <MDTypography component="p" variant="button" color="text" display="flex">
                {t("dayOfWeek.title.long")}:&nbsp;{dayOfWeekLabel}
              </MDTypography>
            )}
            {atTimeLabel && (
              <MDTypography component="p" variant="button" color="text" display="flex">
                {t("time.title.long")}:&nbsp;{atTimeLabel}
              </MDTypography>
            )}
            {interval && (
              <MDTypography component="p" variant="button" color="text" display="flex">
                Интервал закуска:&nbsp;
                {formatDistanceStrict(new Date(interval * 1000), new Date(0), { locale: ru })}
              </MDTypography>
            )}
            {maxRuns && (
              <MDTypography component="p" variant="button" color="text" display="flex">
                Максимальное кол-во запусков:&nbsp;{maxRuns}
              </MDTypography>
            )}
          </MDBox>
        </MDBox>
      </MDBox>
    </Card>
  );
}

JobCard.defaultProps = {};

JobCard.propTypes = {
  job: PropTypes.shape({
    id: PropTypes.string.isRequired,
    name: PropTypes.string.isRequired,
    description: PropTypes.string.isRequired,
    loader: PropTypes.shape({
      weekDays: PropTypes.arrayOf(PropTypes.string),
      atTimes: PropTypes.arrayOf(PropTypes.string),
      interval: PropTypes.number,
      maxRuns: PropTypes.number,
      nextRun: PropTypes.instanceOf(Date),
      lastRun: PropTypes.instanceOf(Date),
    }),
  }).isRequired,
};

export default JobCard;
