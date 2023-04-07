import Grid from "@mui/material/Grid";
import DashboardLayout from "../../examples/LayoutContainers/DashboardLayout";
import DashboardNavbar from "../../examples/Navbars/DashboardNavbar";
import RegularCard from "./components/RegularCard";

function Schedulers() {
  return (
    <DashboardLayout>
      <DashboardNavbar />
      <Grid container spacing={3}>
        <Grid item xs={12} md={6}>
          <RegularCard />
        </Grid>
        <Grid item xs={12} md={6}>
          <RegularCard />
        </Grid>
      </Grid>
    </DashboardLayout>
  );
}

export default Schedulers;
