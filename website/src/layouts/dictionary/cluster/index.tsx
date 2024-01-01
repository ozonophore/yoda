import * as React from 'react';
import {Fragment, useEffect, useState} from 'react';
import Typography from '@mui/joy/Typography';
import Button from '@mui/joy/Button';
import DownloadRoundedIcon from '@mui/icons-material/DownloadRounded';
import FilterAltRoundedIcon from '@mui/icons-material/FilterAltRounded';
import Box from '@mui/joy/Box';
import JoyDataGrid, {IColumn} from 'components/JoyDataGrid';
import Select from '@mui/joy/Select';
import Option from '@mui/joy/Option';
import Input from '@mui/joy/Input';
import {Autocomplete, Chip, CircularProgress} from '@mui/joy';
import {useController} from 'context';
import {SetError, SetMenuActive} from 'context/actions';
import {DictionariesService, Warehouse} from 'client';
import IconButton from '@mui/joy/IconButton';
import deepEqual from "deep-equal";

interface IFilter {
    cluster: string | undefined
    code: string | undefined,
    sources: string[]
}

function ClusterAutocomplete({value, onChange}: {
    value: string,
    onChange: (value: string | null) => void
}): React.JSX.Element {
    const [open, setOpen] = React.useState(false);
    const [options, setOptions] = React.useState<string[]>([]);
    const loading = open && options.length === 0;

    useEffect(() => {
        let active = true;

        if (!loading) {
            return undefined;
        }
        DictionariesService.getClusters()
            .then(data => {
                if (active) {
                    setOptions(data)
                }
            })

        return () => {
            active = false;
        };
    }, [loading]);

    useEffect(() => {
        if (!open) {
            setOptions([]);
        }
    }, [open]);

    return <Autocomplete
        size='sm'
        variant="plain" sx={{
            border: 0,
            background: 'transparent',
            color: 'var(--joy-palette-text-tertiary, var(--joy-palette-neutral-600, #555E68))'
    }}
        open={open}
        onChange={(event, newInputValue) => onChange(newInputValue)}
        onInputChange={(event, newInputValue) => {
            onChange(newInputValue);
        }}
        onOpen={() => {
            setOpen(true);
        }}
        onClose={() => {
            setOpen(false);
        }}
        value={value}
        options={options}
        loading={loading}
        freeSolo
        endDecorator={
            loading ? (
                <CircularProgress
                    size="sm" sx={{bgcolor: 'background.surface'}}/>
            ) : null
        }
    />
}

const columns: IColumn[] = [
    {
        field: 'code',
        minWith: '140px',
        headerName: 'Код склада',
        noWrap: true,
    }, {
        field: 'cluster',
        headerName: 'Кластер',
        width: '280px',
        editable: true,
        renderComponent: ClusterAutocomplete
    }, {
        field: 'source',
        width: '80px',
        headerName: 'Искочник'
    }
]

export function Cluster() {
    const {dispatch} = useController()
    const [sources, setSources] = useState(['WB', 'OZON'])
    const [code, setCode] = useState<string>()
    const [cluster, setCluster] = useState<string>()
    const [filter, setFilter] = useState<IFilter>({
        cluster: undefined,
        code: undefined,
        sources: ['WB', 'OZON']
    })

    const [page, setPage] = useState(0)
    const [pageSize, setPageSize] = useState(25)
    const [total, setTotal] = useState(0)
    const [rows, setRows] = useState<Warehouse[]>([])
    const [isLoading, setIsLoading] = useState(false)
    const [isDownload, setIsDownload] = useState(false)

    const refreshRows = () => {
        setIsLoading(true)
        DictionariesService.getWarehouses(filter.sources, pageSize, pageSize * page, filter.cluster, filter.code)
            .then(req => {
                setTotal(req.count)
                setRows(req.items)
            })
            .catch(err => {
                dispatch(SetError(err.body.description))
            })
            .finally(() => setIsLoading(false))
    }

    useEffect(() => {
        dispatch(SetMenuActive("menu-dict-clusters-id"))
    }, []);

    useEffect(() => {
        refreshRows()
    }, [page, pageSize]);

    useEffect(() => {
        setPage(0)
        refreshRows()
    }, [filter]);

    const renderFilters = () => (
        <Box sx={{
            display: 'flex',
            gap: 1,
            flexDirection: 'row',
        }}>
            <Box width='100%' sx={{display: 'flex', gap: 1, flexDirection: 'row'}}>
                <Input
                    fullWidth={true}
                    value={code}
                    onChange={(e) => setCode(e.target.value)}
                    size='sm' placeholder='Код склада'/>
                <Input
                    value={cluster}
                    fullWidth={true}
                    onChange={(e) => setCluster(e.target.value)}
                    size='sm' placeholder='Кластер'/>
                <Select value={sources}
                        onChange={(e,
                                   newValue) => {
                            setSources(newValue)
                        }}
                        sx={{width: '350px'}} size="sm"
                        multiple
                        defaultValue={['WB', 'OZON']}
                        renderValue={(selected) => (
                            <Box sx={{display: 'flex', gap: '0.25rem'}}>
                                {selected.map((selectedOption, index) => (
                                    <Chip key={`source_${index}`} variant="soft" color="primary">
                                        {selectedOption.label}
                                    </Chip>
                                ))}
                            </Box>
                        )}
                        placeholder="Источник">
                    <Option value="WB">WB</Option>
                    <Option value="OZON">OZON</Option>
                </Select>
            </Box>
            <Box alignContent='center'>
                <IconButton size='sm' variant='outlined' onClick={() => {
                    //setPage(0)
                    setFilter({
                        sources: sources,
                        code: code,
                        cluster: cluster
                    })
                }}>
                    <FilterAltRoundedIcon/>
                </IconButton>
            </Box>
        </Box>
    );

    function handleDownloadFile() {
        setIsDownload(true)
        DictionariesService.exportWarehouses(
            filter.sources,
            filter.cluster,
            filter.code,
        ).then((blob: Blob) => {
                    const fileURL = URL.createObjectURL(blob);
                    const fileLink = document.createElement('a');
                    fileLink.href = fileURL;
                    fileLink.download = 'warehouses.xlsx';
                    fileLink.click();
                }
            ).finally(() => setIsDownload(false))
    }

    function handleOnSave(oldValue: Warehouse, newValue: Warehouse) {
        //setIsLoading(true)
         DictionariesService.updateWarehouse(
            newValue
        ).then(data => {
            const newRows = rows.map(item => {
                if (deepEqual(item, oldValue)) {
                    return newValue
                } else {
                    return item
                }
            })
            setRows(newRows)
         })
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
                <Typography level="h3">Кластеры</Typography>
                <Button
                    color="primary"
                    startDecorator={<DownloadRoundedIcon/>}
                    size="sm"
                    disabled={isLoading}
                    onClick={handleDownloadFile}
                >
                    Скачать Excel
                </Button>
            </Box>
            {renderFilters()}
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
                page={page}
                onSave={handleOnSave}
                pageSize={pageSize}
                onRefresh={refreshRows}
            />
        </Fragment>
    )
}