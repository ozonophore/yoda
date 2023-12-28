import React from 'react';
import CssBaseline from '@mui/joy/CssBaseline';
import { CssVarsProvider } from '@mui/joy/styles';
import Box from '@mui/joy/Box';
import { formLabelClasses } from '@mui/joy/FormLabel';
import Typography from '@mui/joy/Typography';
import { LinearProgress } from '@mui/joy';

function ProfileLoader(): React.JSX.Element {
    return (
        <CssVarsProvider defaultMode="light" disableTransitionOnChange>
            <CssBaseline/>
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
                        [`& .${formLabelClasses.asterisk}`]: {
                            visibility: 'hidden',
                        },
                    }}
                >
                    <div style={{ textAlign: 'center' }}>
                        <Typography sx={{ py: '5px' }} level="body-sm" textAlign="center">
                            Загрузка профиля...
                        </Typography>
                        <LinearProgress/>
                    </div>
                </Box>
                <Box component="footer" sx={{py: 3}}>
                    <Typography level="body-xs" textAlign="center">
                        © {new Date().getFullYear()}
                    </Typography>
                </Box>
            </Box>
        </CssVarsProvider>
    )
}

export default ProfileLoader