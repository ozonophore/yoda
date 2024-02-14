import * as React from 'react';
import { Fragment, useEffect, useState } from 'react';
import Typography from '@mui/joy/Typography';
import Button from '@mui/joy/Button';
import DownloadRoundedIcon from '@mui/icons-material/DownloadRounded';
import Box from '@mui/joy/Box';
import OrderTable from './OrderTable';

import 'dayjs/locale/ru';
import dayjs from 'dayjs';
import { CustomOrdersService } from 'layouts/orderByDay/customOrderService';
import { SetMenuActive } from "../../context/actions";
import { useController } from "../../context";
import PickerWithJoyField from 'components/PickerWithJoyField';

export function OrderByDay() {
    const {dispatch} = useController()
    const [date, setDate] = useState(dayjs())
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
                    <PickerWithJoyField
                        defaultValue={date}
                        minDate={dayjs(Date.parse('2023-01-01'))}
                        maxDate={dayjs()}
                        onChange={(event) => setDate(event ?? date)}
                    />
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
