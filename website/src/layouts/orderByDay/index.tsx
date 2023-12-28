import * as React from 'react';
import {Fragment, useEffect, useState} from 'react';
import Typography from '@mui/joy/Typography';
import Button from '@mui/joy/Button';
import DownloadRoundedIcon from '@mui/icons-material/DownloadRounded';
import Box from '@mui/joy/Box';
import OrderTable from './OrderTable';
import {AdapterDayjs} from '@mui/x-date-pickers/AdapterDayjs';
import JoyDatePicker from 'components/JoyDatePicker';
import {LocalizationProvider} from '@mui/x-date-pickers';

import 'dayjs/locale/ru';
import dayjs from 'dayjs';
import {CustomOrdersService} from 'layouts/orderByDay/customOrderService';
import {SetMenuActive} from "../../context/actions";
import {useController} from "../../context";

export function OrderByDay() {
    const {dispatch} = useController()
    const [date, setDate] = useState(dayjs().subtract(1, 'day'))
    const [isLoading, setIsLoading] = useState(false)

    useEffect(() => {
        dispatch(SetMenuActive("menu-orders-day-id"))
    }, []);

    const handleDownloadFile = () => {
        CustomOrdersService.getOrdersReport(
            () => setIsLoading(true),
            () => setIsLoading(false),
            date.format('YYYY-MM-DD')
        )
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
                <Typography level="h3">Заказы на день</Typography>
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
            <OrderTable date={date.format('YYYY-MM-DD')}/>
        </Fragment>
    )
}
