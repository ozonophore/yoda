import React, { useEffect } from 'react';
import LoginPage from 'layouts/loginPage';
import { LoadProfile } from 'context/actions';
import ProfileLoader from 'layouts/profileLoader';
import { useController } from 'context';

interface IAuthProviderProps {
    children: React.ReactNode;
}

function AuthProvider(props: IAuthProviderProps): React.JSX.Element {
    const {state, dispatch} = useController();
    const {isAuth} = state.auth;
    const {profile, dicts} = state;

    useEffect(() => {
        if (isAuth) {
            dispatch(LoadProfile());
        }
    }, [isAuth]);
    return <>
        {isAuth && !profile &&
            <ProfileLoader/>
        }
        {isAuth && profile && props.children}
        {!isAuth &&
            <LoginPage/>
        }
    </>;
}

export default AuthProvider;