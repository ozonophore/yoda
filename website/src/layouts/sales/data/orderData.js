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
      { Header: "Орг", accessor: "organization", align: "left", minWidth: 170 },
      { Header: "Маркетплейс", accessor: "marketplace", align: "center", minWidth: 170 },
      { Header: "Артикл", accessor: "article", align: "center", minWidth: 170 },
      { Header: "Баркод", accessor: "barcode", align: "center", minWidth: 170 },
      { Header: "Кол-во", accessor: "quantity", align: "right", minWidth: 50, maxWidth: 50 },
      { Header: "Цена", accessor: "price", align: "right", minWidth: 50, maxWidth: 50 },
      { Header: "Статус", accessor: "status", align: "left", minWidth: 50, maxWidth: 50 },
    ],
    onRenderData: (item) => ({
      ...item,
      name: (
        <MDBox maxWidth={150} mb={2} lineHeight={1}>
          <MDTypography variant="button" color="text">
            {item.name}
          </MDTypography>
        </MDBox>
      ),
    }),
  };
}
