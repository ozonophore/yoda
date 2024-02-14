import Box from "@mui/joy/Box";
import Typography from "@mui/joy/Typography";
import Button from "@mui/joy/Button";
import DownloadRoundedIcon from "@mui/icons-material/DownloadRounded";
import JoyDataGrid, { IColumn } from "../../components/JoyDataGrid";
import * as React from "react";
import { Fragment, useEffect, useState } from "react";
import { SetError, SetMenuActive } from "../../context/actions";
import { useController } from "../../context";
import dayjs from "dayjs";
import { type StockFull, StocksService } from "../../client";
import Select from '@mui/joy/Select';
import Option from '@mui/joy/Option';
import Input from '@mui/joy/Input';
import SearchIcon from '@mui/icons-material/Search';
import { KeyboardArrowDown } from '@mui/icons-material';
import { Chip } from '@mui/joy';
import PickerWithJoyField from 'components/PickerWithJoyField';

const columns: IColumn[] = [
    {
        field: 'stockDate',
        headerName: 'Дата',
        width: '100px',
        type: 'date'
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
    const {dispatch, state} = useController()

    const {marketplaces} = state.dicts
    const defaultMarketplaces = marketplaces.map(mp => mp.code)

    const [page, setPage] = useState(0)
    const [pageSize, setPageSize] = useState(25)
    const [total, setTotal] = useState(0)
    const [rows, setRows] = useState<StockFull[]>([])
    const [isLoading, setIsLoading] = useState(false)
    const [isDownload, setIsDownload] = useState(false)
    const [sources, setSources] = useState<string[]>(defaultMarketplaces)
    const [filter, setFilter] = useState<string | undefined>(undefined)

    const [date, setDate] = useState(dayjs())

    useEffect(() => {
        dispatch(SetMenuActive("menu-stocks-id"))
    }, []);

    function handleDownloadFile() {
        setIsDownload(true)
        StocksService.exportStocks(date.format("YYYY-MM-DD"), sources, filter)
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

    const refreshRows = () => {
        setIsLoading(true)
        StocksService.getStocksWithPages(
            date.format('YYYY-MM-DD'),
            pageSize,
            page * pageSize,
            sources,
            filter
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
    }, [date, page, pageSize, sources, filter]);

    const handleSourceChange = (
        event: React.SyntheticEvent | null,
        newValue: string[] | null,
    ) => {
        setSources(newValue ?? sources)
        setPage(0)
    };

    const handleOnKeyDown = (event: any) => {
        if (event.key === 'Enter') {
            setFilter(event.target.value)
            setPage(0)
            event.preventDefault();
        }
    }

    function renderFilters() {
        return <React.Fragment>
            <Box
                width='100%'
                sx={{
                    gap: 2,
                    display: "flex",
                }}
            >
                <Select
                    size="sm"
                    multiple
                    indicator={<KeyboardArrowDown/>}
                    placeholder="Площадка..."
                    defaultValue={defaultMarketplaces}
                    onChange={handleSourceChange}
                    renderValue={(selected) => (
                        <Box sx={{display: 'flex', gap: '0.25rem'}}>
                            {selected.map((selectedOption) => (
                                <Chip key={selectedOption.id} size="sm" variant="soft" color="primary">
                                    {selectedOption.label}
                                </Chip>
                            ))}
                        </Box>
                    )}
                    sx={{
                        minWidth: '270px',
                        width: '270px'
                    }}
                    slotProps={{
                        listbox: {
                            sx: {
                                width: '100%',
                            },
                        },
                    }}
                >
                    {
                        marketplaces.map(item => (
                            <Option key={item.code} value={item.code}>{item.shortName}</Option>
                        ))
                    }
                </Select>
                <Input
                    fullWidth={true}
                    size="sm"
                    placeholder="Поиск"
                    startDecorator={<SearchIcon/>}
                    // onChange={handleChangeFilterText}
                    onKeyDown={handleOnKeyDown}
                />
            </Box>
        </React.Fragment>
    }

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
                    <PickerWithJoyField
                        defaultValue={date}
                        minDate={dayjs(Date.parse('2023-01-01'))}
                        maxDate={dayjs()}
                        onChange={(event) => {
                            setDate(event ?? date)
                            setPage(0)
                        }}
                    />

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