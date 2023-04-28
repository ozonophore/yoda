/* eslint-disable react/prop-types */
/* eslint-disable react/function-component-definition */

import { Chip } from "@mui/material";
// @mui material components
import MDBox from "components/MDBox";
import { lowerCase, upperCase } from "lodash";
import React from "react";

const columns = [
  { Header: "Дата начала", accessor: "startDate", width: "25%", align: "left" },
  { Header: "Дата окончания", accessor: "endDate", align: "left" },
  { Header: "Статус", accessor: "status", align: "center" },
];

const parseStatus = (status) => {
  switch (upperCase(status)) {
    case "COMPLETED":
      return "success";
    case "CANCELED":
      return "warning";
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
      <MDBox display="flex" py={1}>
        <Chip label={lowerCase(status)} color={parseStatus(status)} size="small" />
      </MDBox>
    ),
  }));

export default { columns, renderRows };
