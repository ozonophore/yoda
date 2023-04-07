import Card from "@mui/material/Card";
import MDBox from "../../../../components/MDBox";
import MDTypography from "../../../../components/MDTypography";

function RegularCard() {
  return (
    <Card>
      <MDBox pt={3} px={2}>
        <MDTypography variant="h6" fontWeight="medium" textTransform="capitalize">
          регулярная загрузка по расписанию
        </MDTypography>
      </MDBox>
      <MDBox pt={1} pb={2} px={2}>
        <MDBox component="ul" display="flex" flexDirection="column" p={0} m={0}>
          Schedule
        </MDBox>
      </MDBox>
    </Card>
  );
}

export default RegularCard;
