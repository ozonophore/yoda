/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type StockFull = {
    stockDate: string;
    source: string;
    /**
     * Организация
     */
    organization: string;
    supplierArticle: string;
    barcode: string;
    sku: string;
    name: string;
    brand: string;
    warehouse: string;
    /**
     * Stock quantity
     */
    quantity: number;
    price: number;
    priceWithDiscount: number;
};

