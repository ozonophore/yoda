import React, { Fragment, useEffect, useState } from 'react';
import Box from '@mui/joy/Box';
import Typography from '@mui/joy/Typography';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { LocalizationProvider } from '@mui/x-date-pickers';
import JoyDateRange from 'components/JoyDateRange';
import DownloadRoundedIcon from '@mui/icons-material/DownloadRounded';
import Button from '@mui/joy/Button';
import JoyDataGrid, { IColumn } from 'components/JoyDataGrid';
import { OpenAPI, OrderProduct, OrdersService, ProductParams } from 'client';
import dayjs, { Dayjs } from 'dayjs';
import FormControl from '@mui/joy/FormControl';
import Input from '@mui/joy/Input';
import SearchIcon from '@mui/icons-material/Search';
import { Switch } from '@mui/joy';
import { useController } from 'context';
import { SetMenuActive } from 'context/actions';

type RangeDate = {
    dateFrom: string;
    dateTo: string;
}

const columnsModel : IColumn[] = [
    {
        field: 'orderDate',
        width: 100,
        headerName: 'Дата',
        headerTextAlign: 'center'
    },
    {
        field: 'source',
        width: 100,
        headerName: 'Площадка',
    },
    {
        field: 'org',
        width: 100,
        headerName: 'Кабинет',
        headerTextAlign: 'center',
    },
    {
        field: 'brand',
        width: 100,
        headerName: 'Бренд',
        headerTextAlign: 'center',
    },
    {
        field: 'code1c',
        width: 120,
        headerName: 'Код 1С',
        headerTextAlign: 'center',
    },
    {
        field: 'externalCode',
        width: 100,
        headerName: 'Код площадки',
        headerTextAlign: 'center',
        textAlign: 'right'
    },
    {
        field: 'name',
        minWith: 100,
        width: 200,
        headerName: 'Наименование',
        headerTextAlign: 'center',
    },
    {
        field: 'orderedQuantity',
        width: 70,
        headerName: 'Кол-во',
        headerTextAlign: 'center',
        textAlign: 'right'
    },
    {
        field: 'orderQuantityCanceled',
        width: 70,
        headerName: 'Кол-во отмен',
        headerTextAlign: 'center',
        textAlign: 'right'
    },
    {
        field: 'orderQuantityDelivered',
        width: 80,
        headerName: 'Кол-во доставлено',
        headerTextAlign: 'center',
        textAlign: 'right'
    }
]

export function OrderProductByDay(): React.JSX.Element {

    const [page, setPage] = useState(0)
    const [pageSize, setPageSize] = useState(25)
    const [total, setTotal] = useState(0)
    const [rows, setRows] = useState<OrderProduct[]>([])
    const [isLoading, setIsLoading] = useState(false)
    const [isDownload, setIsDownload] = useState(false)
    const [ rangeDate, setRangeDate ] = useState<RangeDate>({
        dateFrom: dayjs(Date()).format('YYYY-MM-DD'),
        dateTo: dayjs(Date()).format('YYYY-MM-DD')
    })
    const [filter, setFilter] = useState()
    const [filterVelue, setFilterValue] = useState()
    const [isGroup, setIsGroup] = useState(false)
    const [columns, setColumns] = useState<IColumn[]>(columnsModel)

    const {dispatch} = useController()
    useEffect(() => {
        dispatch(SetMenuActive("menu-orders-period-id"))
    }, []);

    function handleDownloadFile() {
        setIsDownload(true)
        OrdersService.exportOrdersProductToExcel(
            {
                dateFrom: rangeDate.dateFrom,
                dateTo: rangeDate.dateTo,
                filter: filterVelue,
                groupBy: isGroup ? ProductParams.groupBy.POSITION : undefined,
            }
        )
          .then((blob: Blob) => {
                //const blob: Blob = new Blob([data], {type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'});
                const fileURL = URL.createObjectURL(blob);
                const fileLink = document.createElement('a');
                fileLink.href = fileURL;
                fileLink.download = 'orders.xlsx';
                fileLink.click();
        }
        ).finally(() => setIsDownload(false))
    }

    function handlePagChange(page: number) {
        setPage(page)
    }

    function handleOnPageSizeChange(pageSize: number) {
        setPageSize(pageSize)
    }

    useEffect(() => {
        const cl = !isGroup ? columnsModel : columnsModel.slice(1)
        setColumns(cl)
        setIsLoading(true)
        OrdersService.getOrdersProduct({
            dateFrom: rangeDate.dateFrom,
            dateTo: rangeDate.dateTo,
            limit: pageSize,
            offset: pageSize * page,
            filter: filterVelue,
            groupBy: isGroup ? ProductParams.groupBy.POSITION : undefined,
        }).then(data => {
            setTotal(data.count)
            setRows(data.items)
        })
            .finally(() => {
            setIsLoading(false)
        })
    }, [page, pageSize, rangeDate, filterVelue, isGroup]);

    function handleRangeChange(dateFrom: Date, dateTo: Date) {
        setRangeDate({
            dateFrom: dayjs(dateFrom).format('YYYY-MM-DD'),
            dateTo: dayjs(dateTo).format('YYYY-MM-DD'),
        })
    }

    function handleChangeFilterText(event: any) {
        if (!event.target.value) {
            setFilter(undefined)
            return
        }
        setFilter(event.target.value);
    }

    function handleOnKeyDown(event: any) {
        if (event.key === 'Enter') {
            setFilterValue(filter);
            event.preventDefault();
        }
    }

    const renderFilters = () => (
        <React.Fragment>
            <FormControl size="sm" sx={{width: '100%'}}>
                <Input
                    fullWidth={true}
                    size="sm"
                    placeholder="Поиск"
                    startDecorator={<SearchIcon/>}
                    onChange={handleChangeFilterText}
                    onKeyDown={handleOnKeyDown}
                />
            </FormControl>
        </React.Fragment>
    );

    function handleChangeGroup(event: React.ChangeEvent<HTMLInputElement>) {
        setIsGroup(event.target.checked)
    }

    return (
        <Fragment>
            <Box
                sx={{
                    display: 'flex',
                    my: 0.5,
                    gap: 1,
                    flexDirection: {xs: 'column', sm: 'row'},
                    alignItems: {xs: 'start', sm: 'center'},
                    flexWrap: 'wrap',
                    justifyContent: 'space-between',
                }}
            >
                <Typography level="h3">Заказы за период</Typography>
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
                    <Switch onChange={handleChangeGroup} size='sm' title='Группировать' startDecorator='Группировка'/>
                    <LocalizationProvider
                        adapterLocale='ru'
                        dateAdapter={AdapterDayjs}>
                        <JoyDateRange width={250} onRangeChange={handleRangeChange} dateFrom={new Date()} dateTo={new Date()} />
                    </LocalizationProvider>
                    <Button
                        color="primary"
                        startDecorator={<DownloadRoundedIcon/>}
                        size="sm"
                        onClick={handleDownloadFile}
                        disabled={isLoading || isDownload}
                    >
                        Скачать Excel
                    </Button>
                </Box>
            </Box>
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
                {renderFilters()}
            </Box>
            <JoyDataGrid isLoading={isLoading} rows={rows} count={total} pageSize={pageSize} page={page}
                         onPageChange={handlePagChange} onPageSizeChange={handleOnPageSizeChange}
                         showColumns={true}
                         columns={ columns }/>
        </Fragment>
    )
}