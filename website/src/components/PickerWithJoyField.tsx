import * as React from 'react';
import { Dayjs } from 'dayjs';
import {
    useTheme as useMaterialTheme,
    useColorScheme as useMaterialColorScheme,
    Experimental_CssVarsProvider as MaterialCssVarsProvider,
} from '@mui/material/styles';
import {
    extendTheme as extendJoyTheme,
    useColorScheme,
    CssVarsProvider,
    THEME_ID,
} from '@mui/joy/styles';
import Input, { InputProps } from '@mui/joy/Input';
import FormControl from '@mui/joy/FormControl';
import FormLabel from '@mui/joy/FormLabel';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { DatePicker, DatePickerProps } from '@mui/x-date-pickers/DatePicker';
import {
    unstable_useDateField as useDateField,
    UseDateFieldProps,
} from '@mui/x-date-pickers/DateField';
import {
    DateFieldSlotsComponent,
    DateFieldSlotsComponentsProps,
} from '@mui/x-date-pickers/DateField/DateField.types';
import { useClearableField } from '@mui/x-date-pickers/hooks';
import {
    BaseSingleInputFieldProps,
    DateValidationError,
    FieldSection,
} from '@mui/x-date-pickers/models';

const joyTheme = extendJoyTheme();

interface JoyFieldProps extends InputProps {
    label?: React.ReactNode;
    InputProps?: {
        ref?: React.Ref<any>;
        endAdornment?: React.ReactNode;
        startAdornment?: React.ReactNode;
    };
    formControlSx?: InputProps['sx'];
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
            InputProps: { ref: containerRef, startAdornment, endAdornment } = {},
            formControlSx,
            endDecorator,
            startDecorator,
            slotProps,
            ...other
        } = props;

        return (
            <FormControl
                disabled={disabled}
                id={id}
                sx={[...(Array.isArray(formControlSx) ? formControlSx : [formControlSx])]}
                ref={ref}
            >
                <FormLabel>{label}</FormLabel>
                <Input
                    ref={ref}
                    disabled={disabled}
                    startDecorator={
                        <React.Fragment>
                            {startAdornment}
                            {startDecorator}
                        </React.Fragment>
                    }
                    endDecorator={
                        <React.Fragment>
                            {endAdornment}
                            {endDecorator}
                        </React.Fragment>
                    }
                    slotProps={{
                        ...slotProps,
                        root: { ...slotProps?.root, ref: containerRef },
                    }}
                    {...other}
                />
            </FormControl>
        );
    },
) as JoyFieldComponent;

interface JoyDateFieldProps
    extends UseDateFieldProps<Dayjs>,
        BaseSingleInputFieldProps<
            Dayjs | null,
            Dayjs,
            FieldSection,
            DateValidationError
        > {}

const JoyDateField = React.forwardRef(
    (props: JoyDateFieldProps, ref: React.Ref<HTMLDivElement>) => {
        const {
            inputRef: externalInputRef,
            slots,
            slotProps,
            ...textFieldProps
        } = props;

        const {
            onClear,
            clearable,
            ref: inputRef,
            ...fieldProps
        } = useDateField<Dayjs, typeof textFieldProps>({
            props: textFieldProps,
            inputRef: externalInputRef,
        });

        /* If you don't need a clear button, you can skip the use of this hook */
        const { InputProps: ProcessedInputProps, fieldProps: processedFieldProps } =
            useClearableField<
                {},
                typeof textFieldProps.InputProps,
                DateFieldSlotsComponent,
                DateFieldSlotsComponentsProps<Dayjs>
            >({
                onClear,
                clearable,
                fieldProps,
                InputProps: fieldProps.InputProps,
                slots,
                slotProps,
            });

        return (
            <JoyField
                ref={ref}
                slotProps={{
                    input: {
                        ref: inputRef,
                    },
                }}
                {...processedFieldProps}
                InputProps={ProcessedInputProps}
            />
        );
    },
);

const JoyDatePicker = React.forwardRef(
    (props: DatePickerProps<Dayjs>, ref: React.Ref<HTMLDivElement>) => {
        return (
            <DatePicker
                ref={ref}
                {...props}
                slots={{ field: JoyDateField, ...props.slots }}
                slotProps={{
                    ...props.slotProps,
                    field: {
                        ...props.slotProps?.field,
                        formControlSx: {
                            flexDirection: 'row',
                        },
                    } as any,
                }}
            />
        );
    },
);

/**
 * This component is for syncing the theme mode of this demo with the MUI docs mode.
 * You might not need this component in your project.
 */
function SyncThemeMode({ mode }: { mode: 'light' | 'dark' }) {
    const { setMode } = useColorScheme();
    const { setMode: setMaterialMode } = useMaterialColorScheme();
    React.useEffect(() => {
        setMode(mode);
        setMaterialMode(mode);
    }, [mode, setMode, setMaterialMode]);
    return null;
}

export default function PickerWithJoyField(props: DatePickerProps<Dayjs>) {
    const materialTheme = useMaterialTheme();
    return (
        <MaterialCssVarsProvider>
            <CssVarsProvider theme={{ [THEME_ID]: joyTheme }}>
                <SyncThemeMode mode={materialTheme.palette.mode} />
                <LocalizationProvider adapterLocale='ru' dateAdapter={AdapterDayjs}>
                    <JoyDatePicker
                        {...props}
                        slotProps={{
                            field: { clearable: false },
                        }}
                    />
                </LocalizationProvider>
            </CssVarsProvider>
        </MaterialCssVarsProvider>
    );
}