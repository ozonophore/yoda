import DashboardLayout from "../../examples/LayoutContainers/DashboardLayout";
import DashboardNavbar from "../../examples/Navbars/DashboardNavbar";
import MDBox from "../../components/MDBox";
import MDTimePicker from "../../components/MDTimePicker";

function Test() {
  return (
    <DashboardLayout>
      <DashboardNavbar />
      <MDBox py={3}>
        <MDTimePicker />
      </MDBox>
    </DashboardLayout>
  );
}

export default Test;
