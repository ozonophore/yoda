import { IAction, IActionFunction, IAuth, useController } from 'context';
import React from 'react';


export default function useAuth(): { state: IAuth, dispatch: React.Dispatch<IAction | IActionFunction>} {

    const {state, dispatch} = useController()

    const {auth} = state

    return {
        state: auth,
        dispatch
    }
}