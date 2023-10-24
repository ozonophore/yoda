/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Sales } from '../models/Sales';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class SalesesService {

    /**
     * Продажи за месяц
     * @param year
     * @param month
     * @returns binary OK
     * @throws ApiError
     */
    public static getSalesByMonthReport(
        year: number,
        month: number,
    ): CancelablePromise<Blob> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/sales/report',
            query: {
                'year': year,
                'month': month,
            },
        });
    }

    /**
     * Продажи за месяц
     * @param year
     * @param month
     * @param page
     * @param size
     * @returns Sales OK
     * @throws ApiError
     */
    public static getSalesByMonthWithPagination(
        year: number,
        month: number,
        page: number,
        size: number,
    ): CancelablePromise<Sales> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/sales',
            query: {
                'year': year,
                'month': month,
                'page': page,
                'size': size,
            },
            errors: {
                401: `Unauthorized`,
            },
        });
    }

}
