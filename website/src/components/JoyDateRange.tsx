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
    dateFormat?: string | undefined;
}

export default function JoyDataRange(props: IProps): ReactElement {

    const dateFormat = props.dateFormat ?? 'DD-MM-YYYY'

    const [dateRange, setDateRange] = useState<Range[]>([
        {
            startDate: new Date(),
            endDate: undefined,
            key: 'selection'
        }
    ])
    const [anchorEl, setAnchorEl] = React.useState<HTMLButtonElement | null>(null);
    const inputRef = React.useRef(null);
    const handleClick = (event: React.MouseEvent<HTMLButtonElement>) => {
        //setAnchorEl(event.currentTarget);
        setAnchorEl(inputRef.current)
    };
    const handleClose = () => {
        setAnchorEl(null);
    };
    const open = Boolean(anchorEl);
    const id = open ? 'simple-popover' : undefined;
    const rangeDate = dateRange[0];
    const startDate = dayjs(rangeDate.startDate).format(dateFormat)
    const endDate = dayjs(rangeDate.endDate).format(dateFormat)
    const msg = startDate === endDate ? `${startDate}` : `${startDate} - ${endDate}`
    return (
        <Fragment>
            <div ref={inputRef}>
            <Input size='sm' aria-describedby={id} value={msg} endDecorator={<IconButton onClick={handleClick}><CalendarIcon/>
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
                        //onChange={item => setState([item.selection])}
                        //showSelectionPreview={true}
                        moveRangeOnFirstSelection={false}
                        months={2}
                        ranges={dateRange}
                        direction="horizontal"
                        locale={ru}
                        startDatePlaceholder='Начало периода'
                        endDatePlaceholder='Окончание периода'
                    />
                </div>
                <Box sx={{
                    pl: 2,
                    pr: 2,
                    pb: 2,
                    display: 'flex',
                    flexWrap: 'wrap',
                    justifyContent: 'end',
                }}>
                    <Button size='sm' variant="outlined" onClick={handleClose}>Готово</Button>
                </Box>
            </Popover>
        </Fragment>
    )
}
