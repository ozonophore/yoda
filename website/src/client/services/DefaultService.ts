/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { OrderProducts } from '../models/OrderProducts';
import type { ProductParams } from '../models/ProductParams';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class DefaultService {

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

}
