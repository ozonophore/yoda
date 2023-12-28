import React, { Fragment, useEffect } from 'react';

import 'dayjs/locale/ru';
import { useController } from 'context';
import { SetMenuActive } from 'context/actions';

export default function Home() {
    const {dispatch} = useController()
    useEffect(() => {
        dispatch(SetMenuActive("menu-home-id"))
    }, []);
    return (
        <Fragment>

        </Fragment>
    )
}