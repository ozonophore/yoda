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
import Grid from "@mui/material/Grid";

// Material Dashboard 2 React components
import MDBox from "components/MDBox";
import { format, parseISO } from "date-fns";
import ComplexStatisticsCard from "examples/Cards/StatisticsCards/ComplexStatisticsCard";
import ReportsBarChart from "examples/Charts/BarCharts/ReportsBarChart";
import Footer from "examples/Footer";

// Material Dashboard 2 React example components
import DashboardLayout from "examples/LayoutContainers/DashboardLayout";
import DashboardNavbar from "examples/Navbars/DashboardNavbar";

// Dashboard components
// Data
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { DefaultService } from "../../generated";

const initTransactionInfo = {
  total: 0,
  success: 0,
};

function Dashboard() {
  const [t] = useTranslation();
  const [salesData, setSalesData] = useState({
    labels: [],
    datasets: { label: "Sales", data: [] },
  });
  const [transactionInfo, setTransactionInfo] = useState(initTransactionInfo);
  const [salesUpdateAt, setSalesUpdateAt] = useState(null);
  useEffect(() => {
    DefaultService.getTransactionsInfo()
      .then((res) => {
        setTransactionInfo(res);
      })
      .catch((err) => {
        console.error(err);
      });

    DefaultService.getSalesForWeek()
      .then((res) => {
        const lastUpdate = parseISO(res.updateAt);
        const days = res.items.map((item) => format(new Date(item.orderDate), "dd-MM"));
        const prices = res.items.map((item) => item.price);
        setSalesData({
          ...salesData,
          labels: days,
          datasets: {
            label: "Sales",
            data: prices,
          },
        });
        setSalesUpdateAt(lastUpdate);
      })
      .catch((err) => {
        console.error(err);
      });
  }, []);

  return (
    <DashboardLayout>
      <DashboardNavbar />
      <MDBox py={3}>
        <Grid container spacing={3}>
          <Grid item xs={12} md={6} lg={3}>
            <MDBox mb={1.5}>
              <ComplexStatisticsCard
                color="dark"
                icon="weekend"
                title="Организаций"
                count={1}
                percentage={{
                  color: "success",
                  amount: "",
                  label: "",
                }}
              />
            </MDBox>
          </Grid>
          <Grid item xs={12} md={6} lg={3}>
            <MDBox mb={1.5}>
              <ComplexStatisticsCard
                icon="schedule"
                title="Загрузок"
                count={transactionInfo.total}
                percentage={{
                  color: "success",
                  amount: transactionInfo.success,
                  label: "Успешных",
                }}
              />
            </MDBox>
          </Grid>
        </Grid>
        <MDBox mt={4.5}>
          <Grid container spacing={3}>
            <Grid item xs={12} md={6} lg={4}>
              <MDBox mb={3}>
                <ReportsBarChart
                  color="info"
                  title={t("dashboard.weeklySales")}
                  description=""
                  date={`Дата расчета: ${
                    salesUpdateAt ? format(salesUpdateAt, "dd-MM-yyyy HH:mm") : " Не определена"
                  }`}
                  chart={salesData}
                />
              </MDBox>
            </Grid>
          </Grid>
        </MDBox>
      </MDBox>
      <Footer />
    </DashboardLayout>
  );
}

export default Dashboard;
