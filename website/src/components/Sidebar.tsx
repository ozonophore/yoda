import * as React from 'react';
import GlobalStyles from '@mui/joy/GlobalStyles';
import Box from '@mui/joy/Box';
import Divider from '@mui/joy/Divider';
import IconButton from '@mui/joy/IconButton';
import List from '@mui/joy/List';
import ListItem from '@mui/joy/ListItem';
import ListItemButton, {listItemButtonClasses} from '@mui/joy/ListItemButton';
import ListItemContent from '@mui/joy/ListItemContent';
import Typography from '@mui/joy/Typography';
import Sheet from '@mui/joy/Sheet';
import HomeRoundedIcon from '@mui/icons-material/HomeRounded';
import ShoppingCartRoundedIcon from '@mui/icons-material/ShoppingCartRounded';
import FolderRoundedIcon from '@mui/icons-material/FolderRounded';
import LogoutRoundedIcon from '@mui/icons-material/LogoutRounded';
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';
import StorefrontRoundedIcon from '@mui/icons-material/StorefrontRounded';

import ColorSchemeToggle from './ColorSchemeToggle';
import {closeSidebar, toggleSidebar} from '../utils';
import {Link} from 'react-router-dom';
import useProfile from 'hooks/useProfile';
import MenuIcon from '@mui/icons-material/Menu';
import {useController} from 'context';

function Toggler({
                     defaultExpanded = true,
                     renderToggle,
                     children,
                 }: {
    defaultExpanded?: boolean;
    children: React.ReactNode;
    renderToggle: (params: {
        open: boolean;
        setOpen: React.Dispatch<React.SetStateAction<boolean>>;
    }) => React.ReactNode;
}) {
    const [open, setOpen] = React.useState(defaultExpanded);
    return (
        <React.Fragment>
            {renderToggle({open, setOpen})}
            <Box
                sx={{
                    display: 'grid',
                    gridTemplateRows: open ? '1fr' : '0fr',
                    transition: '0.2s ease',
                    '& > *': {
                        overflow: 'hidden',
                    },
                }}
            >
                {children}
            </Box>
        </React.Fragment>
    );
}

export default function Sidebar(this: any) {
    const {state} = useController()
    const {profile, dispatch} = useProfile()

    const {active} = state.sidebar

    function handleOnClickMenu(event: any) {
    }

    return (
        <Sheet
            className="Sidebar"
            sx={{
                position: {xs: 'fixed', md: 'sticky'},
                transform: {
                    xs: 'translateX(calc(100% * (var(--SideNavigation-slideIn, 0) - 1)))',
                    md: 'none',
                },
                transition: 'transform 0.4s, width 0.4s',
                zIndex: 10000,
                height: '100dvh',
                width: 'var(--Sidebar-width)',
                top: 0,
                p: 2,
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
                        '--Sidebar-width': '220px',
                        [theme.breakpoints.up('lg')]: {
                            '--Sidebar-width': '240px',
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
                    opacity: 'var(--SideNavigation-slideIn)',
                    backgroundColor: 'var(--joy-palette-background-backdrop)',
                    transition: 'opacity 0.4s',
                    transform: {
                        xs: 'translateX(calc(100% * (var(--SideNavigation-slideIn, 0) - 1) + var(--SideNavigation-slideIn, 0) * var(--Sidebar-width, 0px)))',
                        lg: 'translateX(-100%)',
                    },
                }}
                onClick={() => closeSidebar()}
            />
            <Box sx={{display: 'flex', gap: 1, alignItems: 'center'}}>
                <IconButton
                    onClick={() => toggleSidebar()}
                    variant="outlined"
                    color="neutral"
                    size='sm'
                    sx={{
                        display: {
                            md: 'none'
                        }
                    }}
                >
                    <MenuIcon/>
                </IconButton>
                <ColorSchemeToggle sx={{ml: 'auto'}}/>
            </Box>
            <Divider/>
            <Box
                sx={{
                    minHeight: 0,
                    overflow: 'hidden auto',
                    flexGrow: 1,
                    display: 'flex',
                    flexDirection: 'column',
                    [`& .${listItemButtonClasses.root}`]: {
                        gap: 1.5,
                    },
                }}
            >
                <List
                    size="sm"
                    sx={{
                        gap: 1,
                        '--List-nestedInsetStart': '30px',
                        '--ListItem-radius': (theme) => theme.vars.radius.sm,
                    }}
                >

                    <ListItem>
                        <ListItemButton id="menu-home-id" selected={active === "menu-home-id"}
                                        onClick={handleOnClickMenu}>
                            <HomeRoundedIcon/>
                            <ListItemContent>
                                <Typography style={{textDecoration: 'none'}} to={"/"} component={Link}
                                            level="title-sm">Home</Typography>
                            </ListItemContent>
                        </ListItemButton>
                    </ListItem>

                    <ListItem nested>
                        <Toggler renderToggle={({open, setOpen}) => (
                            <ListItemButton onClick={() => setOpen(!open)}>
                                <ShoppingCartRoundedIcon/>
                                <ListItemContent>
                                    <Typography level="title-sm">Заказы</Typography>
                                </ListItemContent>
                                <KeyboardArrowDownIcon
                                    sx={{transform: open ? 'rotate(180deg)' : 'none'}}
                                />
                            </ListItemButton>
                        )}>
                            <List sx={{gap: 0.5}}>
                                <ListItem sx={{mt: 0.5}}>
                                    <ListItemButton
                                        id="menu-orders-day-id"
                                        selected={active === "menu-orders-day-id"}
                                        onClick={handleOnClickMenu}
                                        to="/order-by-day"
                                        component={Link}>
                                        На день
                                    </ListItemButton>
                                </ListItem>
                                <ListItem sx={{mt: 0.5}}>
                                    <ListItemButton
                                        id="menu-orders-period-id"
                                        selected={active === "menu-orders-period-id"}
                                        onClick={handleOnClickMenu}
                                        to="/order-product-by-day"
                                        component={Link}>
                                        За период
                                    </ListItemButton>
                                </ListItem>
                            </List>
                        </Toggler>
                    </ListItem>

                    <ListItem nested>
                        <Toggler renderToggle={({open, setOpen}) => (
                            <ListItemButton onClick={() => setOpen(!open)}>
                                <StorefrontRoundedIcon/>
                                <ListItemContent>
                                    <Typography level="title-sm">Остатки</Typography>
                                </ListItemContent>
                                <KeyboardArrowDownIcon
                                    sx={{transform: open ? 'rotate(180deg)' : 'none'}}
                                />
                            </ListItemButton>
                        )}>
                            <List sx={{gap: 0.5}}>
                                <ListItem sx={{mt: 0.5}}>
                                    <ListItemButton
                                        id="menu-stocks-id"
                                        selected={active === "menu-stocks-id"}
                                        onClick={handleOnClickMenu}
                                        to="/stocks"
                                        component={Link}>
                                        На день
                                    </ListItemButton>
                                </ListItem>
                            </List>
                        </Toggler>
                    </ListItem>

                    <ListItem nested>
                        <Toggler
                            renderToggle={({open, setOpen}) => (
                                <ListItemButton onClick={() => setOpen(!open)}>
                                    <FolderRoundedIcon/>
                                    <ListItemContent>
                                        <Typography level="title-sm">Справочники</Typography>
                                    </ListItemContent>
                                    <KeyboardArrowDownIcon
                                        sx={{transform: open ? 'rotate(180deg)' : 'none'}}
                                    />
                                </ListItemButton>
                            )}
                        >
                            <List sx={{gap: 0.5}}>
                                <ListItem sx={{mt: 0.5}}>
                                    <ListItemButton
                                        id="menu-dict-item1c-id"
                                        selected={active === "menu-dict-item1c-id"}
                                        onClick={handleOnClickMenu}
                                        to="/dict-position"
                                        component={Link}>Позиций 1С</ListItemButton>
                                </ListItem>
                                <ListItem>
                                    <ListItemButton
                                        id="menu-dict-clusters-id"
                                        selected={active === "menu-dict-clusters-id"}
                                        onClick={handleOnClickMenu}
                                        to="/clusters"
                                        component={Link}>Кластеры</ListItemButton>
                                </ListItem>
                            </List>
                        </Toggler>
                    </ListItem>

                    {/*<ListItem>*/}
                    {/*    <ListItemButton*/}
                    {/*        role="menuitem"*/}
                    {/*        component="a"*/}
                    {/*        href="/joy-ui/getting-started/templates/messages/"*/}
                    {/*    >*/}
                    {/*        <QuestionAnswerRoundedIcon/>*/}
                    {/*        <ListItemContent>*/}
                    {/*            <Typography level="title-sm">Messages</Typography>*/}
                    {/*        </ListItemContent>*/}
                    {/*        <Chip size="sm" color="primary" variant="solid">*/}
                    {/*            4*/}
                    {/*        </Chip>*/}
                    {/*    </ListItemButton>*/}
                    {/*</ListItem>*/}

                    {/*<ListItem nested>*/}
                    {/*    <Toggler*/}
                    {/*        renderToggle={({open, setOpen}) => (*/}
                    {/*            <ListItemButton onClick={() => setOpen(!open)}>*/}
                    {/*                <GroupRoundedIcon/>*/}
                    {/*                <ListItemContent>*/}
                    {/*                    <Typography level="title-sm">Users</Typography>*/}
                    {/*                </ListItemContent>*/}
                    {/*                <KeyboardArrowDownIcon*/}
                    {/*                    sx={{transform: open ? 'rotate(180deg)' : 'none'}}*/}
                    {/*                />*/}
                    {/*            </ListItemButton>*/}
                    {/*        )}*/}
                    {/*    >*/}
                    {/*        <List sx={{gap: 0.5}}>*/}
                    {/*            <ListItem sx={{mt: 0.5}}>*/}
                    {/*                <ListItemButton*/}
                    {/*                    role="menuitem"*/}
                    {/*                    component="a"*/}
                    {/*                    href="/joy-ui/getting-started/templates/profile-dashboard/"*/}
                    {/*                >*/}
                    {/*                    My profile*/}
                    {/*                </ListItemButton>*/}
                    {/*            </ListItem>*/}
                    {/*            <ListItem>*/}
                    {/*                <ListItemButton>Create a new user</ListItemButton>*/}
                    {/*            </ListItem>*/}
                    {/*            <ListItem>*/}
                    {/*                <ListItemButton>Roles & permission</ListItemButton>*/}
                    {/*            </ListItem>*/}
                    {/*        </List>*/}
                    {/*    </Toggler>*/}
                    {/*</ListItem>*/}
                </List>

            </Box>
            <Divider/>
            <Box sx={{display: 'flex', gap: 1, alignItems: 'center'}}>
                <Box sx={{minWidth: 0, flex: 1}}>
                    <Typography level="title-sm">{profile?.name ?? 'Неизвесный пользователь'}</Typography>
                    <Typography level="body-xs">{profile?.email}</Typography>
                </Box>
                <IconButton size="sm" variant="plain" color="neutral">
                    <LogoutRoundedIcon/>
                </IconButton>
            </Box>
        </Sheet>
    );
}
