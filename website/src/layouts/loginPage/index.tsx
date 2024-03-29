import * as React from 'react';
import { CssVarsProvider, useColorScheme } from '@mui/joy/styles';
import GlobalStyles from '@mui/joy/GlobalStyles';
import CssBaseline from '@mui/joy/CssBaseline';
import Box from '@mui/joy/Box';
import Button from '@mui/joy/Button';
import Checkbox from '@mui/joy/Checkbox';
import FormControl from '@mui/joy/FormControl';
import FormLabel, { formLabelClasses } from '@mui/joy/FormLabel';
import IconButton, { IconButtonProps } from '@mui/joy/IconButton';
import Link from '@mui/joy/Link';
import Input from '@mui/joy/Input';
import Typography from '@mui/joy/Typography';
import DarkModeRoundedIcon from '@mui/icons-material/DarkModeRounded';
import LightModeRoundedIcon from '@mui/icons-material/LightModeRounded';

import dayImg from '../../assets/day.avif';
import nightImg from '../../assets/night.avif';
import { authentication } from 'context/actions';
import { Alert } from '@mui/joy';
import useAuth from 'hooks/useAuth';

interface FormElements extends HTMLFormControlsCollection {
    email: HTMLInputElement;
    password: HTMLInputElement;
    persistent: HTMLInputElement;
}

interface SignInFormElement extends HTMLFormElement {
    readonly elements: FormElements;
}

function ColorSchemeToggle({onClick, ...props}: IconButtonProps) {
    const {mode, setMode} = useColorScheme();
    const storedMode = localStorage.getItem("mode") ?? mode;
    const [mounted, setMounted] = React.useState(false);
    React.useEffect(() => {
        setMounted(true);
    }, []);
    if (!mounted) {
        return <IconButton size="sm" variant="plain" color="neutral" disabled/>;
    }
    return (
        <IconButton
            id="toggle-mode"
            size="sm"
            variant="plain"
            color="neutral"
            aria-label="toggle light/dark mode"
            {...props}
            onClick={(event) => {
                if (storedMode === 'dark') {
                    localStorage.setItem("mode", "light")
                    setMode('light');
                } else {
                    localStorage.setItem("mode", "dark")
                    setMode('dark');
                }
                onClick?.(event);
            }}
        >
            {storedMode === 'light' ? <DarkModeRoundedIcon/> : <LightModeRoundedIcon/>}
        </IconButton>
    );
}

export default function LoginPage() {

    const {state, dispatch} = useAuth()

    const handleOnSubmit = (event: React.FormEvent<SignInFormElement>) => {
        event.preventDefault();
        const formElements = event.currentTarget.elements;
        const data = {
            email: formElements.email.value,
            password: formElements.password.value,
            //persistent: formElements.persistent.checked,
        };

        dispatch(authentication(data.email, data.password))
    }

    return (
        <CssVarsProvider defaultMode="light" disableTransitionOnChange>
            <CssBaseline/>
            <GlobalStyles
                styles={{
                    ':root': {
                        '--Collapsed-breakpoint': '769px', // form will stretch when viewport is below `769px`
                        '--Cover-width': '40vw', // must be `vw` only
                        '--Form-maxWidth': '700px',
                        '--Transition-duration': '0.4s', // set to `none` to disable transition
                    },
                }}
            />
            <Box
                sx={(theme) => ({
                    width:
                        'clamp(100vw - var(--Cover-width), (var(--Collapsed-breakpoint) - 100vw) * 999, 100vw)',
                    transition: 'width var(--Transition-duration)',
                    transitionDelay: 'calc(var(--Transition-duration) + 0.1s)',
                    position: 'relative',
                    zIndex: 1,
                    display: 'flex',
                    justifyContent: 'flex-end',
                    backdropFilter: 'blur(4px)',
                    backgroundColor: 'rgba(255 255 255 / 0.6)',
                    [theme.getColorSchemeSelector('dark')]: {
                        backgroundColor: 'rgba(19 19 24 / 0.4)',
                    },
                })}
            >
                <Box
                    sx={{
                        display: 'flex',
                        flexDirection: 'column',
                        minHeight: '100dvh',
                        width:
                            'clamp(var(--Form-maxWidth), (var(--Collapsed-breakpoint) - 100vw) * 999, 100%)',
                        maxWidth: '100%',
                        px: 2,
                    }}
                >
                    <Box
                        component="header"
                        sx={{
                            py: 3,
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: 'space-between',
                        }}
                    >
                        <Typography
                            fontWeight="lg"
                            startDecorator={
                                <Box
                                    component="span"
                                    sx={{
                                        width: 24,
                                        height: 24,
                                        background: (theme) =>
                                            `linear-gradient(45deg, ${theme.vars.palette.primary.solidBg}, ${theme.vars.palette.primary.solidBg} 30%, ${theme.vars.palette.primary.softBg})`,
                                        borderRadius: '50%',
                                        boxShadow: (theme) => theme.shadow.md,
                                        '--joy-shadowChannel': (theme) =>
                                            theme.vars.palette.primary.mainChannel,
                                    }}
                                />
                            }
                        >
                        </Typography>
                        <ColorSchemeToggle/>
                    </Box>
                    <Box
                        component="main"
                        sx={{
                            my: 'auto',
                            py: 2,
                            pb: 5,
                            display: 'flex',
                            flexDirection: 'column',
                            gap: 2,
                            width: 400,
                            maxWidth: '100%',
                            mx: 'auto',
                            borderRadius: 'sm',
                            '& form': {
                                display: 'flex',
                                flexDirection: 'column',
                                gap: 2,
                            },
                            [`& .${formLabelClasses.asterisk}`]: {
                                visibility: 'hidden',
                            },
                        }}
                    >
                        <div>
                            <Typography component="h1" fontSize="xl2" fontWeight="lg">
                                Войти
                            </Typography>
                            <Typography level="body-sm" sx={{my: 1, mb: 3}}>
                                Добро пожаловать
                            </Typography>
                        </div>
                        {state.error &&
                            <Alert key='Title' size="sm" sx={{alignItems: 'flex-start'}} variant="soft" color='danger'>
                                {state.error}
                            </Alert>
                        }
                        <form
                            onSubmit={handleOnSubmit}
                        >
                            <FormControl required>
                                <FormLabel>Логин</FormLabel>
                                <Input type="email" name="email"/>
                            </FormControl>
                            <FormControl required>
                                <FormLabel>Пароль</FormLabel>
                                <Input type="password" name="password"/>
                            </FormControl>
                            <Box
                                height="10px"
                                sx={{
                                    display: 'flex',
                                    justifyContent: 'space-between',
                                    alignItems: 'center',
                                }}
                            >
                            {/*    <Checkbox size="sm" label="Remember for 30 days" name="persistent"/>*/}
                            {/*    <Link fontSize="sm" href="#replace-with-a-link" fontWeight="lg">*/}
                            {/*        Забыли пароль?*/}
                            {/*    </Link>*/}
                            </Box>
                            <Button type="submit" fullWidth disabled={state.isLoading}>
                                Войти
                            </Button>
                        </form>
                    </Box>
                    <Box component="footer" sx={{py: 3}}>
                        <Typography level="body-xs" textAlign="center">
                            © {new Date().getFullYear()} ({process.env.REACT_APP_VERSION})
                        </Typography>
                    </Box>
                </Box>
            </Box>
            <Box
                sx={(theme) => ({
                    height: '100%',
                    position: 'fixed',
                    right: 0,
                    top: 0,
                    bottom: 0,
                    left: 'clamp(0px, (100vw - var(--Collapsed-breakpoint)) * 999, 100vw - var(--Cover-width))',
                    transition:
                        'background-image var(--Transition-duration), left var(--Transition-duration) !important',
                    transitionDelay: 'calc(var(--Transition-duration) + 0.1s)',
                    backgroundColor: 'background.level1',
                    backgroundSize: 'cover',
                    backgroundPosition: 'center',
                    backgroundRepeat: 'no-repeat',
                    backgroundImage:
                        `url(${dayImg})`,
                    [theme.getColorSchemeSelector('dark')]: {
                        backgroundImage:
                            `url(${nightImg})`,
                    },
                })}
            />
        </CssVarsProvider>
    );
}
