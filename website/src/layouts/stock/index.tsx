import Box from "@mui/joy/Box";
import Typography from "@mui/joy/Typography";
import Button from "@mui/joy/Button";
import DownloadRoundedIcon from "@mui/icons-material/DownloadRounded";
import JoyDataGrid, {IColumn} from "../../components/JoyDataGrid";
import * as React from "react";
import {Fragment, useEffect, useState} from "react";
import {SetError, SetMenuActive} from "../../context/actions";
import {useController} from "../../context";
import {LocalizationProvider} from "@mui/x-date-pickers";
import {AdapterDayjs} from "@mui/x-date-pickers/AdapterDayjs";
import JoyDatePicker from "../../components/JoyDatePicker";
import dayjs from "dayjs";
import {DictionariesService, type StockFull, StocksService} from "../../client";

const columns: IColumn[] = [
    {
        field: 'stockDate',
        headerName: 'Дата',
        width: '100px',
        type:'date'
    }, {
        field: 'source',
        headerName: 'МП',
        width: '70px'
    }, {
        field: 'organization',
        headerName: 'Кабинет',
        width: '140px'
    }, {
        field: 'supplierArticle',
        headerName: 'Артикл',
        width: '130px'
    }, {
        field: 'barcode',
        headerName: 'Штрих код',
        width: '120px'
    }, {
        field: 'sku',
        headerName: 'Код МП',
        width: '100px'
    }, {
        field: 'item1C',
        headerName: 'Код 1C',
        width: '70px'
    }, {
        field: 'name',
        headerName: 'Наименование',
        width: '210px'
    }, {
        field: 'brand',
        headerName: 'Бренд',
        width: '100px'
    }, {
        field: 'warehouse',
        headerName: 'Склад',
        width: '120px',
        noWrap: true
    }, {
        field: 'quantity',
        headerName: 'Кол-во',
        width: '70px',
        textAlign: 'right'
    }, {
        field: 'price',
        headerName: 'Цена',
        width: '70px',
        textAlign: 'right'
    }, {
        field: 'priceWithDiscount',
        headerName: 'Цена со скидкой',
        width: '100px',
        textAlign: 'right'
    }
]

export default function Stocks() {
    const {dispatch} = useController()
    const [page, setPage] = useState(0)
    const [pageSize, setPageSize] = useState(25)
    const [total, setTotal] = useState(0)
    const [rows, setRows] = useState<StockFull[]>([])
    const [isLoading, setIsLoading] = useState(false)

    const [date, setDate] = useState(dayjs().subtract(1, 'day'))

    useEffect(() => {
        dispatch(SetMenuActive("menu-stocks-id"))
    }, []);

    function handleDownloadFile() {

    }

    const refreshRows = () => {
        setIsLoading(true)
        StocksService.getStocksWithPages(
            date.format('YYYY-MM-DD'),
            pageSize,
            page * pageSize
        ).then(req => {
                setTotal(req.count)
                setRows(req.items)
            })
            .catch(err => {
                dispatch(SetError(err.body.description))
            })
            .finally(() => setIsLoading(false))
    }

    useEffect(() => {
        refreshRows()
    }, [date, page, pageSize]);

    return (
        <Fragment>
            <Box
                sx={{
                    display: 'flex',
                    my: 1,
                    gap: 1,
                    flexDirection: {xs: 'column', sm: 'row'},
                    alignItems: {xs: 'start', sm: 'center'},
                    flexWrap: 'wrap',
                    justifyContent: 'space-between',
                }}
            >
                <Typography level="h3">Остатки на день</Typography>
                <Box
                    sx={{
                        display: 'flex',
                        my: 0.5,
                        gap: 2,
                        flexDirection: {xs: 'column', sm: 'row'},
                        alignItems: {xs: 'start', sm: 'center'},
                        flexWrap: 'wrap',
                        justifyContent: 'space-between',
                    }}
                >
                    <LocalizationProvider
                        adapterLocale='ru'
                        dateAdapter={AdapterDayjs}>
                        <JoyDatePicker
                            size="sm"
                            sx={{
                                width: '150px'
                            }}
                            defaultValue={date}
                            minDate={dayjs(Date.parse('2023-01-01'))}
                            maxDate={dayjs().subtract(1, 'day')}
                            onChange={(event) => setDate(event ?? date)}
                        />
                    </LocalizationProvider>

                    <Button
                        color="primary"
                        startDecorator={<DownloadRoundedIcon/>}
                        size="sm"
                        onClick={handleDownloadFile}
                        disabled={isLoading}
                    >
                        Скачать Excel
                    </Button>
                </Box>
            </Box>
            {/*{renderFilters()}*/}
            <JoyDataGrid
                showColumns
                columns={columns}
                rows={rows} size='sm'
                count={total}
                isLoading={isLoading}
                onPageChange={(value) => {
                    setPage(value)
                }}
                onPageSizeChange={(value) => {
                    setPageSize(value)
                }}
                onRefresh={refreshRows}
                page={page}
                pageSize={pageSize}/>
        </Fragment>
    )
}