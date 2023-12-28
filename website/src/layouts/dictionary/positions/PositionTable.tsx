import * as React from 'react';
import { Fragment, useEffect, useState } from 'react';
import FormControl from '@mui/joy/FormControl';
import Input from '@mui/joy/Input';
// icons
import SearchIcon from '@mui/icons-material/Search';
import { DictionariesService, DictPosition } from 'client';
import JoyDataGrid, { IColumn } from 'components/JoyDataGrid';

import Select from '@mui/joy/Select';
import Option from '@mui/joy/Option';
import { Box, Chip } from '@mui/joy';
import { KeyboardArrowDown, Logout } from '@mui/icons-material';
import { useController } from 'context';
import { SetLogout } from 'context/actions';

const columnsModel: IColumn[] = [
    {
        field: 'name',
        width: 170,
        headerName: 'Наименование',
        headerTextAlign: 'center'
    },
    {
        field: 'marketplace',
        width: 60,
        headerName: 'Площадка',
    },
    {
        field: 'org',
        width: 100,
        headerName: 'Кабинет',
        headerTextAlign: 'center',
    },
    {
        field: 'barcode',
        width: 80,
        headerName: 'Штрихкод',
        headerTextAlign: 'center',
    },
    {
        field: 'code1c',
        width: 100,
        headerName: 'Код 1С',
        headerTextAlign: 'center',
    },
]

export default function PositionTable(props: {
    date: string
}): React.JSX.Element {
    const { dispatch } = useController()

    const [page, setPage] = useState(0)
    const [pageSize, setPageSize] = useState(25)
    const [total, setTotal] = useState(0)
    const [rows, setRows] = useState<DictPosition[]>([])
    const [isLoading, setIsLoading] = useState(false)
    const [filter, setFilter] = useState()
    const [filterVelue, setFilterValue] = useState()
    const [isGroup, setIsGroup] = useState(false)
    const [columns, setColumns] = useState<IColumn[]>(columnsModel)
    const [source, setSource] = useState<string[]>(["WB", "OZON"])

    function handleDownloadFile() {

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
        const req = {
            limit: pageSize,
            offset: pageSize * page,
            source: source,
            filter: filter,
        }
        DictionariesService.getPositions(req)
            .then(data => {
                setTotal(data.count)
                setRows(data.items)
            })
            .catch(reason => {
                dispatch(SetLogout(reason.description))
            })
            .finally(() => {
                setIsLoading(false)
            })
    }, [page, pageSize, filterVelue, isGroup, source]);


    function handleChangeFilterText(event: any) {
        if (!event.target.value) {
            setFilter(undefined)
            return
        }
        setPage(0)
        setFilter(event.target.value);
    }

    function handleOnKeyDown(event: any) {
        if (event.key === 'Enter') {
            setFilterValue(filter);
            event.preventDefault();
        }
    }

    const handleSourceChange = (
        event: React.SyntheticEvent | null,
        newValue: string[] | null,
    ) => {
        setSource(newValue ?? source)
    };

    const renderFilters = () => (
        <React.Fragment>
            <FormControl size="sm" sx={{width: '100%'}}>
                <Box
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
                        defaultValue={['WB', 'OZON']}
                        onChange={handleSourceChange}
                        renderValue={(selected) => (
                            <Box sx={{display: 'flex', gap: '0.25rem'}}>
                                {selected.map((selectedOption) => (
                                    <Chip size="sm" variant="soft" color="primary">
                                        {selectedOption.label}
                                    </Chip>
                                ))}
                            </Box>
                        )}
                        sx={{
                            minWidth: '140px',
                            width: '140px'
                        }}
                        slotProps={{
                            listbox: {
                                sx: {
                                    width: '100%',
                                },
                            },
                        }}
                    >
                        <Option value="WB">WB</Option>
                        <Option value="OZON">OZON</Option>
                    </Select>
                    <Input
                        fullWidth={true}
                        size="sm"
                        placeholder="Поиск"
                        startDecorator={<SearchIcon/>}
                        onChange={handleChangeFilterText}
                        onKeyDown={handleOnKeyDown}
                    />
                </Box>
            </FormControl>
        </React.Fragment>
    );

    return (
        <Fragment>
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
                         columns={columns}/>
        </Fragment>
    )
}
