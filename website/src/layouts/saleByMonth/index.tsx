import * as React from 'react';
import { Fragment, useState } from 'react';
import Typography from '@mui/joy/Typography';
import Button from '@mui/joy/Button';
import DownloadRoundedIcon from '@mui/icons-material/DownloadRounded';
import Box from '@mui/joy/Box';
import SaleTable from 'layouts/saleByMonth/SaleTable';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import JoyDatePicker from 'components/JoyDatePicker';
import { LocalizationProvider } from '@mui/x-date-pickers';

import 'dayjs/locale/ru';
import dayjs from 'dayjs';
import { CustomSalesService } from 'layouts/saleByMonth/customSalesService';
import { OpenAPI } from 'client';

export function SaleByMonth() {
    const [date, setDate] = useState(dayjs().subtract(1, 'month'))
    const [isLoading, setIsLoading] = useState(false)

    const handleDownloadFile = () => {
        OpenAPI.TOKEN = sessionStorage.getItem('access_token') ?? ''
        CustomSalesService.getSalesByMonthReport(
            () => setIsLoading(true),
            () => setIsLoading(false),
            date.year(),
            date.month() + 1
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
                <Typography level="h3">Продажи за месяц</Typography>
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
                            openTo="year"
                            views={["year", "month"]}
                            size="sm"
                            sx={{
                                width: '170px'
                            }}
                            defaultValue={date}
                            minDate={dayjs(Date.parse('2023-01-01'))}
                            maxDate={dayjs().subtract(1, 'month')}
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
            <SaleTable year={date.year()} month={date.month() + 1}/>
        </Fragment>
    )
}
