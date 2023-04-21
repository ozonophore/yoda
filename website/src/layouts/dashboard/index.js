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
import ReportsLineChart from "examples/Charts/LineCharts/ReportsLineChart";
import Footer from "examples/Footer";

// Material Dashboard 2 React example components
import DashboardLayout from "examples/LayoutContainers/DashboardLayout";
import DashboardNavbar from "examples/Navbars/DashboardNavbar";

// Dashboard components
// Data
import reportsLineChartData from "layouts/dashboard/data/reportsLineChartData";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { DefaultService } from "../../generated";

function Dashboard() {
  const [t] = useTranslation();
  const { sales, tasks } = reportsLineChartData;
  const [salesData, setSalesData] = useState({
    labels: [],
    datasets: { label: "Sales", data: [] },
  });
  const [salesUpdateAt, setSalesUpdateAt] = useState(null);
  useEffect(() => {
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
                icon="leaderboard"
                title="Колво загрузок сегодня"
                count="2"
                percentage={{
                  color: "",
                  amount: "",
                  label: "",
                }}
              />
            </MDBox>
          </Grid>
          <Grid item xs={12} md={6} lg={3}>
            <MDBox mb={1.5}>
              <ComplexStatisticsCard
                color="success"
                icon="store"
                title="Revenue"
                count="34k"
                percentage={{
                  color: "success",
                  amount: "+1%",
                  label: "than yesterday",
                }}
              />
            </MDBox>
          </Grid>
          <Grid item xs={12} md={6} lg={3}>
            <MDBox mb={1.5}>
              <ComplexStatisticsCard
                color="primary"
                icon="person_add"
                title="Followers"
                count="+91"
                percentage={{
                  color: "success",
                  amount: "",
                  label: "Just updated",
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
            <Grid item xs={12} md={6} lg={4}>
              <MDBox mb={3}>
                <ReportsLineChart
                  color="success"
                  title="daily sales"
                  description={
                    <>
                      (<strong>+15%</strong>) increase in today sales.
                    </>
                  }
                  date="updated 4 min ago"
                  chart={sales}
                />
              </MDBox>
            </Grid>
            <Grid item xs={12} md={6} lg={4}>
              <MDBox mb={3}>
                <ReportsLineChart
                  color="dark"
                  title="completed tasks"
                  description="Last Campaign Performance"
                  date="just updated"
                  chart={tasks}
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
