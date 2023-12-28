import Box from '@mui/joy/Box';
import Select from '@mui/joy/Select';
import Option from '@mui/joy/Option';
import Typography from '@mui/joy/Typography';
import IconButton from '@mui/joy/IconButton';
import KeyboardArrowLeftIcon from '@mui/icons-material/KeyboardArrowLeft';
import KeyboardArrowRightIcon from '@mui/icons-material/KeyboardArrowRight';
import {Grid} from '@mui/joy';
import * as React from 'react';

interface IProps {
    page: number
    count: number
    rowsPerPage: number
    pageLength: number
    onChangeRowsPerPage: (event: any, newValue: number | null) => void
    onChangePage: (page: number) => void
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
    return `${from}–${to} из ${count !== -1 ? count : `больше чем ${to}`}`;
}

const YodaPagination = (props: IProps): React.JSX.Element => {

    const getLabelDisplayedRowsTo = () => {
        if (props.pageLength === -1) {
            return (props.page + 1) * props.rowsPerPage;
        }
        return props.rowsPerPage === -1
            ? props.pageLength
            : Math.min(props.pageLength, (props.page + 1) * props.rowsPerPage);
    };

    return (
        <Grid container direction='row'
              justifyContent='flex-end'
              alignItems='flex-end'>
            <Box
                sx={{
                    display: 'flex',
                    alignItems: 'center',
                    gap: 1,
                    justifyContent: 'flex-end',
                }}
            >
                <Typography level="body-sm">Строк на странице:</Typography>
                <Select size='sm' onChange={props.onChangeRowsPerPage} value={props.rowsPerPage}>
                    <Option value={25}>25</Option>
                    <Option value={30}>30</Option>
                    <Option value={50}>50</Option>
                </Select>
                <Typography
                    level="body-sm"
                    textAlign="center"
                    sx={{
                        minWidth: 100,
                    }}>
                    {labelDisplayedRows({
                        from: props.pageLength === 0 ? 0 : props.page * props.rowsPerPage + 1,
                        to: getLabelDisplayedRowsTo(),
                        count: props.count,
                    })}
                </Typography>
                <Box sx={{display: 'flex', gap: 1}}>
                    <IconButton
                        size="sm"
                        color="neutral"
                        variant="outlined"
                        disabled={props.page === 0}
                        onClick={() => props.onChangePage(props.page - 1)}
                        sx={{bgcolor: 'background.surface'}}
                    >
                        <KeyboardArrowLeftIcon/>
                    </IconButton>
                    <IconButton
                        size="sm"
                        color="neutral"
                        variant="outlined"
                        disabled={
                            props.pageLength !== 0
                                ? props.page >= Math.ceil(props.count / props.rowsPerPage) - 1
                                : true
                        }
                        onClick={() => props.onChangePage(props.page + 1)}
                        sx={{bgcolor: 'background.surface'}}
                    >
                        <KeyboardArrowRightIcon/>
                    </IconButton>
                </Box>
            </Box>
        </Grid>
    )
}

export default YodaPagination