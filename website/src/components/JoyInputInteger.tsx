import Input from '@mui/joy/Input';
import React from 'react';

export function JoyInputInt(defaultValue: number) {
    const inputRef = React.useRef<HTMLInputElement | null>(null);
    return <Input type="number" defaultValue={defaultValue}
                  slotProps={{
                      input: {
                          ref: inputRef,
                          step: 1,
                      },
                  }}
    />
}