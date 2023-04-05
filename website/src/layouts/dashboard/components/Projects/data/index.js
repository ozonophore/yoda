/* eslint-disable react/prop-types */
/* eslint-disable react/function-component-definition */
/**
 =========================================================
 * Material Dashboard 2 React - v2.1.0
 =========================================================

 * Product Page: https://www.creative-tim.com/product/material-dashboard-react
 * Copyright 2022 Creative Tim (https://www.creative-tim.com)

 Coded by www.creative-tim.com

 =========================================================

 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 */

// @mui material components
import MDBox from "components/MDBox";
import MDTypography from "components/MDTypography";
import MDProgress from "components/MDProgress";
import Icon from "@mui/material/Icon";

export default function data() {
  const rows = [].map(({ company, jobName }) => ({
    companies: (
      <MDTypography variant="button" fontWeight="medium" ml={1} lineHeight={1}>
        {company}
      </MDTypography>
    ),
    tasks: (
      <MDBox display="flex" py={1}>
        {jobName}
      </MDBox>
    ),
    completion: (
      <MDBox width="8rem" textAlign="left">
        <MDProgress value={90} color="info" variant="gradient" label={false} />
      </MDBox>
    ),
    action: (
      <MDTypography component="a" href="#" color="text">
        <Icon>more_vert</Icon>
      </MDTypography>
    ),
  }));

  return {
    columns: [
      { Header: "companies", accessor: "companies", width: "25%", align: "left" },
      { Header: "tasks", accessor: "tasks", align: "left" },
      { Header: "completion", accessor: "completion", align: "center" },
      { Header: "action", accessor: "action", align: "center" },
    ],

    rows,
  };
}
