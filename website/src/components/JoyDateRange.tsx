import React, { Fragment, ReactElement, useState } from 'react';
import Input from '@mui/joy/Input';
import IconButton from '@mui/joy/IconButton';
import { CalendarIcon } from '@mui/x-date-pickers';
import { Popover } from '@mui/material';
import { DateRange, Range, RangeKeyDict } from 'react-date-range';
import ru from 'date-fns/locale/ru';
import Box from '@mui/joy/Box';
import Button from '@mui/joy/Button';
import dayjs from 'dayjs';

import 'react-date-range/dist/styles.css';
import 'react-date-range/dist/theme/default.css';

interface IProps {
    dateFrom: Date
    dateTo: Date
    width?: number | string | undefined
    dateFormat?: string | undefined
    onRangeChange: (startDate: Date, endDate: Date) => void
}

function getMessage(startDate: Date, endDate: Date, dateFormat: string): string {
    const startDateStr = dayjs(startDate).format(dateFormat)
    const endDateStr = dayjs(endDate).format(dateFormat)
    return startDate.getDate() === endDate.getDate() ? `${startDateStr}` : `${startDateStr} - ${endDateStr}`
}

export default function JoyDateRange({width, onRangeChange, dateFrom, dateTo, dateFormat = 'DD.MM.YYYY'}: IProps): ReactElement {

    const [dateRange, setDateRange] = useState<Range[]>([
        {
            startDate: dateFrom,
            endDate: dateTo,
            key: 'selection'
        }
    ])
    const [anchorEl, setAnchorEl] = React.useState<HTMLButtonElement | null>(null);
    const inputRef = React.useRef(null);
    const [inputValue, setInputValue] = useState(getMessage(dateFrom, dateTo, dateFormat))
    const handleClick = () => {
        setAnchorEl(inputRef.current)
    };
    const handleClose = () => {
        setAnchorEl(null);
    };
    const open = Boolean(anchorEl);
    const id = open ? 'simple-popover' : undefined;

    function handleSubmit() {
        const {startDate, endDate} = dateRange[0]
        setInputValue(getMessage(startDate ?? dateFrom, endDate ?? dateTo, dateFormat))
        onRangeChange(startDate ?? new Date(), endDate ?? new Date())
        handleClose()
    }

    function handleSetToday() {
        const today = new Date()
        setDateRange([{
            startDate: today,
            endDate: today,
            key: 'selection'
        }])
    }

    function handleSetYesterday() {
        const yesterday = new Date()
        yesterday.setDate(yesterday.getDate() - 1)
        setDateRange([{
            startDate: yesterday,
            endDate: yesterday,
            key: 'selection'
        }])
    }

    return (
        <Fragment>
            <div ref={inputRef} style={{width: width}}>
                <Input size='sm' aria-describedby={id} value={inputValue} onClick={handleClick} onKeyUp={handleClick}
                       endDecorator={<IconButton onClick={handleClick}><CalendarIcon/>
                       </IconButton>}/></div>
            <Popover id={id}
                     open={open}
                     anchorEl={anchorEl}
                     onClose={handleClose}
                     anchorOrigin={{
                         vertical: 'bottom',
                         horizontal: 'right',
                     }}
                     transformOrigin={{
                         vertical: 'top',
                         horizontal: 'right',
                     }}
            >
                <div>
                    <DateRange
                        onChange={(item: RangeKeyDict) => {
                            setDateRange([item.selection])
                        }}
                        moveRangeOnFirstSelection={false}
                        showPreview={true}
                        months={2}
                        ranges={dateRange}
                        direction="horizontal"
                        locale={ru}
                        startDatePlaceholder='Начало периода'
                        endDatePlaceholder='Окончание периода'
                        classNames={{
                            dayToday: 'Сегодня'
                        }}
                    />
                </div>
                <Box sx={{
                    pl: 2,
                    pr: 2,
                    pb: 2,
                    display: 'flex',
                    flexWrap: 'wrap',
                    justifyContent: 'space-between',
                }}>
                    <Box sx={{
                        gap: 2,
                        display: 'flex',
                        flexWrap: 'wrap',
                        justifyContent: 'space-between'}}>
                        <Button size='sm' onClick={handleSetToday}>Сегодня</Button>
                        <Button size='sm' onClick={handleSetYesterday}>Вчера</Button>
                    </Box>
                    <Button size='sm' variant="outlined" onClick={handleSubmit}>Готово</Button>
                </Box>
            </Popover>
        </Fragment>
    )
}
