/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { DictPositions } from '../models/DictPositions';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class DictionariesService {

    /**
     * Получение списка позиций
     * @returns DictPositions OK
     * @throws ApiError
     */
    public static getPositions(): CancelablePromise<DictPositions> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/dictionaries/positions',
        });
    }

}
