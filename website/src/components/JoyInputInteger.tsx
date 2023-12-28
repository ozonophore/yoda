import Input from '@mui/joy/Input';
import React from 'react';

interface IProps {
    defaultValue: number
}

export function JoyInputInteger(props: IProps) {
    const inputRef = React.useRef<HTMLInputElement | null>(null);
    return <Input
        variant="plain" sx={{border: 0, background: 'transparent'}}
        type="number" defaultValue={props.defaultValue}
                  slotProps={{
                      input: {
                          ref: inputRef,
                          step: 1,
                      },
                  }}
    />
}

export function JoyInputNumber(props: IProps) {
    const inputRef = React.useRef<HTMLInputElement | null>(null);
    return <Input
        size='lg'
        variant="plain" sx={{border: 0, background: 'transparent'}}
        type="number" defaultValue={props.defaultValue}
        slotProps={{
            input: {
                ref: inputRef,
                step: 0.01,
            },
        }}
    />
}