import MDBox from "../../components/MDBox";
import Footer from "../../examples/Footer";
import DashboardLayout from "../../examples/LayoutContainers/DashboardLayout";
import DashboardNavbar from "../../examples/Navbars/DashboardNavbar";
import JobGrid from "./components/JobGrid";

function Schedulers() {
  return (
    <DashboardLayout>
      <DashboardNavbar />
      <MDBox pb={3}>
        <JobGrid />
      </MDBox>
      <Footer />
    </DashboardLayout>
  );
}

export default Schedulers;
