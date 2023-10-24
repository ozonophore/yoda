import {
    BaseSingleInputFieldProps,
    DatePicker,
    DatePickerProps,
    DateValidationError,
    FieldSection
} from '@mui/x-date-pickers';
import Input, { InputProps } from '@mui/joy/Input';
import React from 'react';
import { unstable_useDateField as useDateField, UseDateFieldProps, } from '@mui/x-date-pickers/DateField';
import { Dayjs } from 'dayjs';
import FormControl from '@mui/joy/FormControl';
import FormLabel from '@mui/joy/FormLabel';
import { OverridableStringUnion } from '@mui/types';
import { InputPropsSizeOverrides } from '@mui/joy/Input/InputProps';

interface JoyFieldProps extends InputProps {
    label?: React.ReactNode;
    InputProps?: {
        ref?: React.Ref<any>;
        endAdornment?: React.ReactNode;
        startAdornment?: React.ReactNode;
    };
    formControlSx?: InputProps['sx'];
}

interface JoyDateFieldProps
    extends UseDateFieldProps<Dayjs>,
        BaseSingleInputFieldProps<
            Dayjs | null,
            Dayjs,
            FieldSection,
            DateValidationError
        > {
}

type JoyFieldComponent = ((
    props: JoyFieldProps & React.RefAttributes<HTMLDivElement>,
) => React.JSX.Element) & { propTypes?: any };

const JoyField = React.forwardRef(
    (props: JoyFieldProps, ref: React.Ref<HTMLDivElement>) => {
        const {
            disabled,
            id,
            label,
            InputProps: {ref: containerRef, startAdornment, endAdornment} = {},
            formControlSx,
            slotProps,
            ...other
        } = props;

        return (
            <FormControl
                disabled={disabled}
                id={id}
                sx={[
                    {
                        flexGrow: 1,
                    },
                    ...(Array.isArray(formControlSx) ? formControlSx : [formControlSx]),
                ]}
                ref={ref}
            >
                <FormLabel>{label}</FormLabel>
                <Input
                    disabled={disabled}
                    slotProps={{
                        ...slotProps,
                        root: {...slotProps?.root, ref: containerRef},
                    }}
                    startDecorator={startAdornment}
                    endDecorator={endAdornment}
                    {...other}
                />
            </FormControl>
        );
    },
) as JoyFieldComponent;

const JoyDateField = React.forwardRef(
    (props: JoyDateFieldProps & {
        size?: OverridableStringUnion<'sm' | 'md' | 'lg', InputPropsSizeOverrides>
    }, ref: React.Ref<HTMLDivElement>) => {
        const {
            inputRef: externalInputRef,
            slots,
            slotProps,
            size,
            ...textFieldProps
        } = props;

        const {ref: inputRef, ...other} = useDateField<Dayjs, typeof textFieldProps>({
            props: textFieldProps,
            inputRef: externalInputRef,
        });

        return (
            <JoyField
                ref={ref}
                size={props.size}
                slotProps={{
                    input: {
                        ref: inputRef,
                    },
                }}
                {...other}
            />
        );
    },
);

const JoyDatePicker = React.forwardRef(
    (props: DatePickerProps<Dayjs> & {
        size?: OverridableStringUnion<'sm' | 'md' | 'lg', InputPropsSizeOverrides>;
    }, ref: React.Ref<HTMLDivElement>) => {
        return (
            <DatePicker
                ref={ref}
                {...props}
                slots={{field: JoyDateField, ...props.slots}}
                slotProps={{
                    field: {
                        formControlSx: {
                            flexDirection: 'row',
                        },
                        size: props.size
                    } as any,
                }}
            />
        );
    },
);

export default JoyDatePicker