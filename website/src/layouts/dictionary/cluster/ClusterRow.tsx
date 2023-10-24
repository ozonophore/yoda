import React, { Fragment, useEffect, useState } from 'react';
import { IRow } from './model';
import Input from '@mui/joy/Input';
import Select from '@mui/joy/Select';
import Option from '@mui/joy/Option';
import Box from '@mui/joy/Box';

import equal from 'fast-deep-equal';

import IconButton from '@mui/joy/IconButton';
import CheckRoundedIcon from '@mui/icons-material/CheckRounded';
import RemoveCircleOutlineRoundedIcon from '@mui/icons-material/RemoveCircleOutlineRounded';
import ClearRoundedIcon from '@mui/icons-material/ClearRounded';

export default function ClusterRow({ id, row, onCreate, onUpdate, onDelete }: {
    id: number,
    row: IRow,
    onCreate: (value:IRow) => void,
    onUpdate: (value: IRow) => void
    onDelete: (index: number, value: IRow) => void
}): React.JSX.Element {

    const [oldData, setOldData] = useState(row)
    const [data, setData] = useState<IRow>(row)

    useEffect(() => {
        setData(row)
        setOldData(row)
    }, [row])

    const handleChange = (
        event: React.SyntheticEvent | null,
        newValue: string | null,
    ) => {
        setData({...data, source: newValue ?? "WB"})
    };

    const handleOnCancel = () => {
        setData(row)
    }

    const handleOnSubmit = () => {
        if (oldData.code == null) {
            onCreate(data)
            setOldData(data)
        } else {
            onUpdate(data)
            setOldData(data)
        }
    }

    const { code, cluster, source } = data

    return <Fragment>
        <tr key={ id }>
            <td>
                <Input sx={{border: "none", bgcolor: "transparent", marginLeft: "-8px", fontSize: "12px"}}
                       onChange={(e) => setData({...data, code: e.target.value})}
                       value={ code ?? "" } variant="plain" size="sm"></Input>
            </td>
            <td>
                <Input sx={{border: "none", bgcolor: "transparent", marginLeft: "-8px", fontSize: "12px"}}
                       onChange={(e) => setData({...data, cluster: e.target.value})}
                       value={ cluster ?? "" } variant="plain" size="sm"></Input>
            </td>
            <td>
                <Box sx={{display: "flex"}}>
                    <Select
                        sx={{border: "none", bgcolor: "transparent", width: "100%"}}
                        size="sm"
                        value={ source ?? "WB"}
                        onChange={ handleChange }
                    >
                        <Option value="WB">WB</Option>
                        <Option value="OZON">OZON</Option>
                    </Select>
                </Box>
            </td>
            <td>
                <Box sx={{ display: 'flex', gap: 1 }}>
                    <IconButton size="sm" variant="soft" color="success"
                                onClick={ handleOnSubmit }
                                disabled={ !!oldData.code && equal(oldData, data) }   >
                        <CheckRoundedIcon/>
                    </IconButton>
                    <IconButton size="sm" variant="soft" color="primary"
                                onClick={ handleOnCancel }
                                disabled={ equal(oldData, data) }>
                        <ClearRoundedIcon/>
                    </IconButton>
                    <IconButton size="sm" variant="soft" color="danger" onClick={ () => onDelete(id, oldData) }>
                        <RemoveCircleOutlineRoundedIcon/>
                    </IconButton>
                </Box>
            </td>
        </tr>
    </Fragment>
}