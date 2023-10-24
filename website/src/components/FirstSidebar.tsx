import * as React from 'react';
import { useEffect, useState } from 'react';
import GlobalStyles from '@mui/joy/GlobalStyles';
import Avatar from '@mui/joy/Avatar';
import Divider from '@mui/joy/Divider';
import List from '@mui/joy/List';
import ListItem from '@mui/joy/ListItem';
import ListItemButton from '@mui/joy/ListItemButton';
import Sheet from '@mui/joy/Sheet';
// icons
import SettingsRoundedIcon from '@mui/icons-material/SettingsRounded';
import SupportRoundedIcon from '@mui/icons-material/SupportRounded';
import { openSidebar } from '../utils';
import { useController } from '../context';
import { FirstSidebarActive } from '../context/actions';

export default function FirstSidebar() {
    const {state, dispatch} = useController()
    const {menu} = state
    const {activeIndex, data} = menu
    const [selectedIndex, setSelectedIndex] = useState(activeIndex)

    useEffect(() => {
    }, [activeIndex])
    const handleOnMenuClick = (index: number) => {
        if (activeIndex !== index) {
            dispatch(FirstSidebarActive(index))
        }
        openSidebar()
    }

    interface IProps {
        index: number
        children?: string | React.JSX.Element | React.JSX.Element[]
    }

    const MenuItem = ({index, children}: IProps): React.JSX.Element => {
        return (<ListItemButton selected={activeIndex === index} variant={activeIndex === index ? "soft" : undefined}
                                onClick={() => handleOnMenuClick(index)}>
            {children}
        </ListItemButton>)
    }

    return (
        <Sheet
            className="FirstSidebar"
            sx={{
                position: {
                    xs: 'fixed',
                    md: 'sticky',
                },
                transform: {
                    xs: 'translateX(calc(100% * (var(--SideNavigation-slideIn, 0) - 1)))',
                    md: 'none',
                },
                transition: 'transform 0.4s',
                zIndex: 1000,
                height: '100dvh',
                width: 'var(--FirstSidebar-width)',
                top: 0,
                p: 2,
                flexShrink: 0,
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                gap: 1,
                borderRight: '1px solid',
                borderColor: 'divider',
            }}
        >
            <GlobalStyles
                styles={{
                    ':root': {
                        '--FirstSidebar-width': '68px',
                    },
                }}
            />
            <List size="sm" sx={{'--ListItem-radius': '6px', '--List-gap': '8px'}}>
                {
                    data.map((item, index) => (
                        <ListItem key={`f_menu_${index}`}>
                            <MenuItem index={index}>
                                {item.icon}
                            </MenuItem>
                        </ListItem>
                    ))
                }
            </List>
            <List
                sx={{
                    mt: 'auto',
                    flexGrow: 0,
                    '--ListItem-radius': '8px',
                    '--List-gap': '4px',
                }}
            >
                <ListItem>
                    <ListItemButton>
                        <SupportRoundedIcon/>
                    </ListItemButton>
                </ListItem>
                <ListItem>
                    <ListItemButton>
                        <SettingsRoundedIcon/>
                    </ListItemButton>
                </ListItem>
            </List>
            <Divider/>
            <Avatar variant="outlined" size="sm" src="/static/images/avatar/3.jpg"/>
        </Sheet>
    );
}
