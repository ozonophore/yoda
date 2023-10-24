/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { AuthInfo } from '../models/AuthInfo';
import type { LoginInfo } from '../models/LoginInfo';
import type { Profile } from '../models/Profile';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class AuthService {

    /**
     * Authetification
     * @param requestBody
     * @returns AuthInfo OK
     * @throws ApiError
     */
    public static login(
        requestBody: LoginInfo,
    ): CancelablePromise<AuthInfo> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/auth/login',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * Refresh token
     * @returns AuthInfo OK
     * @throws ApiError
     */
    public static refresh(): CancelablePromise<AuthInfo> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/auth/refresh',
        });
    }

    /**
     * Get profile
     * @returns Profile OK
     * @throws ApiError
     */
    public static profile(): CancelablePromise<Profile> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/auth/profile',
        });
    }

}
