import * as React from 'react';
import { useEffect, useRef, useState } from 'react';
import Box from '@mui/joy/Box';
import Button from '@mui/joy/Button';
import Divider from '@mui/joy/Divider';
import FormControl from '@mui/joy/FormControl';
import FormLabel from '@mui/joy/FormLabel';
import Link from '@mui/joy/Link';
import Input from '@mui/joy/Input';
import Modal from '@mui/joy/Modal';
import ModalDialog from '@mui/joy/ModalDialog';
import ModalClose from '@mui/joy/ModalClose';
import Select from '@mui/joy/Select';
import Option from '@mui/joy/Option';
import Table from '@mui/joy/Table';
import Sheet from '@mui/joy/Sheet';
import IconButton, { iconButtonClasses } from '@mui/joy/IconButton';
import Typography from '@mui/joy/Typography';
import Menu from '@mui/joy/Menu';
import MenuButton from '@mui/joy/MenuButton';
import MenuItem from '@mui/joy/MenuItem';
import Dropdown from '@mui/joy/Dropdown';
// icons
import FilterAltIcon from '@mui/icons-material/FilterAlt';
import SearchIcon from '@mui/icons-material/Search';
import ArrowDropDownIcon from '@mui/icons-material/ArrowDropDown';
import MoreHorizRoundedIcon from '@mui/icons-material/MoreHorizRounded';
import AddRoundedIcon from '@mui/icons-material/AddRounded';
import { IRow } from './model';
import ClusterRow from './ClusterRow';


function descendingComparator<T>(a: T, b: T, orderBy: keyof T) {
    if (b[orderBy] < a[orderBy]) {
        return -1;
    }
    if (b[orderBy] > a[orderBy]) {
        return 1;
    }
    return 0;
}

type Order = 'asc' | 'desc';

function getComparator<Key extends keyof any>(
    order: Order,
    orderBy: Key,
): (
    a: { [key in Key]: number | string },
    b: { [key in Key]: number | string },
) => number {
    return order === 'desc'
        ? (a, b) => descendingComparator(a, b, orderBy)
        : (a, b) => -descendingComparator(a, b, orderBy);
}

// Since 2020 all major browsers ensure sort stability with Array.prototype.sort().
// stableSort() brings sort stability to non-modern browsers (notably IE11). If you
// only support modern browsers you can replace stableSort(exampleArray, exampleComparator)
// with exampleArray.slice().sort(exampleComparator)
function stableSort<T>(array: readonly T[], comparator: (a: T, b: T) => number) {
    const stabilizedThis = array.map((el, index) => [el, index] as [T, number]);
    stabilizedThis.sort((a, b) => {
        const order = comparator(a[0], b[0]);
        if (order !== 0) {
            return order;
        }
        return a[1] - b[1];
    });
    return stabilizedThis.map((el) => el[0]);
}

function RowMenu() {
    return (
        <Dropdown>
            <MenuButton
                slots={{root: IconButton}}
                slotProps={{root: {variant: 'plain', color: 'neutral', size: 'sm'}}}
            >
                <MoreHorizRoundedIcon/>
            </MenuButton>
            <Menu size="sm" sx={{minWidth: 140}}>
                <MenuItem>Edit</MenuItem>
                <MenuItem>Rename</MenuItem>
                <MenuItem>Move</MenuItem>
                <Divider/>
                <MenuItem color="danger">Delete</MenuItem>
            </Menu>
        </Dropdown>
    );
}

interface IRowWrapper<R> {
    isNew: boolean
    wasChanged: boolean
    newRow: R
    oldRow?: R
}

interface IEditableCellProps {
    value: string
    isEdit: boolean
    onClick: () => void
    onChange: (value: string) => void
}

function EditableCell(props: IEditableCellProps): React.JSX.Element {
    const inputRef = useRef<HTMLInputElement>(null)
    return (
        <td onClick={() => {
            console.log("# " + inputRef.current)
            inputRef.current?.focus()
            props.onClick()
        }}>
            {props.isEdit &&
                <Input ref={inputRef}
                       sx={{border: "none", bgcolor: "transparent", marginLeft: "-8px", fontSize: "12px"}}
                    // onBlur={() => {
                    //     setEditableRow(null)
                    // }}
                       onKeyUp={(e) => {
                           if (e.key === "Enter") {
                               props.onClick()
                           }
                       }}
                       onKeyDown={(e) => {
                           if (e.key === "Tab") {
                               console.log("#Tab")
                           }
                       }}
                       onChange={(e) => props.onChange(e.target.value)}
                       value={props.value} variant="plain" size="sm"></Input>
            }
            {!props.isEdit &&
                <Typography level="body-xs">{props.value}</Typography>
            }
        </td>
    )
}

export default function ClusterTable() {
    const [order, setOrder] = React.useState<Order>('desc');
    const [selected, setSelected] = React.useState<readonly string[]>([]);
    const [open, setOpen] = React.useState(false);
    const [rows, setRows] = useState<(IRow)[]>([])
    const [count, setCount] = useState<number>(0)
    const [page, setPage] = useState<number>(0)
    const [size, setSize] = useState<number>(0)
    const [needRefresh, setNeedRefresh] = useState<boolean>(true)


    useEffect(() => {

        fetch("/rest/clusters")
            .catch(error => console.error("# "  + error))

            // .then(data => {
            //
            // })
            // .then(json => {
            //         setCount(json.count)
            //         setPage(json.page)
            //         setSize(json.size)
            //         setRows(json.data)
            // }).catch(error => console.error("# "  + error))
        // const fetchData = async () => {
        //     const rest = await fetch("/rest/clusters")
        //     const json = await rest.json()
        //     return json
        // }

        // fetchData().catch(console.error).then(json => {
        //     setCount(json.count)
        //     setPage(json.page)
        //     setSize(json.size)
        //     setRows(json.data)
        // })
    }, [])

    const deleteRow = (index: number, value: IRow) => {
        if (value.code == null) {
            setRows(rows.filter((row, ind) => ind !== index))
        }
    }

    const addRow = (value: IRow) => {
        const requestOptions = {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify(value)
        };
        fetch("/rest/clusters", requestOptions).then(res => {
            if (res.ok) {
                // res.json().then(json => {
                //     const newRows = [json, ...rows]
                //     setRows(newRows)
                // })
            }
        })
    }

    const updateRow = (value: IRow) => {
        const requestOptions = {
            method: 'PUT',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify(value)
        };
        fetch("/rest/clusters", requestOptions).then(res => {
            if (res.ok) {
                // res.json().then(json => {
                //     const newRows = [json, ...rows]
                //     setRows(newRows)
                // })
            }
        })
    }

    const handleAddNewRow = () => {
        const newRow : IRow = {}
        const newRows = [newRow, ...rows]
        setRows(newRows)
        // const requestOptions = {
        //     method: 'POST',
        //     headers: {'Content-Type': 'application/json'},
        //     body: JSON.stringify(record)
        // };
        // fetch("/rest/clusters", requestOptions).then(res => {
        //     if (res.ok) {
        //         const newRows = [record, ...rows]
        //         setRows(newRows)
        //     }
        // })
    }


    const handleSubmit = (row: IRow) => {
        const requestOptions = {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify(row)
        };
        fetch("/rest/clusters", requestOptions).then(res => {
            if (res.ok) {
                setNeedRefresh(true)
            }
        })
    }

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

    const handleOnSelectSource = (row: IRow, value: string) => {
        row.source = value
    }
    const renderFilters = () => (
        <React.Fragment>
            <FormControl size="sm">
                <FormLabel>Status</FormLabel>
                <Select
                    size="sm"
                    placeholder="Filter by status"
                    slotProps={{button: {sx: {whiteSpace: 'nowrap'}}}}
                >
                    <Option value="paid">Paid</Option>
                    <Option value="pending">Pending</Option>
                    <Option value="refunded">Refunded</Option>
                    <Option value="cancelled">Cancelled</Option>
                </Select>
            </FormControl>

            <FormControl size="sm">
                <FormLabel>Category</FormLabel>
                <Select size="sm" placeholder="All">
                    <Option value="all">All</Option>
                    <Option value="refund">Refund</Option>
                    <Option value="purchase">Purchase</Option>
                    <Option value="debit">Debit</Option>
                </Select>
            </FormControl>

            <FormControl size="sm">
                <FormLabel>Customer</FormLabel>
                <Select size="sm" placeholder="All">
                    <Option value="all">All</Option>
                    <Option value="olivia">Olivia Rhye</Option>
                    <Option value="steve">Steve Hampton</Option>
                    <Option value="ciaran">Ciaran Murray</Option>
                    <Option value="marina">Marina Macdonald</Option>
                    <Option value="charles">Charles Fulton</Option>
                    <Option value="jay">Jay Hoper</Option>
                </Select>
            </FormControl>
        </React.Fragment>
    );
    return (
        <React.Fragment>
            <Sheet
                className="SearchAndFilters-mobile"
                sx={{
                    display: {
                        xs: 'flex',
                        sm: 'none',
                    },
                    my: 1,
                    gap: 1,
                }}
            >
                <Input
                    size="sm"
                    placeholder="Search"
                    startDecorator={<SearchIcon/>}
                    sx={{flexGrow: 1}}
                />
                <IconButton
                    size="sm"
                    variant="outlined"
                    color="neutral"
                    onClick={() => setOpen(true)}
                >
                    <FilterAltIcon/>
                </IconButton>
                <Modal open={open} onClose={() => setOpen(false)}>
                    <ModalDialog aria-labelledby="filter-modal" layout="fullscreen">
                        <ModalClose/>
                        <Typography id="filter-modal" level="h2">
                            Filters
                        </Typography>
                        <Divider sx={{my: 2}}/>
                        <Sheet sx={{display: 'flex', flexDirection: 'column', gap: 2}}>
                            {renderFilters()}
                            <Button color="primary" onClick={() => setOpen(false)}>
                                Submit
                            </Button>
                        </Sheet>
                    </ModalDialog>
                </Modal>
            </Sheet>
            <Box
                className="SearchAndFilters-tabletUp"
                sx={{
                    borderRadius: 'sm',
                    py: 2,
                    display: {
                        xs: 'none',
                        sm: 'flex',
                    },
                    flexWrap: 'wrap',
                    gap: 1.5,
                    '& > *': {
                        minWidth: {
                            xs: '120px',
                            md: '160px',
                        },
                    },
                }}
            >
                <FormControl sx={{flex: 1}} size="sm">
                    <FormLabel>Search for order</FormLabel>
                    <Input size="sm" placeholder="Search" startDecorator={<SearchIcon/>}/>
                </FormControl>
                {renderFilters()}
            </Box>
            <Box
                sx={{
                display: 'flex',
                justifyContent: 'flex-end'
            }}>
                <IconButton variant="outlined" size="sm" onClick={handleAddNewRow}>
                    <AddRoundedIcon/>
                </IconButton>
            </Box>
            <Sheet
                className="OrderTableContainer"
                variant="outlined"
                sx={{
                    borderRadius: 'sm',
                    '--TableCell-height': '40px',
                    // the number is the amount of the header rows.
                    '--TableHeader-height': 'calc(1 * var(--TableCell-height))',
                    '--Table-firstColumnWidth': '80px',
                    '--Table-lastColumnWidth': '128px',
                    // background needs to have transparency to show the scrolling shadows
                    '--TableRow-stripeBackground': 'rgba(0 0 0 / 0.04)',
                    '--TableRow-hoverBackground': 'rgba(0 0 0 / 0.08)',
                    overflow: 'auto',
                    background: (
                        theme,
                    ) => `linear-gradient(to right, ${theme.vars.palette.background.surface} 30%, rgba(255, 255, 255, 0)),
            linear-gradient(to right, rgba(255, 255, 255, 0), ${theme.vars.palette.background.surface} 70%) 0 100%,
            radial-gradient(
              farthest-side at 0 50%,
              rgba(0, 0, 0, 0.12),
              rgba(0, 0, 0, 0)
            ),
            radial-gradient(
                farthest-side at 100% 50%,
                rgba(0, 0, 0, 0.12),
                rgba(0, 0, 0, 0)
              )
              0 100%`,
                    backgroundSize:
                        '40px calc(100% - var(--TableCell-height)), 40px calc(100% - var(--TableCell-height)), 14px calc(100% - var(--TableCell-height)), 14px calc(100% - var(--TableCell-height))',
                    backgroundRepeat: 'no-repeat',
                    backgroundAttachment: 'local, local, scroll, scroll',
                    backgroundPosition:
                        '0 var(--TableCell-height), calc(100% - var(--Table-lastColumnWidth)) var(--TableCell-height), 0 var(--TableCell-height), calc(100% - var(--Table-lastColumnWidth)) var(--TableCell-height)',
                    backgroundColor: 'background.surface',
                }}
            >
                <Table
                    aria-labelledby="tableTitle"
                    stickyHeader
                    hoverRow
                    sx={{
                        height: '100%',
                        '--TableCell-headBackground': 'var(--joy-palette-background-level1)',
                        '--TableCell-headBackgroundLastColumn': 'var(--joy-palette-common-white)',
                        '& tr > th:last-child': {
                            position: 'sticky',
                            right: 0,
                            bgcolor: 'var(--TableCell-headBackground)',
                        },
                        '& tr > td:last-child': {
                            position: 'sticky',
                            right: 0,
                            maxWidth: '128px',
                            bgcolor: 'background.surface',
                        },
                }}
                >
                    <thead>
                    <tr>
                        <th style={{width: 140, maxWidth: 200, minWidth: 80, padding: '12px 6px'}}>
                            <Link
                                underline="none"
                                color="primary"
                                component="button"
                                onClick={() => setOrder(order === 'asc' ? 'desc' : 'asc')}
                                fontWeight="lg"
                                endDecorator={<ArrowDropDownIcon/>}
                                sx={{
                                    '& svg': {
                                        transition: '0.2s',
                                        transform:
                                            order === 'desc' ? 'rotate(0deg)' : 'rotate(180deg)',
                                    },
                                }}
                            >
                                Склад
                            </Link>
                        </th>
                        <th style={{width: 140, maxWidth: 200, minWidth: 80, padding: '12px 6px'}}>Кластер</th>
                        <th style={{width: 140, maxWidth: 200, minWidth: 80, padding: '12px 6px'}}>МП</th>
                        <th
                            aria-label="last"
                            style={{width: 'var(--Table-lastColumnWidth)', minWidth: 'var(--Table-lastColumnWidth)', maxWidth: 'var(--Table-lastColumnWidth)'}}
                        />
                    </tr>
                    </thead>
                    <tbody>
                    {rows.map((row, index) => (
                        <ClusterRow row={ row } id={ index } onCreate={addRow} onUpdate={updateRow} onDelete={deleteRow}/>
                    ))}
                    </tbody>
                </Table>
            </Sheet>
            <Box
                className="Pagination-laptopUp"
                sx={{
                    pt: 2,
                    gap: 2,
                    [`& .${iconButtonClasses.root}`]: {borderRadius: '50%'},
                    display: {
                        xs: 'flex',
                        md: 'flex',
                    },
                    alignItems: 'center',
                    justifyContent: 'flex-end'
                }}
            >
                <Box sx={{flex: 1}}/>
                {['1', '2', '3', '…', '8', '9', '10'].map((page) => (
                    <IconButton
                        key={page}
                        size="sm"
                        variant={Number(page) ? 'outlined' : 'plain'}
                        color="neutral"
                    >
                        {page}
                    </IconButton>
                ))}
                <Box sx={{flex: 1}}/>
                {/*<Box*/}
                {/*    sx={{*/}
                {/*        display: 'flex',*/}
                {/*        alignItems: 'center',*/}
                {/*        gap: 2,*/}
                {/*        justifyContent: 'flex-end',*/}
                {/*    }}*/}
                {/*>*/}
                <Typography textAlign="center" sx={{minWidth: 80}}>
                    {labelDisplayedRows({
                        from: rows.length === 0 ? 0 : (page - 1) * size + 1,
                        to: rows.length === 0 ? (page - 1) * size : (page - 1) * size + rows.length,
                        count: count,
                    })}
                </Typography>
                {/*</Box>*/}
            </Box>
        </React.Fragment>
    );
}