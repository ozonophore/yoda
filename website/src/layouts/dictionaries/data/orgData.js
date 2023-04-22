import MDBox from "../../../components/MDBox";
import MDTypography from "../../../components/MDTypography";

export default function data() {
  return {
    columns: [
      {
        Header: "Наименование",
        accessor: "name",
        width: 170,
        maxWidth: 170,
        align: "left",
        style: { overflow: "visible" },
      },
      { Header: "Инн", accessor: "inn", align: "left", minWidth: 170 },
      { Header: "Крр", accessor: "kpp", align: "дуае", minWidth: 170 },
    ],
    onRenderData: (item) => ({
      ...item,
      name: (
        <MDBox mb={2} lineHeight={1}>
          <MDTypography variant="button" color="text">
            {item.name}
          </MDTypography>
        </MDBox>
      ),
    }),
  };
}
