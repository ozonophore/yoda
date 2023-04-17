import Card from "@mui/material/Card";
import Grid from "@mui/material/Grid";
import { format } from "date-fns";
import MDBox from "../../components/MDBox";
import MDTypography from "../../components/MDTypography";
import Footer from "../../examples/Footer";
import DashboardLayout from "../../examples/LayoutContainers/DashboardLayout";
import DashboardNavbar from "../../examples/Navbars/DashboardNavbar";
import { DefaultService } from "../../generated";
import StockDataTable from "./components/StockDataTable";
import orderData from "./data/orderData";
import stockData from "./data/stockData";

function Sales() {
  const onFetchStockData = (
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
    DefaultService.getStocks(
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

  const onFetchOrderData = (
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
                  Остатки
                </MDTypography>
              </MDBox>
              <MDBox>
                <StockDataTable
                  canSearch
                  isSorted={false}
                  showTotalEntries
                  noEndBorder
                  columns={stockData().columns}
                  onFetchData={onFetchStockData}
                  onRenderCell={stockData().onRenderData}
                />
              </MDBox>
            </Card>
          </Grid>
          <Grid item xs={12}>
            <Card>
              <MDBox
                mx={2}
                mt={-3}
                py={3}
                px={2}
                variant="gradient"
                bgColor="info"
                borderRadius="lg"
                coloredShadow="info"
              >
                <MDTypography variant="h6" color="white">
                  Продажи
                </MDTypography>
              </MDBox>
              <MDBox>
                <StockDataTable
                  canSearch
                  isSorted={false}
                  showTotalEntries
                  noEndBorder
                  columns={orderData().columns}
                  onFetchData={onFetchOrderData}
                  onRenderCell={orderData().onRenderData}
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

export default Sales;
