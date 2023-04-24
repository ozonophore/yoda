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

const renderRows = (items) => {
  console.log("#", items);
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
        <Chip label={status} color={status} size="small" />
      </MDTypography>
    ),
  }));
};

export default { columns, renderRows };
