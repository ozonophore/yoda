import { useController } from 'context';

interface IError {
    error?: string,
    dispatch: React.Dispatch<any>
}
export default function useError() {
    const {state, dispatch} = useController()
    const {error} = state
    return {
        error,
        dispatch
    }
}