import * as React from 'react';
import {createContext, Fragment, useEffect, useState} from 'react';
import Table from '@mui/joy/Table';
import Sheet from '@mui/joy/Sheet';
// icons
import {Orders} from 'client';
import YodaPagination from 'components/YodaPagination';
import {LinearProgress, Tooltip} from '@mui/joy';
import Typography from '@mui/joy/Typography';
import Input from '@mui/joy/Input';
import deepEqual from 'deep-equal';
import Box from '@mui/joy/Box';

import CheckRoundedIcon from '@mui/icons-material/CheckRounded';
import ClearRoundedIcon from '@mui/icons-material/ClearRounded';
import IconButton from '@mui/joy/IconButton';
import {jsx} from '@emotion/react';
import dayjs from "dayjs";
import JSX = jsx.JSX;

interface IDataGridContext {

}

export const DataGridContext = createContext<IDataGridContext>({})

function labelDisplayedRows({
                                from,
                                to,
                                count,
                            }: {
    from: number;
    to: number;
    count: number;
}) {
    return `${from}–${to} of ${count !== -1 ? count : `more than ${to}`}`;
}

const initOrders: Orders = {items: [], count: 0}

interface IComponentProps {
    value: string
    onChange: (value: string | null) => void
}

export interface IColumn {
    field: string,
    editable?: boolean,
    noWrap?: boolean,
    type?: 'int' | 'number' | 'string' | 'bool' | 'date',
    width?: number | string | undefined,
    minWith?: number | string | undefined,
    headerName?: string | undefined,
    description?: string | undefined,
    textAlign?: 'left' | 'right' | 'center' | undefined,
    headerTextAlign?: 'left' | 'right' | 'center' | undefined,
    renderComponent?: (props: IComponentProps) => JSX.Element,
}

interface IJoyDataSortModel {
}

interface IJoyDataFilterModel {
}

interface IJoyDataPaginationModel {
    page: number
    pageSize: number
}

export interface IJoyDataRowsParams {
    start: number;
    offset: number;
    sortModel?: IJoyDataSortModel;
    filterModel?: IJoyDataFilterModel;
    paginationModel?: IJoyDataPaginationModel;
}

export interface IGetRowsResponse<R> {
    rows: R[];
    rowCount: number;
    pageInfo?: {
        hasNextPage?: boolean;
        truncated?: number;
    };
}

export interface IJoyDataSource<R> {
    getRows(params?: IJoyDataRowsParams): Promise<IGetRowsResponse<R>>

    onPageChange(page: number): void
}

interface IProps<R> {
    size?: 'sm' | 'md' | 'lg' | undefined,
    columns: IColumn[],
    rows: R[],
    page: number,
    pageSize: number,
    count: number,
    onFilterModelChange?: (filterModel: IJoyDataFilterModel) => void | undefined
    onPaginationModelChange?: (paginationModel: IJoyDataPaginationModel) => void | undefined
    onPageChange: (page: number) => void
    onPageSizeChange: (pageSize: number) => void
    showColumns?: boolean
    isLoading: boolean,
    onSave?: (oldValue: R, newValue: R) => void,
    onRefresh?: (page: number) => void,
}

interface IData<R> {
    oldData: R
    newData: R
}

interface IDataGridInputProps {
    value: any,
    onChange: (value: string) => void,
    type?: 'int' | 'number' | 'string' | 'bool' | 'date'
}

function DataGridInput({value, onChange, type}: IDataGridInputProps) {
    return (<Input size='sm' value={value} variant="plain"
                   onChange={(e) => {
                       onChange(e.target.value)
                   }}
                   sx={{
                       border: 0,
                       color: 'var(--joy-palette-text-tertiary, var(--joy-palette-neutral-600, #555E68))',
                       background: 'transparent'
                   }}/>)
}

interface IRowDataGridProps<R> {
    keyValue?: number | string | null,
    columns: IColumn[],
    row: IData<R>,
    onSave: (oldValue: R, newValue: R) => void,
    editable?: boolean
}

function JoyRowDataGrid<R>({keyValue, columns, row, onSave, editable}: IRowDataGridProps<R>) {
    const [saveDisable, setSaveDisabled] = useState(deepEqual(row.oldData, row.newData))
    const [data, setData] = useState(row.newData)

    function handleOnCancel() {
        setData(row.oldData)
        setSaveDisabled(true)
    }

    useEffect(() => {
        setData(row.newData)
    }, [row]);

    useEffect(() => {

    }, [saveDisable]);

    function getValue(data: R, column: IColumn): string {
        if (column.type === 'date') {
            const v = (data as Record<string, any>)[column.field]
            const date = dayjs(v).format('YYYY-MM-DD')
            return date
        }
        return (data as Record<string, any>)[column.field]
    }

    return (<tr key={keyValue}>
        {
            columns.map(column => <td key={`${keyValue}_${column.field}`} style={{textAlign: column.textAlign}}>
                {column.editable &&
                    (column.renderComponent ?
                        column.renderComponent({
                            value: (data as Record<string, any>)[column.field],
                            onChange: (value) => {
                                const newData = {
                                    ...data,
                                    [column.field]: value
                                }
                                setData(newData)
                                setSaveDisabled(deepEqual(row.oldData, newData))
                            }
                        }) :
                        <DataGridInput value={(data as Record<string, any>)[column.field]}
                                       onChange={(value: string) => {
                                           const newData = {
                                               ...data,
                                               [column.field]: value
                                           }
                                           setData(newData)
                                           setSaveDisabled(deepEqual(row.oldData, newData))
                                       }} type={column.type}/>)
                }
                {
                    !Boolean(column.editable) &&
                    <Tooltip title={getValue(data, column)}>
                        <Typography noWrap={column.noWrap} level='body-xs' style={{minWidth: column.minWith}}>
                            {getValue(data, column)}
                        </Typography>
                    </Tooltip>
                }
            </td>)
        }
        {editable &&
            <td>
                <Box sx={{display: 'flex', gap: 1}}>
                    <IconButton disabled={saveDisable} variant='outlined' color='success'
                                size='sm'
                                onClick={() => {
                                    onSave(row.oldData, data)
                                    setSaveDisabled(!saveDisable)
                                }}
                    >
                        <CheckRoundedIcon/>
                    </IconButton>
                    <IconButton disabled={saveDisable} onClick={handleOnCancel} variant='outlined' color='primary'
                                size='sm'>
                        <ClearRoundedIcon/>
                    </IconButton>
                </Box>
            </td>
        }
    </tr>)
}

export default function JoyDataGrid<R>({
                                           columns,
                                           rows,
                                           page,
                                           pageSize,
                                           count,
                                           onPageChange,
                                           onPageSizeChange,
                                           isLoading,
                                           onSave,
                                           onRefresh,
                                           size = 'sm',
                                           showColumns = false
                                       }: IProps<R>): React.JSX.Element {

    const [data, setData] = useState<{ oldData: R, newData: R }[]>([])
    useEffect(() => {
        const newData: IData<R>[] = rows.map(row => {
            return {
                oldData: row,
                newData: row,
            }
        })
        setData(newData)
    }, [rows]);

    useEffect(() => {
        if (isLoading) {
            setData([])
        }
    }, [isLoading]);

    function handleOnSave(oldValue: R, newValue: R) {
        if (onSave) {
            onSave(oldValue, newValue)
        }
    }

    const renderRows = (editable: boolean) => {

        return (
            <Fragment>
                {
                    data.map((row, index) =>
                        (<JoyRowDataGrid key={`row_${index}`} columns={columns} row={row} onSave={handleOnSave}
                                         editable={editable}/>)
                    )
                }
            </Fragment>
        )
    }

    const handleChangePage = (newPage: number) => {
        onPageChange(newPage)
    };

    const handleChangeRowsPerPage = (event: any, newValue: number | null) => {
        const newPageSize = newValue ?? 25
        onPageSizeChange(newPageSize)
        onPageChange(0)
    };

    const handleChangeFilterText = (event: any) => {
        if (!event.target.value || event.target.value.length < 3) {
            return
        }
        // setFilter(event.target.value);
    }

    const handleOnKeyDown = (event: any) => {
        if (event.key === 'Enter') {
            event.preventDefault();
        }
    }

    const editableTable = columns.filter(colum => colum.editable).length > 0

    let tableSx = {
        '--TableCell-headBackground': 'var(--joy-palette-background-level1)',
        '--Table-headerUnderlineThickness': '1px',
        '--TableRow-hoverBackground': 'var(--joy-palette-background-level1)',
        '--TableCell-paddingY': '4px',
        '--TableCell-paddingX': '8px',
        '& tr > th': {textAlign: 'center'},
    }

    if (editableTable) {
        tableSx = Object.assign(tableSx, {
            '& tr > *:last-child': {
                position: 'sticky',
                right: 0,
                bgcolor: 'var(--TableCell-headBackground)'
            }
        })
    }

    return (
        <React.Fragment>
            <Sheet
                className="OrderTableContainer"
                variant="outlined"
                sx={{
                    display: {xs: 'none', sm: 'initial'},
                    width: '100%',
                    borderRadius: 'sm',
                    flexShrink: 1,
                    overflow: 'auto',
                    minHeight: 0,
                    height: '100%',
                    '--Table-lastColumnWidth': '88px',
                }}
            >
                <Table size={size}
                       aria-labelledby="tableTitle"
                       stickyHeader
                       hoverRow
                       borderAxis={showColumns ? 'both' : 'xBetween'}
                       sx={tableSx}
                >
                    <thead>
                    <tr key='grid_header'>
                        {
                            columns.map(column =>
                                <th key={`grid_header_${column.field}`} style={{
                                    width: column.width,
                                    padding: '12px 6px',
                                    textAlign: column.headerTextAlign,
                                    minWidth: column.minWith,
                                    whiteSpace: 'normal',
                                }}>
                                    <Tooltip size='sm' title={column.headerName} component="div">
                                        <Typography level='title-sm'>
                                            {column.headerName ?? column.field}
                                        </Typography>
                                    </Tooltip>
                                </th>
                            )
                        }
                        {editableTable &&
                            <th
                                aria-label="last"
                                style={{width: 'var(--Table-lastColumnWidth)'}}
                            />
                        }
                    </tr>
                    </thead>
                    <tbody>
                    {
                        !isLoading && renderRows(editableTable)
                    }
                    </tbody>
                </Table>
                {
                    isLoading &&
                    <LinearProgress size={size}/>
                }
            </Sheet>
            <YodaPagination page={page} count={count} rowsPerPage={pageSize}
                            pageLength={page * pageSize + rows.length}
                            onChangeRowsPerPage={handleChangeRowsPerPage}
                            onChangePage={handleChangePage}
                            onRefresh={onRefresh}
            />
        </React.Fragment>
    );
}
