import * as React from 'react';
import { Fragment } from 'react';
import { CssVarsProvider, StyledEngineProvider } from '@mui/joy/styles';
import CssBaseline from '@mui/joy/CssBaseline';
import Box from '@mui/joy/Box';
import Breadcrumbs from '@mui/joy/Breadcrumbs';

import {
    Experimental_CssVarsProvider as MaterialCssVarsProvider,
    experimental_extendTheme as materialExtendTheme,
    THEME_ID as MATERIAL_THEME_ID,
} from '@mui/material/styles';
// icons
import ChevronRightRoundedIcon from '@mui/icons-material/ChevronRightRounded';
import ColorSchemeToggle from './components/ColorSchemeToggle';
import Header from './components/Header';
import Sidebar from './components/Sidebar';
import { Outlet, Route, Routes } from 'react-router-dom';
import { Cluster } from './layouts/dictionary/cluster';
import Dashboard from './layouts/dashboard';
import Home from './layouts/home';
import { OrderByDay } from './layouts/orderByDay';
import { SaleByMonth } from 'layouts/saleByMonth';
import { Positions } from 'layouts/dictionary/positions';
import { Snackbar } from '@mui/material';
import { Alert } from '@mui/joy';
import useError from 'hooks/useError';
import { SetError } from 'context/actions';
import IconButton from '@mui/joy/IconButton';

import Close from '@mui/icons-material/Close';

const useEnhancedEffect =
    typeof window !== 'undefined' ? React.useLayoutEffect : React.useEffect;

function Layout() {

    const {error, dispatch} = useError()
    const handleClose = () => {
        dispatch(SetError(undefined))
    }
    return (<Fragment>
        <Sidebar/>
        <Box
            component="main"
            className="MainContent"
            sx={{
                px: {
                    xs: 2,
                    md: 6,
                },
                pt: {
                    xs: 'calc(12px + var(--Header-height))',
                    sm: 'calc(12px + var(--Header-height))',
                    md: 3,
                },
                pb: {
                    xs: 2,
                    sm: 2,
                    md: 3,
                },
                flex: 1,
                display: 'flex',
                flexDirection: 'column',
                minWidth: 0,
                height: '100dvh',
                gap: 1,
            }}
        >
            <Box sx={{display: 'flex', alignItems: 'center'}}>
                <Breadcrumbs
                    size="sm"
                    aria-label="breadcrumbs"
                    separator={<ChevronRightRoundedIcon fontSize="small"/>}
                >

                </Breadcrumbs>
                <ColorSchemeToggle
                    sx={{ml: 'auto', display: {xs: 'none', md: 'inline-flex'}}}
                />
            </Box>
            <Outlet/>
            <Snackbar
                open={!!error}
                autoHideDuration={6000}
                onClose={handleClose}
                anchorOrigin={{vertical: "bottom", horizontal: "center"}}
            >
                <Alert size="md"
                       color="danger"
                       endDecorator={
                           <IconButton
                               variant="plain"
                               sx={{
                                   '--IconButton-size': '32px',
                                   transform: 'translate(0.5rem, -0.5rem)',
                               }}
                               onClick={handleClose}
                           >
                               <Close/>
                           </IconButton>
                       }
                >
                    {error}
                </Alert>
            </Snackbar>
        </Box>
    </Fragment>)
}

const materialTheme = materialExtendTheme();

export default function App() {
    return (
        <MaterialCssVarsProvider theme={{[MATERIAL_THEME_ID]: materialTheme}}>
            <StyledEngineProvider injectFirst>
                <CssVarsProvider disableTransitionOnChange>
                    <CssBaseline/>
                    <Box sx={{display: 'flex', minHeight: '100dvh'}}>
                        <Header/>
                        <Routes>
                            <Route path="/" element={<Layout/>}>
                                <Route index element={<Home/>}/>
                                <Route path="dashboard" element={<Dashboard/>}/>
                                <Route path="clusters" element={<Cluster/>}/>
                                <Route path="order-by-day" element={<OrderByDay/>}/>
                                <Route path="sales-by-month" element={<SaleByMonth/>}/>
                                <Route path="positions" element={<Positions/>}/>
                            </Route>
                        </Routes>
                    </Box>
                </CssVarsProvider>
            </StyledEngineProvider>
        </MaterialCssVarsProvider>
    );
}
