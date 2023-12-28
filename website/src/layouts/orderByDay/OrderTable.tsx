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
import { OpenAPI, Orders, OrdersService } from 'client';
import YodaPagination from 'components/YodaPagination';
import { useController } from 'context';
import { SetLogout } from 'context/actions';
import { Highlight } from '@mui/icons-material';

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

export default function OrderTable(props: { date: string }): React.JSX.Element {
    const {dispatch} = useController()
    const [selected, setSelected] = React.useState<readonly string[]>([]);
    const [open, setOpen] = React.useState(false);
    const [filter, setFilter] = React.useState(undefined);
    const [source, setSource] = React.useState<string | undefined>(undefined);

    const [page, setPage] = React.useState(0);
    const [rowsPerPage, setRowsPerPage] = React.useState(25);

    const [orders, setOrders] = useState<Orders>(initOrders)
    const {date} = props
    const {count, items} = orders
    const [isLoading, setLoading] = useState(false)

    const refreshData = () => {
        setLoading(true)
        OrdersService.getOrders(page + 1, rowsPerPage, date, filter, source)
            .then(result => {
                setOrders(result)
            }).catch(error => {
            if (error.status === 401) {
                dispatch(SetLogout(error.body.description))
            }
        }).finally(() => {
            setLoading(false)
        })
    }

    useEffect(() => {
        refreshData()
    }, [page, rowsPerPage, date, filter, source]);

    const handleOnSourceChange = (
        event: React.SyntheticEvent | null,
        newValue: string | null,
    ) => {
        setSource(newValue ?? "");
    }
    const renderFilters = () => (
        <React.Fragment>
            <FormControl size="sm">
                <FormLabel>МП</FormLabel>
                <Select
                    size="sm"
                    placeholder="Фильтр по МП"
                    slotProps={{button: {sx: {whiteSpace: 'nowrap'}}}}
                    onChange={handleOnSourceChange}
                    defaultValue=""
                >
                    <Option value="">Все</Option>
                    <Option value="OZON">OZON</Option>
                    <Option value="WB">WB</Option>
                </Select>
            </FormControl>
        </React.Fragment>
    );

    const handleChangePage = (newPage: number) => {
        setPage(newPage);
    };

    const handleChangeRowsPerPage = (event: any, newValue: number | null) => {
        setRowsPerPage(parseInt(newValue!.toString(), 10));
        setPage(0);
    };

    const handleChangeFilterText = (event: any) => {
        if (!event.target.value || event.target.value.length < 3) {
            return
        }
        setFilter(event.target.value);
    }

    const handleOnKeyDown = (event: any) => {
        if (event.key === 'Enter') {
            setFilter(event.target.value);
            event.preventDefault();
        }
    }

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
                    <FormLabel>Поиск</FormLabel>
                    <Input
                        size="sm"
                        placeholder="Поиск"
                        startDecorator={<SearchIcon/>}
                        onChange={handleChangeFilterText}
                        onKeyDown={handleOnKeyDown}
                    />
                </FormControl>
                {renderFilters()}
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
                        <th style={{width: "70px", padding: '12px 6px'}}>
                            МП
                        </th>
                        <th style={{width: 250, padding: '12px 6px'}}>Юр.лицо</th>
                        <th style={{width: 250, padding: '12px 6px'}}>Наименование</th>
                        <th style={{width: 140, padding: '12px 6px'}}>Штрих-код</th>
                        <th style={{width: 140, padding: '12px 6px'}}>Артикл</th>
                        <th style={{width: 140, padding: '12px 6px'}}>Код 1С</th>
                        <th style={{width: 140, padding: '12px 6px'}}>Бренд</th>
                        <th style={{width: 100, padding: '12px 6px'}}>SKU МП</th>
                        <th style={{width: 100, padding: '12px 6px'}}>Остаток</th>
                        <th style={{width: 140, padding: '12px 6px'}}>Сумма заказа</th>
                        <th style={{width: 130, padding: '12px 6px'}}>Кол-во заказа</th>
                    </tr>
                    </thead>
                    <tbody>
                    {!isLoading && orders.items.map((row) => (
                        <tr key={`row_${row.id}`}>
                            <td style={{textAlign: 'center', width: 120}}>
                                {row.source}
                            </td>
                            <td>
                                {row.org}
                            </td>
                            <td>
                                {row.name}
                            </td>
                            <td>
                                {row.barcode}
                            </td>
                            <td>
                                {row.supplierArticle}
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
