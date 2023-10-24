import * as React from 'react';
import * as ReactDOM from 'react-dom/client';
import { StyledEngineProvider } from '@mui/joy/styles';
import App from './App';
import { BrowserRouter } from 'react-router-dom';
import { YContextProvider } from './context';
import AuthProvider from 'layouts/auth';


if (process.env.NODE_ENV === 'development') {
    const {worker} = require('./mocks/browser')
    worker.start()
}

ReactDOM.createRoot(document.querySelector("#root")!).render(
    <React.StrictMode>
        <StyledEngineProvider injectFirst>
            <BrowserRouter>
                <YContextProvider>
                    <AuthProvider>
                        <App/>
                    </AuthProvider>
                </YContextProvider>
            </BrowserRouter>
        </StyledEngineProvider>
    </React.StrictMode>
);