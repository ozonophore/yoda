import DashboardLayout from "../../examples/LayoutContainers/DashboardLayout";
import DashboardNavbar from "../../examples/Navbars/DashboardNavbar";
import MDBox from "../../components/MDBox";

function Test() {
  return (
    <DashboardLayout>
      <DashboardNavbar />
      <MDBox py={3}>Test</MDBox>
    </DashboardLayout>
  );
}

export default Test;
