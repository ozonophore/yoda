import * as React from 'react';
import { Fragment } from 'react';
import Typography from '@mui/joy/Typography';
import Button from '@mui/joy/Button';
import DownloadRoundedIcon from '@mui/icons-material/DownloadRounded';
import Box from '@mui/joy/Box';
import ClusterTable from './ClusterTable';

export function Cluster() {
    return (
        <Fragment>
            <Box
                sx={{
                    display: 'flex',
                    my: 1,
                    gap: 1,
                    flexDirection: {xs: 'column', sm: 'row'},
                    alignItems: {xs: 'start', sm: 'center'},
                    flexWrap: 'wrap',
                    justifyContent: 'space-between',
                }}
            >
                <Typography level="h3">Кластеры</Typography>
                <Button
                    color="primary"
                    startDecorator={<DownloadRoundedIcon/>}
                    size="sm"
                >
                    Скачать Excel
                </Button>
            </Box>
            <ClusterTable/>
        </Fragment>
    )
}