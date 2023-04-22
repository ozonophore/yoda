import Card from "@mui/material/Card";
import Grid from "@mui/material/Grid";
import { format } from "date-fns";
import MDBox from "../../components/MDBox";
import MDTypography from "../../components/MDTypography";
import Footer from "../../examples/Footer";
import DashboardLayout from "../../examples/LayoutContainers/DashboardLayout";
import DashboardNavbar from "../../examples/Navbars/DashboardNavbar";
import { DefaultService } from "../../generated";
import StockDataTable from "../sales/components/StockDataTable";
import orgData from "./data/orgData";
import mpData from "./data/mpData";

export default function Dictionaries() {
  const onFetchOrgData = (
    pageIndex,
    pageSize,
    transactionDate,
    search,
    setItems,
    setTotalCount,
    setPageCount,
    setCanNextPage,
    setCanPreviousPage
  ) => {
    DefaultService.getOrders(
      pageSize,
      pageIndex * pageSize,
      format(transactionDate, "yyyy-LL-dd"),
      search || null
    )
      .then((res) => {
        setItems(res.items);
        setTotalCount(res.total);
        const count = Math.ceil(res.total / pageSize);
        setPageCount(count);
        setCanNextPage(pageIndex < count - 1);
        setCanPreviousPage(pageIndex > 0);
      })
      .catch(() => {
        setItems([]);
      });
  };

  return (
    <DashboardLayout>
      <DashboardNavbar />
      <MDBox pt={4} pb={3}>
        <Grid container spacing={6}>
          <Grid item xs={12}>
            <Card>
              <MDBox
                mx={2}
                mt={-3}
                py={2}
                px={2}
                variant="gradient"
                bgColor="info"
                borderRadius="lg"
                coloredShadow="info"
              >
                <MDTypography variant="h6" color="white">
                  Организации
                </MDTypography>
              </MDBox>
              <MDBox>
                <StockDataTable
                  canDateFilter={false}
                  canSearch
                  isSorted={false}
                  showTotalEntries
                  noEndBorder
                  columns={orgData().columns}
                  onFetchData={onFetchOrgData}
                  onRenderCell={orgData().onRenderData}
                />
              </MDBox>
            </Card>
          </Grid>
          <Grid item xs={12}>
            <Card>
              <MDBox
                mx={2}
                mt={-3}
                py={2}
                px={2}
                variant="gradient"
                bgColor="info"
                borderRadius="lg"
                coloredShadow="info"
              >
                <MDTypography variant="h6" color="white">
                  Маркетплейсы
                </MDTypography>
              </MDBox>
              <MDBox>
                <StockDataTable
                  canDateFilter={false}
                  canSearch
                  isSorted={false}
                  showTotalEntries
                  noEndBorder
                  columns={mpData().columns}
                  onFetchData={onFetchOrgData}
                  onRenderCell={mpData().onRenderData}
                />
              </MDBox>
            </Card>
          </Grid>
          <Grid item xs={12}>
            <Card>
              <MDBox
                mx={2}
                mt={-3}
                py={2}
                px={2}
                variant="gradient"
                bgColor="info"
                borderRadius="lg"
                coloredShadow="info"
              >
                <MDTypography variant="h6" color="white">
                  Товары
                </MDTypography>
              </MDBox>
              <MDBox>
                <StockDataTable
                  canDateFilter={false}
                  canSearch
                  isSorted={false}
                  showTotalEntries
                  noEndBorder
                  columns={mpData().columns}
                  onFetchData={onFetchOrgData}
                  onRenderCell={mpData().onRenderData}
                />
              </MDBox>
            </Card>
          </Grid>
        </Grid>
      </MDBox>
      <Footer />
    </DashboardLayout>
  );
}
