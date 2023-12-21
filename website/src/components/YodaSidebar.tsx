import * as React from 'react';
import { Fragment, useEffect } from 'react';
import { styled } from '@mui/joy/styles';
import GlobalStyles from '@mui/joy/GlobalStyles';
import Avatar from '@mui/joy/Avatar';
import Box from '@mui/joy/Box';
import Divider from '@mui/joy/Divider';
import IconButton from '@mui/joy/IconButton';
import List from '@mui/joy/List';
import ListItem from '@mui/joy/ListItem';
import ListItemButton from '@mui/joy/ListItemButton';
import ListItemContent from '@mui/joy/ListItemContent';
import ListItemDecorator from '@mui/joy/ListItemDecorator';
import Typography from '@mui/joy/Typography';
import Sheet from '@mui/joy/Sheet';
import { closeSidebar } from '../utils';

import AutoStoriesRoundedIcon from '@mui/icons-material/AutoStoriesRounded';
import HomeRoundedIcon from '@mui/icons-material/HomeRounded';
import DynamicFeedRoundedIcon from '@mui/icons-material/DynamicFeedRounded';
import DashboardRoundedIcon from '@mui/icons-material/DashboardRounded';
import FolderRoundedIcon from '@mui/icons-material/FolderRounded';
import DeviceHubRoundedIcon from '@mui/icons-material/DeviceHubRounded';
import ShoppingCartIcon from '@mui/icons-material/ShoppingCart';
import PointOfSaleRoundedIcon from '@mui/icons-material/PointOfSaleRounded';
import LogoutRoundedIcon from '@mui/icons-material/LogoutRounded';
import DensitySmallRoundedIcon from '@mui/icons-material/DensitySmallRounded';

import ListSubheader from '@mui/joy/ListSubheader';
import { Link, useLocation, useNavigate } from 'react-router-dom';
import { SetLogout } from 'context/actions';
import useProfile from 'hooks/useProfile';
import { Permission } from 'client';

import { Sidebar, Menu , MenuItem } from "react-pro-sidebar";

interface IMenuItem {
    id: string
    title: string
    icon: React.JSX.Element
    href?: string
    menu?: IMenuItem[]
    permission?: string
}

const menu = [
    {
        id: "/",
        title: "Главная",
        permission: Permission.HOME,
        icon: <HomeRoundedIcon/>,
        href: "/"
    }, {
        id: "menu_1",
        title: "Dashboard",
        permission: Permission.DASHBOARD,
        icon: <DashboardRoundedIcon/>,
        href: "dashboard"
    }, {
        id: "menu_2",
        title: "Справочники",
        icon: <DynamicFeedRoundedIcon/>,
        menu: [
            {
                id: "/clusters",
                title: "Кластеры",
                icon: <DeviceHubRoundedIcon/>,
                href: "clusters"
            },
            {
                id: "/positions",
                title: "Позиции",
                icon: <DensitySmallRoundedIcon/>,
                href: "positions"
            }
        ]
    }, {
        id: "menu_3",
        title: "Отчеты",
        icon: <FolderRoundedIcon/>,
        menu: [
            {
                id: "/order-by-day",
                title: "Заказы за день",
                permission: Permission.ORDERS,
                icon: <ShoppingCartIcon/>,
                href: "order-by-day"
            },
            {
                id: "/sales-by-month",
                title: "Продажи за месяц",
                permission: Permission.SALES,
                icon: <PointOfSaleRoundedIcon/>,
                href: "sales-by-month"
            }
        ]
    }, {
        id: "menu_4",
        title: "Заказы",
        icon: <FolderRoundedIcon/>,
        menu: [
            {
                id: "/order-product-by-day",
                title: "За период",
                permission: Permission.ORDERS,
                icon: <ShoppingCartIcon/>,
                href: "order-product-by-day"
            }
        ]
    }, {
        id: "menu_5",
        title: "Справочники",
        icon: <FolderRoundedIcon/>,
        menu: [
            {
                id: "/dict-position",
                title: "Позиций 1С",
                permission: Permission.DICTIONARY,
                icon: <AutoStoriesRoundedIcon/>,
                href: "dict-position"
            }
        ]
    }
]

const Dropdown = styled('i')(({theme}) => ({
    color: theme.vars.palette.text.tertiary,
}));

export default function Sidebar() {

    const {profile, dispatch} = useProfile()
    const [permissions, setPermissions] = React.useState<Permission[]>([])

    const navigate = useNavigate();

    const {pathname} = useLocation();

    const selected = !pathname ? "/" : pathname

    useEffect(() => {
        setPermissions(profile?.permissions ?? [])
    }, [profile]);

    function contains(items: Permission[], item: IMenuItem): boolean {
        if (!item.permission && item.menu) {
            for (const i of item.menu) {
                if (contains(items, i)) return true
            }
        }
        const value = item?.permission ?? ""
        if (!value) return false
        return items.includes(value as Permission)
    }

    function renderMenu(menu: IMenuItem[]): React.JSX.Element[] {
        return (menu.map((item: IMenuItem) => (
                <Fragment>
                    {contains(permissions, item) &&
                        <ListItem key={item.id} nested={!!item.menu}>
                            {!!(item.href) && <ListItemButton
                                to={item.href}
                                state={{selected: item.id}}
                                selected={item.id === selected}
                                style={{
                                    color: (item.id === selected) ? '#0052cc' : ''
                                }}
                                component={Link}>
                                <ListItemDecorator>
                                    {item.icon}
                                </ListItemDecorator>
                                <ListItemContent>
                                    {item.title}
                                </ListItemContent>
                            </ListItemButton>}
                            {!(item.href) && <ListSubheader>
                                <ListItemDecorator>
                                    {item.icon}
                                </ListItemDecorator>
                                <ListItemContent>{item.title}</ListItemContent>
                            </ListSubheader>}
                            {!!(item.menu) && <List>
                                {renderMenu(item.menu)}
                            </List>}
                        </ListItem>
                }
                </Fragment>
            ))
        )
    }

    const handleLogout = (event: React.FormEvent) => {
    }

    return (
        <Sheet
            className="Sidebar"
            sx={{
                position: {
                    xs: 'fixed',
                    md: 'sticky',
                },
                transform: {
                    xs: 'translateX(calc(100% * (var(--SideNavigation-slideIn, 0) - 1)))',
                    md: 'none',
                },
                transition: 'transform 0.4s, width 0.4s',
                zIndex: 1000,
                height: '100dvh',
                width: 'var(--Sidebar-width)',
                top: 0,
                p: 1.5,
                py: 3,
                flexShrink: 0,
                display: 'flex',
                flexDirection: 'column',
                gap: 2,
                borderRight: '1px solid',
                borderColor: 'divider',
            }}
        >
            <GlobalStyles
                styles={(theme) => ({
                    ':root': {
                        '--Sidebar-width': '230px',
                        [theme.breakpoints.up('lg')]: {
                            '--Sidebar-width': '230px',
                        },
                    },
                })}
            />
            <Box
                className="Sidebar-overlay"
                sx={{
                    position: 'fixed',
                    zIndex: 9998,
                    top: 0,
                    left: 0,
                    width: '100vw',
                    height: '100vh',

                    opacity: 'calc(var(--SideNavigation-slideIn, 0) - 0.2)',
                    transition: 'opacity 0.4s',
                    transform: {
                        xs: 'translateX(calc(100% * (var(--SideNavigation-slideIn, 0) - 1) + var(--SideNavigation-slideIn, 0) * var(--Sidebar-width, 0px)))',
                        lg: 'translateX(-100%)',
                    },
                }}
                onClick={() => closeSidebar()}
            />
            <Box
                sx={{
                    minHeight: 0,
                    overflow: 'hidden auto',
                    flexGrow: 1,
                    display: 'flex',
                    flexDirection: 'column',
                }}
            >
                {/*<List*/}
                {/*    size="sm"*/}
                {/*    sx={{*/}
                {/*        '--ListItem-radius': '6px',*/}
                {/*        '--List-gap': '4px',*/}
                {/*        '--List-nestedInsetStart': '20px',*/}
                {/*    }}*/}
                {/*>*/}
                {/*    {renderMenu(menu)}*/}
                {/*</List>*/}
                <Sidebar>
                    <Menu>

                    </Menu>
                </Sidebar>
            </Box>
            <Divider/>
            <Box sx={{display: 'flex', gap: 1, alignItems: 'center'}}>
                <Avatar variant="outlined" src="/static/images/avatar/3.jpg"/>
                <Box sx={{minWidth: 0, flex: 1}}>
                    <Typography fontSize="sm" fontWeight="lg">
                        {profile?.name ?? 'Неизвесный пользователь'}
                    </Typography>
                    <Typography level="body-xs"></Typography>
                </Box>
                <IconButton
                    onClick={handleLogout}
                    // component={Link}
                    // to="/logout"
                    variant="plain"
                    color="neutral">
                    <LogoutRoundedIcon/>
                </IconButton>
            </Box>
        </Sheet>
    );
}
