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
import { OpenAPI, Sales, SalesesService } from 'client';
import YodaPagination from 'components/YodaPagination';
import { SetLogout } from 'context/actions';
import { useController } from 'context';


const initData: Sales = {
    items: [],
    count: 0
}

interface IProps {
    year: number
    month: number
}

export default function SaleTable({year, month}: IProps) {
    const {dispatch} = useController()
    const [selected, setSelected] = React.useState<readonly string[]>([]);
    const [open, setOpen] = React.useState(false);

    const [page, setPage] = React.useState(0);
    const [rowsPerPage, setRowsPerPage] = React.useState(25);
    const [data, setData] = useState<Sales>(initData)
    const {items, count} = data

    useEffect(() => {
        SalesesService.getSalesByMonthWithPagination(year, month, page + 1, rowsPerPage)
            .then(resp => {
                setData(resp)
            }).catch(error => {
            if (error.status === 401) {
                dispatch(SetLogout(error.body.description))
            }
        })
    }, [year, month, page, rowsPerPage]);

    const renderFilters = () => (
        <React.Fragment>
            <FormControl size="sm">
                <FormLabel>Страна</FormLabel>
                <Select
                    size="sm"
                    placeholder="Страна"
                    slotProps={{button: {sx: {whiteSpace: 'nowrap'}}}}
                >
                    <Option value="paid">Paid</Option>
                    <Option value="pending">Pending</Option>
                    <Option value="refunded">Refunded</Option>
                    <Option value="cancelled">Cancelled</Option>
                </Select>
            </FormControl>

            <FormControl size="sm">
                <FormLabel>Область</FormLabel>
                <Select size="sm" placeholder="Область">
                    <Option value="all">All</Option>
                    <Option value="refund">Refund</Option>
                    <Option value="purchase">Purchase</Option>
                    <Option value="debit">Debit</Option>
                </Select>
            </FormControl>

            <FormControl size="sm">
                <FormLabel>Регион</FormLabel>
                <Select size="sm" placeholder="Регион">
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

    const handleChangePage = (newPage: number) => {
        setPage(newPage);
    };

    const handleChangeRowsPerPage = (event: any, newValue: number | null) => {
        setRowsPerPage(parseInt(newValue!.toString(), 10));
        setPage(0);
    };

    const getLabelDisplayedRowsTo = () => {
        if (items.length === -1) {
            return (page + 1) * rowsPerPage;
        }
        return rowsPerPage === -1
            ? items.length
            : Math.min(items.length, (page + 1) * rowsPerPage);
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
                            Фильтр
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
                    <Input size="sm" placeholder="Поиск" startDecorator={<SearchIcon/>}/>
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
                        <th style={{width: "50px", padding: '12px 6px'}}>МП</th>
                        <th style={{width: "270px", padding: '12px 6px'}}>
                            Наименование
                        </th>
                        <th style={{width: "240px", padding: '12px 6px'}}>Адрес</th>
                        <th style={{width: "120px", padding: '12px 6px'}}>Штрих-код</th>
                        <th style={{width: '120px', padding: '12px 6px'}}>Код 1С</th>
                        <th style={{width: 120, padding: '12px 6px'}}>Артикл</th>
                        <th style={{width: '80px', padding: '12px 6px'}}>Кол-во</th>
                        <th style={{width: 160, padding: '12px 6px'}}>Сумма со скидкой</th>
                    </tr>
                    </thead>
                    <tbody>
                    {items.map((row) => (
                        <tr key={row.id}>
                            <td>
                                {row.source}
                            </td>
                            <td>
                                {row.name}
                            </td>
                            <td>
                                <Table size='sm'>
                                    <tbody>
                                    <tr>
                                        <td style={{width: 70}}>Старан</td>
                                        <td>{row.country}</td>
                                    </tr>
                                    {!!row.region &&
                                        <tr>
                                            <td>Область</td>
                                            <td>{row.region}</td>
                                        </tr>
                                    }
                                    {!!row.oblast &&
                                        <tr>
                                            <td>Округ</td>
                                            <td>{row.oblast}</td>
                                        </tr>
                                    }
                                    </tbody>
                                </Table>
                            </td>
                            <td>
                                {row.barcode}
                            </td>
                            <td>
                                {row.code1c}
                            </td>
                            <td>
                                {row.supplierArticle}
                            </td>
                            <td style={{textAlign: 'right'}}>
                                {row.quantity}
                            </td>
                            <td style={{textAlign: 'right'}}>
                                {row.total_price.toFixed(2)}
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
