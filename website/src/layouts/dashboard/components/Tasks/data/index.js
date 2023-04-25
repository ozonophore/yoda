/* eslint-disable react/prop-types */
/* eslint-disable react/function-component-definition */

import { Chip } from "@mui/material";
// @mui material components
import MDBox from "components/MDBox";
import MDTypography from "components/MDTypography";

const columns = [
  { Header: "Дата начала", accessor: "startDate", width: "25%", align: "left" },
  { Header: "Дата окончания", accessor: "endDate", align: "left" },
  { Header: "Статус", accessor: "status", align: "center" },
];

const parseStatus = (status) => {
  switch (status) {
    case "COMPLETED":
      return "success";
    case "CANCELED":
      return "error";
    case "REJECTED":
      return "error";
    case "BEGIN":
      return "info";
    default:
      return "warning";
  }
};

const renderRows = (items) =>
  items.map(({ startDate, endDate, status }) => ({
    startDate: (
      <MDBox display="flex" py={1}>
        {startDate}
      </MDBox>
    ),
    endDate: (
      <MDBox display="flex" py={1}>
        {endDate}
      </MDBox>
    ),
    status: (
      <MDTypography color="white">
        <Chip label={parseStatus(status)} color={parseStatus(status)} size="small" />
      </MDTypography>
    ),
  }));

export default { columns, renderRows };
