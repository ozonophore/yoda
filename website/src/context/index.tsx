import * as React from 'react';
import {
    createContext,
    Dispatch,
    ReactNode,
    ReducerStateWithoutAction,
    ReducerWithoutAction,
    useContext,
    useReducer
} from 'react';
import HomeRoundedIcon from '@mui/icons-material/HomeRounded';
import DashboardRoundedIcon from '@mui/icons-material/DashboardRounded';
import DynamicFeedRoundedIcon from '@mui/icons-material/DynamicFeedRounded';
import FolderRoundedIcon from '@mui/icons-material/FolderRounded';
import BubbleChartIcon from '@mui/icons-material/BubbleChart';
import QrCode2RoundedIcon from '@mui/icons-material/QrCode2Rounded';
import BlockRoundedIcon from '@mui/icons-material/BlockRounded';
import DeviceHubRoundedIcon from '@mui/icons-material/DeviceHubRounded';
import { Cluster } from '../layouts/dictionary/cluster';
import { Profile } from 'client';

interface ISubMenuData {
    name: string,
    type: "menu" | "header"
    index?: number,
    icon?: React.JSX.Element,
    component?: React.JSX.Element
}
interface IMenuData {
    name: string,
    href?: string,
    icon?: React.JSX.Element,
    subMenu?: ISubMenuData[]
}

interface IContent {
    auth: IAuth
    profile?: Profile,
    menu: {
        activeIndex: number,
        activeSubIndex: number,
        data: IMenuData[]
    },
    error?: string
}

interface IAction {
    type: string,
    payload: any
}

type IActionFunction = (dispatcher: Dispatch<IAction | IActionFunction>) => void;

interface IContext {
    state: IContent,
    dispatch: Dispatch<IAction | IActionFunction>
}

export interface IAuth {
    isAuth: boolean,
    isLoading: boolean
    error?: string
}

const InitState: IContent = {
    auth: {
        isAuth: sessionStorage.getItem("access_token") ? true : false,
        isLoading: false
    },
    menu: {
        activeIndex: 0,
        activeSubIndex: 0,
        data: [
            {
                name: "Home",
                href: "/home",
                icon: <HomeRoundedIcon href="/home"/>,
                subMenu: [
                    {
                        name: "Dashboard",
                        type: "header"
                    },
                    {
                        name: "Overview",
                        icon: <BubbleChartIcon/>,
                        type: "menu"
                    }
                ]
            }, {
                name: "Dashboard",
                href: "/dashboard",
                icon: <DashboardRoundedIcon/>,
                subMenu: []
            }, {
                name: "Справочники",
                href: "/dictionary",
                icon: <DynamicFeedRoundedIcon/>,
                subMenu: [
                    {
                        name: "Справочники",
                        type: "header"
                    },{
                        name: "Кластеры",
                        type: "menu",
                        index: 0,
                        icon: <DeviceHubRoundedIcon/>,
                        component: <Cluster/>
                    },{
                        name: "Выведенные позиции",
                        type: "menu",
                        index: 1,
                        icon: <BlockRoundedIcon/>
                    },{
                        name: "Штрих-коды",
                        type: "menu",
                        index: 2,
                        icon: <QrCode2RoundedIcon/>
                    }
                ]
            }, {
                name: "Folder",
                href: "/folder",
                icon: <FolderRoundedIcon/>,
                subMenu: []
            }
        ]
    }
}

const YContext = createContext<IContext | null>(null)
YContext.displayName = "YContext"

function reducer(state: IContent, action: IAction): IContent {
    switch (action.type) {
        case "SET_AUTH_LOADING": {
            return {
                ...state, auth: { ...state.auth, isLoading: action.payload }
            }
        }
        case "SET_ERROR": {
            return {
                ...state, error: action.payload
            }
        }
        case "SET_AUTH": {
            return {
                ...state, auth: { ...state.auth, ...action.payload }
            }
        }
        case "FIRST_MENU_SET_ACTIVE": {
            return {
                ...state, menu: {
                    ...state.menu,
                    activeIndex: action.payload,
                    activeSubIndex: 0
                }
            }
        }
        case "SECOND_MENU_SET_ACTIVE": {
            return {
                ...state, menu: {
                    ...state.menu,
                    activeSubIndex: action.payload,
                }
            }
        }
        case "SET_PROFILE": {
            return {
                ...state, profile: action.payload
            }
        }
    }
    return state
}


function useReducerWithThunk(reducer: (state: IContent, action: IAction) => IContent, initialState: IContent): [IContent, React.Dispatch<IAction | IActionFunction>] {
    const [state, dispatch] = useReducer(reducer, initialState);
    let customDispatch = (action: IAction | IActionFunction) => {
        if (typeof action === 'function') {
            action(customDispatch);
        } else {
            dispatch(action);
        }
    };
    return [state, customDispatch];
}

function YContextProvider({children}: { children: ReactNode }) {
    const [state, dispatch] = useReducerWithThunk(reducer, InitState)
    //const value = useMemo(() => [state, dispatch], [state, dispatch]);
    return <YContext.Provider value={{state, dispatch}}>
        {children}
    </YContext.Provider>
}

function useController(): IContext {
    const context = useContext(YContext)
    if (!context) {
        throw new Error(
            "useController should be used inside the YContextProvider."
        );
    }
    return context
}

export {
    YContext,
    YContextProvider,
    useController
};
export type { IAction, IActionFunction };
