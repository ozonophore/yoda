import React, {Fragment, useState} from 'react';
import JoyDataGrid, {IColumn} from 'components/JoyDataGrid';
import Typography from "@mui/joy/Typography";
import Input from "@mui/joy/Input";

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

export function Test() {

    const [lines, setLines] = useState(0)

    const style = {
        overflow: "hidden",
        textOverflow: "ellipsis",
        display: "-webkit-box",
        WebkitLineClamp: lines,
        WebkitBoxOrient: "vertical",
    }

    return <Fragment>
        <div>
            <Input type='number' onChange={(value) => {
                setLines(Number(value.target.value))
            }}></Input>
            <Typography noWrap={false} sx={lines !==0 ? style : {}}>
                Для вас мы собрали самые известные и проверенные временем сказки для детей. Здесь размещены русские народные сказки и авторские сказки, которые точно стоит прочитать ребенку. Детские сказки этого раздела подходят абсолютно всем ребятам: подобраны сказки для самых маленьких и для школьников. Некоторые произведения Вы найдете только у нас, в оригинальном изложении!
            </Typography>
        </div>
        <JoyDataGrid columns={columns} rows={rows} page={0} pageSize={2} count={2} onPageChange={() => {
        }}
                     onPageSizeChange={() => {
                     }} isLoading={false} showColumns={true}/>
    </Fragment>
}