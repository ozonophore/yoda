/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { DictPositions } from '../models/DictPositions';
import type { PageProductParams } from '../models/PageProductParams';
import type { Warehouse } from '../models/Warehouse';
import type { Warehouses } from '../models/Warehouses';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class DictionariesService {

    /**
     * Получение списка позиций
     * @param requestBody
     * @returns DictPositions OK
     * @throws ApiError
     */
    public static getPositions(
        requestBody: PageProductParams,
    ): CancelablePromise<DictPositions> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/dictionaries/positions',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * @param filter
     * @returns string OK
     * @throws ApiError
     */
    public static getClusters(
        filter?: string,
    ): CancelablePromise<Array<string>> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/dictionaries/clusters',
            query: {
                'filter': filter,
            },
        });
    }

    /**
     * Получение списка складов
     * @param source
     * @param cluster
     * @param code
     * @returns binary OK
     * @throws ApiError
     */
    public static exportWarehouses(
        source: Array<string>,
        cluster?: string,
        code?: string,
    ): CancelablePromise<Blob> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/dictionaries/warehouses/export',
            query: {
                'source': source,
                'cluster': cluster,
                'code': code,
            },
        });
    }

    /**
     * Получение списка складов
     * @param source
     * @param limit
     * @param offset
     * @param cluster
     * @param code
     * @returns Warehouses OK
     * @throws ApiError
     */
    public static getWarehouses(
        source?: Array<string>,
        limit?: number,
        offset?: number,
        cluster?: string,
        code?: string,
    ): CancelablePromise<Warehouses> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/dictionaries/warehouses',
            query: {
                'source': source,
                'limit': limit,
                'offset': offset,
                'cluster': cluster,
                'code': code,
            },
        });
    }

    /**
     * Добавление кластера к существующему складу
     * @param requestBody
     * @returns Warehouse OK
     * @throws ApiError
     */
    public static updateWarehouse(
        requestBody: Warehouse,
    ): CancelablePromise<Warehouse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/dictionaries/warehouses',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

}
