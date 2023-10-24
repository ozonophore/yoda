import * as React from 'react';
import { useEffect } from 'react';
import Box from '@mui/joy/Box';
import List from '@mui/joy/List';
import ListSubheader from '@mui/joy/ListSubheader';
import ListItem from '@mui/joy/ListItem';
import ListItemContent from '@mui/joy/ListItemContent';
import ListItemDecorator from '@mui/joy/ListItemDecorator';
import ListItemButton from '@mui/joy/ListItemButton';

import { closeSidebar } from '../utils';
import { useController } from '../context';
import { SecondSidebarActive } from '../context/actions';
import Sheet from '@mui/joy/Sheet';

export default function SecondSidebar() {
    const {state, dispatch} = useController()
    const {data, activeIndex, activeSubIndex} = state.menu
    const menu = data[state.menu.activeIndex].subMenu ?? []
    const isNotEmpty = menu.length !== 0

    useEffect(() => {
        console.log("# change")
    }, [activeIndex, activeSubIndex])

    const handleOnClick = (index: number) => {
        dispatch(SecondSidebarActive(index))
        closeSidebar()
    }

    return (
        <React.Fragment>
            <Box
                className="SecondSidebar-overlay"
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
                        xs: 'translateX(calc(100% * (var(--SideNavigation-slideIn, 0) - 1) + var(--SideNavigation-slideIn, 0) * var(--FirstSidebar-width, 0px)))',
                        lg: 'translateX(-100%)',
                    },
                }}
                onClick={() => closeSidebar()}
            />
            {isNotEmpty && (
            <Sheet
                className="SecondSidebar"
                color="neutral"
                sx={{
                    position: {
                        xs: 'fixed',
                        lg: 'sticky',
                    },
                    transform: {
                        xs: 'translateX(calc(100% * (var(--SideNavigation-slideIn, 0) - 1) + var(--SideNavigation-slideIn, 0) * var(--FirstSidebar-width, 0px)))',
                        lg: 'none',
                    },
                    transition: 'transform 0.4s',
                    zIndex: 9999,
                    height: '100dvh',
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
                <List
                    size="sm"
                    sx={{
                        '--ListItem-radius': '6px',
                        '--List-gap': '6px',
                    }}
                >
                    {
                        menu.map((item, index) => {
                                if (item.type === "menu") {
                                    return (<ListItem key={`s_menu_${index}`}>
                                        <ListItemButton onClick={() => handleOnClick(item.index ?? 0)}
                                                        selected={activeSubIndex === item.index}>
                                            <ListItemDecorator>
                                                {item.icon}
                                            </ListItemDecorator>
                                            <ListItemContent>{item.name}</ListItemContent>
                                        </ListItemButton>
                                    </ListItem>)
                                }
                                return (<ListSubheader key={`sh_menu_${index}`} role="presentation" sx={{fontWeight: 'lg'}}>
                                    {item.name}
                                </ListSubheader>)
                            }
                        )
                    }
                </List>
            </Sheet>)}
        </React.Fragment>
    );
}
