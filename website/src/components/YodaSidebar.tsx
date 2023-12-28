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
import { closeSidebar, toggleSidebar } from '../utils';

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
import useProfile from 'hooks/useProfile';
import { Permission } from 'client';
import MenuIcon from '@mui/icons-material/Menu';

type Theme = 'light' | 'dark';

const themes = {
    light: {
        sidebar: {
            backgroundColor: '#ffffff',
            color: '#607489',
        },
        menu: {
            menuContent: '#fbfcfd',
            icon: '#0098e5',
            hover: {
                backgroundColor: '#c5e4ff',
                color: '#44596e',
            },
            disabled: {
                color: '#9fb6cf',
            },
        },
    },
    dark: {
        sidebar: {
            backgroundColor: '#0b2948',
            color: '#8ba1b7',
        },
        menu: {
            menuContent: '#082440',
            icon: '#59d0ff',
            hover: {
                backgroundColor: '#00458b',
                color: '#b6c8d9',
            },
            disabled: {
                color: '#3e5e7e',
            },
        },
    },
};

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

// hex to rgba converter
const hexToRgba = (hex: string, alpha: number) => {
    const r = parseInt(hex.slice(1, 3), 16);
    const g = parseInt(hex.slice(3, 5), 16);
    const b = parseInt(hex.slice(5, 7), 16);

    return `rgba(${r}, ${g}, ${b}, ${alpha})`;
};

const Dropdown = styled('i')(({theme}) => ({
    color: theme.vars.palette.text.tertiary,
}));

export default function YodaSidebar() {

    const {profile, dispatch} = useProfile()
    const [permissions, setPermissions] = React.useState<Permission[]>([])
    const [theme, setTheme] = React.useState<Theme>('light');
    const [hasImage, setHasImage] = React.useState(false);
    const [collapsed, setCollapsed] = React.useState(false);

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

    function hadleOnCollapse() {
        toggleSidebar()
    }

    return (
        <Sheet
            className="Sidebar"
            sx={{
                pt: {
                    xs: 'calc(12px + var(--Header-height))',
                    sm: 'calc(12px + var(--Header-height))',
                    md: 2,
                },
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
                    //width: '100vw',
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
            <Box sx={{display: {md: 'flex', xs: 'none', sm: 'none'}, gap: 1, alignItems: 'center'}}>
                <IconButton
                    onClick={() => toggleSidebar()}
                    variant="outlined"
                    color="neutral"
                    size="sm"
                    id='--SideToggleButton'
                >
                    <MenuIcon/>
                </IconButton>
            </Box>
            <Divider/>
            <Box
                sx={{
                    minHeight: 0,
                    overflow: 'hidden auto',
                    flexGrow: 1,
                    display: 'flex',
                    flexDirection: 'column',
                }}
            >
                <List
                    size="sm"
                    sx={{
                        '--ListItem-radius': '6px',
                        '--List-gap': '4px',
                        '--List-nestedInsetStart': '20px',
                    }}
                >
                    {renderMenu(menu)}
                </List>
            </Box>
            <Divider/>
            <Box sx={{
                display: 'flex',
                gap: 1,
                alignItems: 'center',
                flexWrap: 'wrap',
                justifyContent: 'space-between',
            }}>
                <Box sx={{
                    gap: 2,
                    display: 'flex',
                    flexWrap: 'wrap',
                    justifyContent: 'space-between',
                    alignItems: 'center',
                }}>
                    <Avatar size='sm' variant="outlined" src="/static/images/avatar/3.jpg"/>
                    <Typography fontSize="sm" fontWeight="lg">
                        {profile?.name ?? 'Неизвесный пользователь'}
                    </Typography>
                </Box>
                <IconButton
                    onClick={handleLogout}
                    // component={Link}
                    // to="/logout"
                    size='sm'
                    variant="outlined"
                    color="neutral">
                    <LogoutRoundedIcon/>
                </IconButton>
            </Box>
        </Sheet>
    );
}
