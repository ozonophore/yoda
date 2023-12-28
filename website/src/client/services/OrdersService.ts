/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { OrderProducts } from '../models/OrderProducts';
import type { Orders } from '../models/Orders';
import type { ProductParams } from '../models/ProductParams';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class OrdersService {

    /**
     * Выгрузка заказа за день
     * @param page
     * @param size
     * @param date
     * @param filter
     * @param source
     * @returns Orders Ok
     * @throws ApiError
     */
    public static getOrders(
        page: number,
        size: number,
        date?: string,
        filter?: string,
        source?: string,
    ): CancelablePromise<Orders> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/orders',
            query: {
                'date': date,
                'page': page,
                'size': size,
                'filter': filter,
                'source': source,
            },
            errors: {
                401: `Unauthorized`,
                404: `Not Found`,
            },
        });
    }

    /**
     * Заказы по продукту
     * @param requestBody
     * @returns OrderProducts Ok
     * @throws ApiError
     */
    public static getOrdersProduct(
        requestBody: ProductParams,
    ): CancelablePromise<OrderProducts> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/orders/product',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                401: `Unauthorized`,
                404: `Not Found`,
                418: `Error`,
            },
        });
    }

    /**
     * Выгрузка в excel заказов по продукту
     * @param requestBody
     * @returns binary Ok
     * @throws ApiError
     */
    public static exportOrdersProductToExcel(
        requestBody: ProductParams,
    ): CancelablePromise<Blob> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/orders/product/report',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                401: `Unauthorized`,
                404: `Not Found`,
                418: `Error`,
            },
        });
    }

    /**
     * Выгрузка заказа за день
     * @param date
     * @returns binary OK
     * @throws ApiError
     */
    public static getOrdersReport(
        date?: string,
    ): CancelablePromise<Blob> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/orders/report',
            query: {
                'date': date,
            },
        });
    }

}
