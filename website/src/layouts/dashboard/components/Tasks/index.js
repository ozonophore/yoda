// @mui material components
import Card from "@mui/material/Card";
import Icon from "@mui/material/Icon";
import Menu from "@mui/material/Menu";
import MenuItem from "@mui/material/MenuItem";

// Material Dashboard 2 React components
import MDBox from "components/MDBox";
import MDTypography from "components/MDTypography";

// Material Dashboard 2 React examples
import DataTable from "examples/Tables/DataTable";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { DefaultService } from "../../../../generated";
import data from "./data";

function Tasks() {
  const [t] = useTranslation();
  const [menu, setMenu] = useState(null);
  const { columns, renderRows } = data;
  const [taskInfo, setTaskInfo] = useState({ items: [], completed: 0, cenceled: 0 });

  const openMenu = ({ currentTarget }) => setMenu(currentTarget);
  const closeMenu = () => setMenu(null);

  const refreshData = () => {
    DefaultService.getTasks()
      .then((res) => {
        console.log("#getTasks", res);
        setTaskInfo(res);
      })
      .catch((err) => {
        console.error(err);
      });
  };

  const runTackImmediately = () => {
    DefaultService.runTask()
      .then((res) => {
        if (res.result) {
          refreshData();
        }
      })
      .catch((err) => {
        console.error(err);
      });
  };

  const handleOnClick = () => {
    runTackImmediately("@runTaskImmediately");
    closeMenu();
  };

  useEffect(() => {
    refreshData();
  }, []);

  const renderMenu = (
    <Menu
      id="simple-menu"
      anchorEl={menu}
      anchorOrigin={{
        vertical: "top",
        horizontal: "left",
      }}
      transformOrigin={{
        vertical: "top",
        horizontal: "right",
      }}
      open={Boolean(menu)}
      onClose={closeMenu}
    >
      <MenuItem onClick={handleOnClick}>Run immediately</MenuItem>
    </Menu>
  );

  return (
    <Card>
      <MDBox display="flex" justifyContent="space-between" alignItems="center" p={3}>
        <MDBox>
          <MDTypography variant="h6" gutterBottom>
            {t("tasks.active")}
          </MDTypography>
          <MDBox display="flex" alignItems="center" lineHeight={0}>
            <Icon
              sx={{
                fontWeight: "bold",
                color: ({ palette: { info } }) => info.main,
                mt: -0.5,
              }}
            >
              done
            </Icon>
            <MDTypography pr={2} variant="button" fontWeight="regular" color="text">
              &nbsp;<strong>{taskInfo.completed} done</strong>
            </MDTypography>
            <Icon
              sx={{
                fontWeight: "bold",
                color: ({ palette: { error } }) => error.main,
                mt: -0.5,
              }}
            >
              close
            </Icon>
            <MDTypography variant="button" fontWeight="regular" color="text">
              &nbsp;<strong>{taskInfo.canceled} undone</strong>
            </MDTypography>
          </MDBox>
        </MDBox>
        <MDBox color="text" px={2}>
          <Icon
            sx={{ cursor: "pointer", fontWeight: "bold" }}
            fontSize="small"
            onClick={refreshData}
          >
            refresh
          </Icon>
          <Icon sx={{ cursor: "pointer", fontWeight: "bold" }} fontSize="small" onClick={openMenu}>
            more_vert
          </Icon>
        </MDBox>
        {renderMenu}
      </MDBox>
      <MDBox>
        <DataTable
          table={{ columns, rows: renderRows(taskInfo.items) }}
          showTotalEntries={false}
          isSorted={false}
          noEndBorder
          entriesPerPage={false}
        />
      </MDBox>
    </Card>
  );
}

export default Tasks;
