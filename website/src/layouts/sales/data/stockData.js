export default function data() {
  return {
    columns: [
      { Header: "Наименование", accessor: "name", width: "30%", align: "left" },
      { Header: "Орг", accessor: "organization", align: "left", minWidth: 170 },
      { Header: "Маркетплейс", accessor: "marketplace", align: "center", minWidth: 170 },
      { Header: "Артикл", accessor: "article", align: "center", minWidth: 170 },
      { Header: "Баркод", accessor: "barcode", align: "center", minWidth: 170 },
      { Header: "Кол-во", accessor: "quantity", align: "right", minWidth: 50, maxWidth: 50 },
    ],
    onRenderData: (item) => item,
  };
}
