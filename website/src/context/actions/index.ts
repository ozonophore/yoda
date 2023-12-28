import { IAction, IActionFunction } from '../index';
import { AuthService, OpenAPI, Profile } from 'client';
import { Dispatch } from 'react';

const SetError = (error?: string): IAction => {
    return {
        type: "SET_ERROR",
        payload: error
    }
}

const SetAuthLoading = (loading: boolean): IAction => {
    return {
        type: "SET_AUTH_LOADING",
        payload: loading
    }
}
const SetLogout = (error?: string): IActionFunction => {
    return (dispatch: Dispatch<IAction>) => {
        sessionStorage.removeItem("access_token")
        dispatch(setAuth(false, error))
        dispatch(CleanProfile())
    }
}

const SetLogin = (token?: string): IActionFunction => {
    return (dispatch: Dispatch<IAction>) => {
        if (token) {
            OpenAPI.TOKEN = token
            sessionStorage.setItem("access_token", token)
            dispatch(setAuth(true))
        } else {
            sessionStorage.removeItem("access_token")
            dispatch(setAuth(false))
        }
    }
}

const setAuth = (auth: boolean, error?: string): IAction => {
    return {
        type: "SET_AUTH",
        payload: {isAuth: auth, error}
    }
}

const FirstSidebarActive = (index: number): IAction => {
    return {
        type: "FIRST_MENU_SET_ACTIVE",
        payload: index
    }
}

const SecondSidebarActive = (index: number): IAction => {
    return {
        type: "SECOND_MENU_SET_ACTIVE",
        payload: index
    }
}

const authentication = (email: string, password: string): IActionFunction => {
    return (dispatch: Dispatch<IAction | IActionFunction>) => {
        dispatch(SetAuthLoading(true))
        AuthService.login({email, password})
            .then(resp => {
                if (resp.access_token) {
                    dispatch(SetLogin(resp.access_token))
                } else {
                    dispatch(SetLogout(resp.description))
                }
            }).catch(err => {
            dispatch(SetLogout(err.body.description))
        }).finally(() => {
            dispatch(SetAuthLoading(false))
        })
    }
}

const SetProfile = (profile: Profile): IAction => {
    return {
        type: "SET_PROFILE",
        payload: profile
    }
}

const CleanProfile = (): IAction => {
    return {
        type: "SET_PROFILE",
        payload: null
    }
}

const LoadProfile = (): IActionFunction => {
    return (dispatch: Dispatch<IActionFunction | IAction>) => {
        AuthService.profile()
            .then(resp => {
                dispatch(SetProfile(resp))
            }).catch(err => {
                const description = err.body.description
                dispatch(SetLogout(description));
        })
    }
}

const SetMenuActive = (key: string): IAction  => {
    console.log("#", key)
    return {
        type: "SET_SIDEBAR_ACTIVE",
        payload: key
    }
}

export {
    FirstSidebarActive,
    SecondSidebarActive,
    SetProfile,
    CleanProfile,
    authentication,
    SetLogin,
    SetLogout,
    SetError,
    SetAuthLoading,
    LoadProfile,
    SetMenuActive
}