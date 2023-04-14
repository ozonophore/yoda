import Divider from "@mui/material/Divider";
import DataTable from "react-data-table-component";
import Card from "@mui/material/Card";
import Icon from "@mui/material/Icon";
import * as React from "react";
import DashboardLayout from "../../examples/LayoutContainers/DashboardLayout";
import DashboardNavbar from "../../examples/Navbars/DashboardNavbar";
import MDBox from "../../components/MDBox";
import MDTimePicker from "../../components/MDTimePicker";
import TimePickers from "./TimePickers";
import YDataTable from "./YDataTable";
import MDPagination from "../../components/MDPagination";

// const columns = [
//   {
//     name: "Title",
//     selector: (row) => row.title,
//     sortable: true,
//     width: "670px",
//     right: true,
//   },
//   {
//     name: "Year",
//     selector: (row) => row.year,
//     sortable: true,
//     width: "470px",
//     right: true,
//   },
// ];
const columns = [
  { id: "name", selector: (row) => row.name, name: "Name", right: true, width: "100px" },
  {
    id: "code",
    selector: (row) => row.code,
    name: "ISO\u00a0Code",
    minWidth: "100px",
    width: "170px",
  },
  {
    id: "population",
    selector: (row) => row.population,
    name: "Population",
    minWidth: "100px",
    width: "170px",
    align: "right",
  },
  {
    id: "size",
    selector: (row) => row.size,
    name: "Size\u00a0(km\u00b2)",
    minWidth: "100px",
    width: "170px",
    align: "right",
  },
  {
    id: "density",
    selector: (row) => row.density,
    name: "Density",
    minWidth: "100px",
    width: "170px",
    align: "right",
  },
];

function createData(id, name, code, population, size) {
  const density = population / size;
  return { id, name, code, population, size, density };
}

const data = [
  createData(1, "India", "IN", 1324171354, 3287263),
  createData(2, "China", "CN", 1403500365, 9596961),
  createData(3, "Italy", "IT", 60483973, 301340),
  createData(4, "United States", "US", 327167434, 9833520),
  createData(5, "Canada", "CA", 37602103, 9984670),
  createData(6, "Australia", "AU", 25475400, 7692024),
  createData(7, "Germany", "DE", 83019200, 357578),
  createData(8, "Ireland", "IE", 4857000, 70273),
  createData(9, "Mexico", "MX", 126577691, 1972550),
  createData(10, "Japan", "JP", 126317000, 377973),
  createData(11, "France", "FR", 67022000, 640679),
  createData(12, "United Kingdom", "GB", 67545757, 242495),
  createData(13, "Russia", "RU", 146793744, 17098246),
  createData(14, "Nigeria", "NG", 200962417, 923768),
  createData(15, "Brazil", "BR", 210147125, 8515767),
];

function CustomMaterialPagination() {
  return (
    <MDBox p={1}>
      <MDPagination variant="gradient" color="info">
        <MDPagination item onClick={() => {}}>
          <Icon sx={{ fontWeight: "bold" }}>chevron_left</Icon>
        </MDPagination>
        <MDPagination item onClick={() => {}}>
          <Icon sx={{ fontWeight: "bold" }}>chevron_right</Icon>
        </MDPagination>
      </MDPagination>
    </MDBox>
  );
}

function Test() {
  return (
    <DashboardLayout>
      <DashboardNavbar />
      <MDBox>
        <MDTimePicker />
      </MDBox>
      <MDBox>
        <TimePickers />
      </MDBox>
      <MDBox py={3}>
        <Card>
          <DataTable
            pagination
            paginationComponent={CustomMaterialPagination}
            columns={columns}
            data={data}
          />
        </Card>
      </MDBox>
      <div>--------------</div>
      <MDBox py={3}>
        <Card>
          <MDBox>
            <YDataTable />
          </MDBox>
        </Card>
        <div>------------</div>
      </MDBox>
      <Divider />
      <MDBox py={3}>
        <
      </MDBox>
    </DashboardLayout>
  );
}

export default Test;
