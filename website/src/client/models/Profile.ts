/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { Permission } from './Permission';

export type Profile = {
    /**
     * Email пользователя
     */
    email: string;
    /**
     * Имя пользователя
     */
    name: string;
    /**
     * Права пользователя
     */
    permissions: Array<Permission>;
};

