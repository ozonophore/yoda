import * as React from 'react';
import {Fragment, useEffect, useState} from 'react';
import Typography from '@mui/joy/Typography';
import Button from '@mui/joy/Button';
import DownloadRoundedIcon from '@mui/icons-material/DownloadRounded';
import Box from '@mui/joy/Box';
import PositionTable from 'layouts/dictionary/positions/PositionTable';

import 'dayjs/locale/ru';
import dayjs from 'dayjs';
import {SetMenuActive} from "../../../context/actions";
import {useController} from "../../../context";

export function DictPositions() {
    const {dispatch} = useController()
    const [date, setDate] = useState(dayjs())
    useEffect(() => {
        dispatch(SetMenuActive("menu-dict-item1c-id"))
    }, []);
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
                <Typography level="h3">Позиции</Typography>
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
                    <a href={`/rest/orders/report?date=${date.format('YYYY-MM-DD')}`} target="_blank" download>
                        <Button
                            color="primary"
                            startDecorator={<DownloadRoundedIcon/>}
                            size="sm"
                        >
                            Скачать Excel
                        </Button>
                    </a>
                </Box>
            </Box>
            <PositionTable date={date.format('YYYY-MM-DD')}/>
        </Fragment>
    )
}
