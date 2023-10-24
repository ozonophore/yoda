/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Stocks } from '../models/Stocks';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class StocksService {

    /**
     * Получение остатков товаров
     * Получение отстатоков за дату
     * @param date Дата (YYYY-MM-DD)
     * @returns Stocks OK
     * @throws ApiError
     */
    public static getStocks(
        date: string,
    ): CancelablePromise<Stocks> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/stocks/{date}',
            path: {
                'date': date,
            },
            errors: {
                404: `Not Found`,
            },
        });
    }

}
