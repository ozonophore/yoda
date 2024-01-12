/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Stocks } from '../models/Stocks';
import type { StocksFull } from '../models/StocksFull';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class StocksService {

    /**
     * Получение остатков товаров
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

    /**
     * Получение остатков товаров
     * Выгрузка отчета отстатоков за дату
     * @param date Дата (YYYY-MM-DD)
     * @returns binary OK
     * @throws ApiError
     */
    public static exportStocks(
        date: string,
    ): CancelablePromise<Blob> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/stocks/export',
            query: {
                'date': date,
            },
        });
    }

    /**
     * Получение остатков товаров
     * Получение отстатоков за дату
     * @param date Дата (YYYY-MM-DD)
     * @param limit
     * @param offset
     * @returns StocksFull OK
     * @throws ApiError
     */
    public static getStocksWithPages(
        date: string,
        limit: number,
        offset: number,
    ): CancelablePromise<StocksFull> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/stocks',
            query: {
                'date': date,
                'limit': limit,
                'offset': offset,
            },
            errors: {
                404: `Not Found`,
            },
        });
    }

}
