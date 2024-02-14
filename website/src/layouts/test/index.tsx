import React, { Fragment, useMemo, useState } from 'react';
import { IColumn } from 'components/JoyDataGrid';
import Typography from "@mui/joy/Typography";
import Input from "@mui/joy/Input";
import Select from '@mui/joy/Select';
import Option from '@mui/joy/Option';
import Box from '@mui/joy/Box';
import { Chip } from '@mui/joy';
import dayjs from 'dayjs';
import { experimental_extendTheme as materialExtendTheme, } from '@mui/material/styles';
import PickerWithJoyField from 'components/PickerWithJoyField';
import { MaterialReactTable, type MRT_ColumnDef, useMaterialReactTable, } from 'material-react-table';
import Sheet from '@mui/joy/Sheet';
import List from '@mui/joy/List';
import ListItem from '@mui/joy/ListItem';
import ListDivider from '@mui/joy/ListDivider';
import { Check } from '@mui/icons-material';
import ListItemDecorator from '@mui/joy/ListItemDecorator';
import CheckRoundedIcon from '@mui/icons-material/CheckRounded';
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

type Person = {
    name: {
        firstName: string;
        lastName: string;
    };
    address: string;
    city: string;
    state: string;
};
const data: Person[] = [
    {
        name: {
            firstName: 'John',
            lastName: 'Doe',
        },
        address: '261 Erdman Ford',
        city: 'East Daphne',
        state: 'Kentucky',
    },
    {
        name: {
            firstName: 'Jane',
            lastName: 'Doe',
        },
        address: '769 Dominic Grove',
        city: 'Columbus',
        state: 'Ohio',
    },
    {
        name: {
            firstName: 'Joe',
            lastName: 'Doe',
        },
        address: '566 Brakus Inlet',
        city: 'South Linda',
        state: 'West Virginia',
    },
    {
        name: {
            firstName: 'Kevin',
            lastName: 'Vandy',
        },
        address: '722 Emie Stream',
        city: 'Lincoln',
        state: 'Nebraska',
    },
    {
        name: {
            firstName: 'Joshua',
            lastName: 'Rolluffs',
        },
        address: '32188 Larkin Turnpike',
        city: 'Omaha',
        state: 'Nebraska',
    },
];

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

    const columns2 = useMemo<MRT_ColumnDef<Person>[]>(
        () => [
            {
                accessorKey: 'name.firstName', //access nested data with dot notation
                header: 'First Name',
                size: 150,
            },
            {
                accessorKey: 'name.lastName',
                header: 'Last Name',
                size: 150,
            },
            {
                accessorKey: 'address', //normal accessorKey
                header: 'Address',
                size: 200,
            },
            {
                accessorKey: 'city',
                header: 'City',
                size: 150,
            },
            {
                accessorKey: 'state',
                header: 'State',
                size: 150,
            },
        ],
        [],
    );
    let tableSx = {
        '--TableCell-headBackground': 'var(--joy-palette-background-level1)',
        '--Table-headerUnderlineThickness': '1px',
        '--TableRow-hoverBackground': 'var(--joy-palette-background-level1)',
        '--TableCell-paddingY': '4px',
        '--TableCell-paddingX': '8px',
        '& tr > th': {textAlign: 'center'},
    }
    const table = useMaterialReactTable({
        columns: columns2,
        data, //data must be memoized or stable (useState, useMemo, defined outside of this component, etc.)

        muiTableHeadCellProps: {
            //no useTheme hook needed, just use the `sx` prop with the theme callback
            sx: (theme) => ({
                //width: column.width,
                padding: '12px 6px',
                //textAlign: column.headerTextAlign,
                // minWidth: column.minWith,
                background: 'var(--joy-palette-background-level1)',
                whiteSpace: 'normal',
            }),
        },
        muiTableContainerProps: {
            sx: {
                maxHeight: '500px',
                height: '100%'
            }
        }
    });

    return <Fragment>
        <Box
            sx={{
                flexGrow: 1,
                display: 'flex',
                justifyContent: 'center',
                gap: 2,
                flexWrap: 'wrap',
                '& > *': { minWidth: 0, flexBasis: 200 },
            }}
        >
            <List
                size="sm"
                variant="outlined"
                sx={{
                    maxWidth: 300,
                    borderRadius: 'sm',
                }}>
                <ListItem>
                    LIKATO
                </ListItem>
                <ListDivider  inset='gutter' />
                <ListItem>
                    Заказы
                </ListItem>
                <ListItem>
                    Склады
                </ListItem>
            </List>
            <List
                size="sm"
                variant="outlined"
                sx={{
                    maxWidth: 300,
                    borderRadius: 'sm',
                }}>
                <ListItem>
                    <ListItemDecorator>
                        <CheckRoundedIcon />
                    </ListItemDecorator>
                    DREAMLAB
                </ListItem>
                <ListDivider  inset='gutter' />
                <ListItem>
                    Заказы
                </ListItem>
                <ListItem>
                    Склады
                </ListItem>
            </List>
        </Box>
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
        <Sheet
            className="OrderTableContainer"
            variant="outlined"
            sx={{
                //display: {xs: 'none', sm: 'initial'},
                width: '100%',
                borderRadius: 'sm',
                flexShrink: 1,
                overflow: 'auto',
                minHeight: 0,
                height: '100vh',
                '--Table-lastColumnWidth': '88px',
            }}
        >
            <MaterialReactTable table={table}/>
        </Sheet>
        {/*<JoyDataGrid columns={columns} rows={rows} page={0} pageSize={2} count={2} onPageChange={() => {*/}
        {/*}}*/}
        {/*             onPageSizeChange={() => {*/}
        {/*             }} isLoading={false} showColumns={true}/>*/}
    </Fragment>
}