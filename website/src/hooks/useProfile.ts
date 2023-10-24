import { Profile } from 'client';
import { IAction, IActionFunction, useController } from 'context';
import { Dispatch } from 'react';

export default function useProfile(): { profile: Profile | undefined, dispatch: Dispatch<IAction | IActionFunction> } {
    const {state, dispatch} = useController()
    return {
        profile: state.profile,
        dispatch
    }
}