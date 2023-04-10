import Grid from "@mui/material/Grid";
import DashboardLayout from "../../examples/LayoutContainers/DashboardLayout";
import DashboardNavbar from "../../examples/Navbars/DashboardNavbar";
import RegularCard from "./components/RegularCard";

function Schedulers() {
  return (
    <DashboardLayout>
      <DashboardNavbar />
      <Grid container spacing={3}>
        <Grid key="job_1" item xs={12} md={6}>
          <RegularCard title="регулярная загрузка по расписанию" />
        </Grid>
        <Grid key="job_2" item xs={12} md={6}>
          <div />
        </Grid>
      </Grid>
    </DashboardLayout>
  );
}

export default Schedulers;
