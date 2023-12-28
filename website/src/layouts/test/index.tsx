import React, { Fragment } from 'react';
import JoyDataGrid, { IColumn } from 'components/JoyDataGrid';
import { Autocomplete } from '@mui/joy';

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
    return <Fragment>
        <JoyDataGrid columns={columns} rows={rows} page={0} pageSize={2} count={2} onPageChange={() => {}}
                     onPageSizeChange={() => {}} isLoading={false} showColumns={true}/>
    </Fragment>
}