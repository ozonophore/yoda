export default function data() {
  return {
    columns: [
      {
        Header: "Наименование",
        accessor: "name",
        width: "100%",
        align: "left",
        style: { overflow: "visible" },
      },
    ],
    onRenderData: (item) => item,
  };
}
