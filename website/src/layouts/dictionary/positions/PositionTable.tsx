import * as React from 'react';
import { useEffect, useState } from 'react';
import Box from '@mui/joy/Box';
import Button from '@mui/joy/Button';
import Divider from '@mui/joy/Divider';
import FormControl from '@mui/joy/FormControl';
import FormLabel from '@mui/joy/FormLabel';
import Input from '@mui/joy/Input';
import Modal from '@mui/joy/Modal';
import ModalDialog from '@mui/joy/ModalDialog';
import ModalClose from '@mui/joy/ModalClose';
import Select from '@mui/joy/Select';
import Option from '@mui/joy/Option';
import Table from '@mui/joy/Table';
import Sheet from '@mui/joy/Sheet';
import IconButton from '@mui/joy/IconButton';
import Typography from '@mui/joy/Typography';
// icons
import FilterAltIcon from '@mui/icons-material/FilterAlt';
import SearchIcon from '@mui/icons-material/Search';
import KeyboardArrowRightIcon from '@mui/icons-material/KeyboardArrowRight';
import KeyboardArrowLeftIcon from '@mui/icons-material/KeyboardArrowLeft';
import { Orders, OrdersService } from 'client';
import YodaPagination from 'components/YodaPagination';
import { Checkbox } from '@mui/joy';

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

export default function PositionTable(props: { date: string }): React.JSX.Element {
    const [selected, setSelected] = React.useState<readonly string[]>([]);
    const [open, setOpen] = React.useState(false);

    const [page, setPage] = React.useState(0);
    const [rowsPerPage, setRowsPerPage] = React.useState(25);

    const [orders, setOrders] = useState<Orders>(initOrders)
    const {date} = props
    const {count, items} = orders

    useEffect(() => {
        OrdersService.getOrders(page + 1, rowsPerPage, date)
            .then(result => {
                setOrders(result)
            })
    }, [page, rowsPerPage, date]);

    const handleChangePage = (newPage: number) => {
        setPage(newPage);
    };

    const handleChangeRowsPerPage = (event: any, newValue: number | null) => {
        setRowsPerPage(parseInt(newValue!.toString(), 10));
        setPage(0);
    };

    const getLabelDisplayedRowsTo = () => {
        if (orders.items.length === -1) {
            return (page + 1) * rowsPerPage;
        }
        return rowsPerPage === -1
            ? orders.items.length
            : Math.min(orders.items.length, (page + 1) * rowsPerPage);
    };

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
                    placeholder="Поиск"
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
                    <FormLabel>Поиск</FormLabel>
                    <Input size="sm" placeholder="Поиск" startDecorator={<SearchIcon/>}/>
                </FormControl>
            </Box>
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
                    height: '100%'
                }}
            >
                <Table
                    aria-labelledby="tableTitle"
                    stickyHeader
                    hoverRow
                    sx={{
                        '--TableCell-headBackground': 'var(--joy-palette-background-level1)',
                        '--Table-headerUnderlineThickness': '1px',
                        '--TableRow-hoverBackground': 'var(--joy-palette-background-level1)',
                        '--TableCell-paddingY': '4px',
                        '--TableCell-paddingX': '8px',
                        '& tr > th': {textAlign: 'center'}
                    }}
                >
                    <thead>
                    <tr>
                        <th style={{width: 48, textAlign: 'center', padding: '12px 6px' }} aria-label='empty'>
                            <Checkbox
                                size="sm"
                                indeterminate={
                                    selected.length > 0 && selected.length !== orders.items.length
                                }
                                checked={selected.length === orders.items.length}
                                onChange={(event) => {

                                }}
                                color={
                                    selected.length > 0 || selected.length === orders.items.length
                                        ? 'primary'
                                        : undefined
                                }
                                sx={{ verticalAlign: 'text-bottom' }}
                            />
                        </th>
                        <th style={{width: 250, padding: '12px 6px'}}>Наименование</th>
                        <th style={{width: 140, padding: '12px 6px'}}>Штрих-код</th>
                        <th style={{width: 140, padding: '12px 6px'}}>Код 1С</th>
                        <th style={{width: 140, padding: '12px 6px'}}>Бренд</th>
                        <th style={{width: 100, padding: '12px 6px'}}>SKU МП</th>
                        <th style={{width: 100, padding: '12px 6px'}}>Дата</th>
                    </tr>
                    </thead>
                    <tbody>
                    {orders.items.map((row) => (
                        <tr key={`row_${row.id}`}>
                            <td style={{textAlign: 'center', width: 120}}>
                                {row.source}
                            </td>
                            <td>
                                {row.name}
                            </td>
                            <td>
                                {row.barcode}
                            </td>
                            <td>
                                {row.code1c}
                            </td>
                            <td>
                                {row.brand}
                            </td>
                            <td>
                                {row.externalCode}
                            </td>
                            <td style={{textAlign: 'right'}}>
                                {row.balance}
                            </td>
                            <td style={{textAlign: 'right'}}>
                                {row.orderSum.toFixed(2)}
                            </td>
                            <td style={{textAlign: 'right'}}>
                                {row.orderedQuantity}
                            </td>
                        </tr>
                    ))}
                    </tbody>
                </Table>
            </Sheet>
            <YodaPagination page={page} count={count} rowsPerPage={rowsPerPage} pageLength={items.length}
                            onChangeRowsPerPage={handleChangeRowsPerPage}
                            onChangePage={handleChangePage}/>
        </React.Fragment>
    );
}
