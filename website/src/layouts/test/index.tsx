import React, { Fragment, useState } from 'react';
import JoyDataGrid, { IColumn } from 'components/JoyDataGrid';
import Typography from "@mui/joy/Typography";
import Input from "@mui/joy/Input";
import Select from '@mui/joy/Select';
import Option from '@mui/joy/Option';
import Box from '@mui/joy/Box';
import { Chip } from '@mui/joy';
import dayjs from 'dayjs';
import { experimental_extendTheme as materialExtendTheme, } from '@mui/material/styles';
import PickerWithJoyField from 'components/PickerWithJoyField';
// import DatePicker from 'react-datepicker';
// import 'react-datepicker/dist/react-datepicker.css';
// import './style.css';
// import ru from 'date-fns/locale/ru'


const columns: IColumn[] = [
    {
        field: "name",
        textAlign: 'right',
        editable: true,
    }, {
        field: "value",
        type: "int",
        editable: true,
    }
]

const rows = [
    {
        name: 'Name 1',
        value: 456
    }, {
        name: 'Name 2',
        value: 123,
    }
]

function SelectMultipleAppearance() {
    return (
        <Select
            multiple
            defaultValue={['dog', 'cat']}
            renderValue={(selected) => (
                <Box sx={{display: 'flex', gap: '0.25rem'}}>
                    {selected.map((selectedOption) => (
                        <Chip key={selectedOption.id} variant="soft" color="primary">
                            {selectedOption.label}
                        </Chip>
                    ))}
                </Box>
            )}
            sx={{
                minWidth: '15rem',
            }}
            slotProps={{
                listbox: {
                    sx: {
                        width: '100%',
                    },
                },
            }}
        >
            <Option key="dog" value="dog">Dog</Option>
            <Option key="cat" value="cat">Cat</Option>
            <Option key="fish" value="fish">Fish</Option>
            <Option key="bird" value="bird">Bird</Option>
        </Select>
    );
}

const materialTheme = materialExtendTheme();

export function Test() {

    const [lines, setLines] = useState(0)
    const [date, setDate] = useState(dayjs().subtract(1, 'day'))

    const style = {
        overflow: "hidden",
        textOverflow: "ellipsis",
        display: "-webkit-box",
        WebkitLineClamp: lines,
        WebkitBoxOrient: "vertical",
    }

    return <Fragment>
        <div>
            {/*<DatePicker popperClassName="calendar-popout" onChange={(date) => console.log(date)}*/}
            {/*locale={ru}*/}
            {/*/>*/}
            <PickerWithJoyField></PickerWithJoyField>
            <Input></Input>
            <SelectMultipleAppearance/>
            <Input type='number' onChange={(value) => {
                setLines(Number(value.target.value))
            }}></Input>
            <Typography noWrap={false} sx={lines !== 0 ? style : {}}>
                Для вас мы собрали самые известные и проверенные временем сказки для детей. Здесь размещены русские
                народные сказки и авторские сказки, которые точно стоит прочитать ребенку. Детские сказки этого раздела
                подходят абсолютно всем ребятам: подобраны сказки для самых маленьких и для школьников. Некоторые
                произведения Вы найдете только у нас, в оригинальном изложении!
            </Typography>
        </div>
        <JoyDataGrid columns={columns} rows={rows} page={0} pageSize={2} count={2} onPageChange={() => {
        }}
                     onPageSizeChange={() => {
                     }} isLoading={false} showColumns={true}/>
    </Fragment>
}