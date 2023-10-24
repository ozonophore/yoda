import React, { Fragment } from 'react';
import Box from '@mui/joy/Box';
import Typography from '@mui/joy/Typography';
import { Card } from '@mui/joy';
import Table from '@mui/joy/Table';

export default function Dashboard() {
    return <Fragment>
        <Box
            sx={{
                display: 'flex',
                my: 0.5,
                gap: 1,
                flexDirection: {xs: 'column', sm: 'row'},
                alignItems: {xs: 'start', sm: 'center'},
                flexWrap: 'wrap',
                justifyContent: 'space-between',
            }}
        >
            <Typography level="h3">Dashboard</Typography>
        </Box>
        <Box
            sx={{
                display: 'flex',
            }}
        >
            <Box sx={{
                backgroundColor: 'red'
            }}>
123
            </Box>
            <Box sx={{
                backgroundColor: 'green'
            }}>
                123
            </Box>
        </Box>
    </Fragment>
}