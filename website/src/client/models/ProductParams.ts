/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type ProductParams = {
    dateFrom: string;
    dateTo: string;
    filter?: string;
    limit?: number;
    offset?: number;
    groupBy?: ProductParams.groupBy;
};

export namespace ProductParams {

    export enum groupBy {
        POSITION = 'POSITION',
    }


}

