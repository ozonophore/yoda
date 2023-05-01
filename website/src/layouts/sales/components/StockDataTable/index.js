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

import Autocomplete from "@mui/material/Autocomplete";
import CircularProgress from "@mui/material/CircularProgress";
import Icon from "@mui/material/Icon";

// @mui material components
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableContainer from "@mui/material/TableContainer";
import TableRow from "@mui/material/TableRow";

// Material Dashboard 2 React components
import MDBox from "components/MDBox";
import MDInput from "components/MDInput";
import MDPagination from "components/MDPagination";
import MDTypography from "components/MDTypography";
import { format, isValid, parseISO } from "date-fns";
import DataTableBodyCell from "examples/Tables/DataTable/DataTableBodyCell";

// Material Dashboard 2 React example components
import DataTableHeadCell from "examples/Tables/DataTable/DataTableHeadCell";

// prop-types is a library for typechecking of props
import PropTypes from "prop-types";
import React, { useEffect, useMemo, useState } from "react";
import { useTranslation } from "react-i18next";

// react-table components
import { useAsyncDebounce, useGlobalFilter, usePagination, useSortBy, useTable } from "react-table";

function StockDataTable({
  entriesPerPage,
  canSearch,
  canDateFilter,
  showTotalEntries,
  pagination,
  isSorted,
  noEndBorder,
  columns,
  onFetchData,
  onRenderCell,
}) {
  const [t] = useTranslation();
  const [search, setSearch] = useState(undefined);
  const [transactionDate, setTransactionDate] = useState(new Date());
  const [totalCount, setTotalCount] = useState(0);
  const [pageIndex, setPageIndex] = useState(0);
  const [canNextPage, setCanNextPage] = useState(true);
  const [pageCount, setPageCount] = useState(0);
  const [canPreviousPage, setCanPreviousPage] = useState(true);
  const [items, setItems] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const defaultValue = entriesPerPage.defaultValue ? entriesPerPage.defaultValue : 10;
  const [pageSize, setPageSize] = useState(defaultValue);
  const entries = entriesPerPage.entries
    ? entriesPerPage.entries.map((el) => el.toString())
    : ["5", "10", "15", "20", "25"];
  const data = useMemo(() => items.map((item) => onRenderCell(item)), [items]);
  const gotoPage = (page) => {
    setPageIndex(page);
  };
  const previousPage = () => {
    setPageIndex(pageIndex - 1);
  };
  const nextPage = () => {
    setPageIndex(pageIndex + 1);
  };
  const setGlobalFilter = (value) => {
    setPageIndex(0);
    setSearch(value);
  };
  const tableInstance = useTable(
    {
      columns,
      data,
      manualPagination: true,
    },
    useGlobalFilter,
    useSortBy,
    usePagination
  );

  const { getTableProps, getTableBodyProps, headerGroups, prepareRow, rows, page, pageOptions } =
    tableInstance;

  const wrapSetItems = (dt) => {
    setItems(dt);
    setIsLoading(false);
  };

  // Set the default value for the entries per page when component mounts
  useEffect(() => {
    setIsLoading(true);
    onFetchData(
      pageIndex,
      pageSize,
      transactionDate,
      search,
      wrapSetItems,
      setTotalCount,
      setPageCount,
      setCanNextPage,
      setCanPreviousPage
    );
  }, [pageIndex, pageSize, transactionDate, search]);

  // Set the entries per page value based on the select value
  const setEntriesPerPage = (value) => setPageSize(value);

  // Render the paginations
  const options = [];
  for (let i = 0; i < pageCount; i += 1) {
    if (i > pageIndex - 2) {
      options.push(i);
    }
    if (i > pageIndex + 1) {
      break;
    }
  }
  const renderPagination = options.map((option) => (
    <MDPagination
      count={2}
      item
      key={option}
      onClick={() => gotoPage(Number(option))}
      active={pageIndex === option}
    >
      {option + 1}
    </MDPagination>
  ));

  // Handler for the input to set the pagination index
  const handleInputPagination = ({ target: { value } }) =>
    value > pageOptions.length || value < 0 ? gotoPage(0) : gotoPage(Number(value));

  // Customized page options starting from 1
  const customizedPageOptions = pageOptions.map((option) => option + 1);

  // Setting value for the pagination input
  const handleInputPaginationValue = ({ target: value }) => gotoPage(Number(value.value - 1));

  // Search input state handle
  const onSearchChange = useAsyncDebounce((value) => {
    setGlobalFilter(value || undefined);
  }, 100);

  // A function that sets the sorted value for the table
  const setSortedValue = (column) => {
    let sortedValue;

    if (isSorted && column.isSorted) {
      sortedValue = column.isSortedDesc ? "desc" : "asce";
    } else if (isSorted) {
      sortedValue = "none";
    } else {
      sortedValue = false;
    }

    return sortedValue;
  };

  // Setting the entries starting point
  const entriesStart = pageIndex === 0 ? pageIndex + 1 : pageIndex * pageSize + 1;

  // Setting the entries ending point
  let entriesEnd;

  if (pageIndex === 0) {
    entriesEnd = pageSize;
  } else if (pageIndex === pageOptions.length - 1) {
    entriesEnd = rows.length;
  } else {
    entriesEnd = pageSize * (pageIndex + 1);
  }

  return (
    <MDBox>
      <MDBox>
        {entriesPerPage || canSearch ? (
          <MDBox display="flex" justifyContent="space-between" alignItems="center" p={3}>
            {entriesPerPage && (
              <MDBox display="flex" alignItems="center">
                <Autocomplete
                  disableClearable
                  value={pageSize.toString()}
                  options={entries}
                  onChange={(event, newValue) => {
                    setEntriesPerPage(parseInt(newValue, 10));
                  }}
                  size="small"
                  sx={{ width: "5rem" }}
                  renderInput={(params) => <MDInput {...params} />}
                />
                <MDTypography variant="caption" color="secondary">
                  &nbsp;&nbsp;{t("dataTable.rowsPerPage")}
                </MDTypography>
              </MDBox>
            )}
            <MDBox display="flex" alignItems="right">
              {canSearch && (
                <MDBox width="12rem" ml="auto">
                  <MDInput
                    placeholder={t("search.label")}
                    value={search || ""}
                    size="small"
                    fullWidth
                    onChange={({ currentTarget }) => {
                      onSearchChange(currentTarget.value);
                    }}
                  />
                </MDBox>
              )}
              {canDateFilter && (
                <MDBox pl={2} width="12rem" ml="auto">
                  <MDInput
                    type="date"
                    size="small"
                    value={format(transactionDate, "yyyy-MM-dd")}
                    fullWidth
                    onChange={({ currentTarget }) => {
                      if (!isValid(parseISO(currentTarget.value))) {
                        return;
                      }
                      const dt = parseISO(currentTarget.value);
                      setTransactionDate(dt);
                      setPageIndex(0);
                    }}
                  />
                </MDBox>
              )}
            </MDBox>
          </MDBox>
        ) : null}
      </MDBox>
      <TableContainer sx={{ boxShadow: "none" }}>
        <Table {...getTableProps()}>
          <MDBox component="thead">
            {headerGroups.map((headerGroup) => (
              <TableRow {...headerGroup.getHeaderGroupProps()}>
                {headerGroup.headers.map((column) => (
                  <DataTableHeadCell
                    {...column.getHeaderProps(isSorted && column.getSortByToggleProps())}
                    width={column.width ? column.width : "auto"}
                    align={column.align ? column.align : "left"}
                    sorted={setSortedValue(column)}
                    style={{ minWidth: column.minWidth }}
                  >
                    {column.render("Header")}
                  </DataTableHeadCell>
                ))}
              </TableRow>
            ))}
          </MDBox>
          {isLoading && (
            <MDBox display="flex" justifyContent="center" alignItems="center" pt={2}>
              <CircularProgress color="info" size={25} />
              <MDTypography pl={2} variant="caption" color="secondary">
                Loading...
              </MDTypography>
            </MDBox>
          )}
          {!isLoading && (
            <TableBody {...getTableBodyProps()}>
              {page.map((row, key) => {
                prepareRow(row);
                return (
                  <TableRow {...row.getRowProps()}>
                    {row.cells.map((cell) => (
                      <DataTableBodyCell
                        noBorder={noEndBorder && rows.length - 1 === key}
                        align={cell.column.align ? cell.column.align : "left"}
                        {...cell.getCellProps()}
                      >
                        {cell.render("Cell")}
                      </DataTableBodyCell>
                    ))}
                  </TableRow>
                );
              })}
            </TableBody>
          )}
        </Table>
      </TableContainer>
      <MDBox
        display="flex"
        flexDirection={{ xs: "column", sm: "row" }}
        justifyContent="space-between"
        alignItems={{ xs: "flex-start", sm: "center" }}
        p={!showTotalEntries && pageOptions.length === 1 ? 0 : 3}
      >
        {showTotalEntries && (
          <MDBox mb={{ xs: 3, sm: 0 }}>
            <MDTypography variant="button" color="secondary" fontWeight="regular">
              {t("dataTable.totalEntries", {
                entriesStart,
                entriesEnd,
                totalEntries: totalCount,
              })}
            </MDTypography>
          </MDBox>
        )}
        {options.length > 1 && (
          <MDPagination
            variant={pagination.variant ? pagination.variant : "gradient"}
            color={pagination.color ? pagination.color : "info"}
          >
            {canPreviousPage && (
              <MDPagination item onClick={() => previousPage()}>
                <Icon sx={{ fontWeight: "bold" }}>chevron_left</Icon>
              </MDPagination>
            )}
            {options.length > 6 ? (
              <MDBox width="5rem" mx={1}>
                <MDInput
                  inputProps={{ type: "number", min: 1, max: customizedPageOptions.length }}
                  value={customizedPageOptions[pageIndex]}
                  onChange={(handleInputPagination, handleInputPaginationValue)}
                />
              </MDBox>
            ) : (
              renderPagination
            )}
            {canNextPage && (
              <MDPagination item onClick={() => nextPage()}>
                <Icon sx={{ fontWeight: "bold" }}>chevron_right</Icon>
              </MDPagination>
            )}
          </MDPagination>
        )}
      </MDBox>
    </MDBox>
  );
}

// Setting default values for the props of DataTable
StockDataTable.defaultProps = {
  entriesPerPage: { defaultValue: 10, entries: [5, 10, 15, 20, 25] },
  canSearch: false,
  canDateFilter: true,
  showTotalEntries: true,
  pagination: { variant: "gradient", color: "info" },
  isSorted: true,
  noEndBorder: false,
};

// Typechecking props for the DataTable
StockDataTable.propTypes = {
  entriesPerPage: PropTypes.oneOfType([
    PropTypes.shape({
      defaultValue: PropTypes.number,
      entries: PropTypes.arrayOf(PropTypes.number),
    }),
    PropTypes.bool,
  ]),
  canSearch: PropTypes.bool,
  canDateFilter: PropTypes.bool,
  showTotalEntries: PropTypes.bool,
  pagination: PropTypes.shape({
    variant: PropTypes.oneOf(["contained", "gradient"]),
    color: PropTypes.oneOf([
      "primary",
      "secondary",
      "info",
      "success",
      "warning",
      "error",
      "dark",
      "light",
    ]),
  }),
  isSorted: PropTypes.bool,
  noEndBorder: PropTypes.bool,
  columns: PropTypes.arrayOf(PropTypes.object).isRequired,
  onFetchData: PropTypes.func.isRequired,
  onRenderCell: PropTypes.func.isRequired,
};

export default StockDataTable;
