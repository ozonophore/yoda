import { OpenAPI } from 'client';
import { downloadFile } from 'utils';

export class CustomOrdersService {
    public static getOrdersReport(
        onStartDownload: () => void,
        onFinishDownload: () => void,
        date?: string,
    ) {

        const url = OpenAPI.BASE + `/orders/report${date ? `?date=${date}` : ''}`
        const authHeader = `Bearer ${OpenAPI.TOKEN}`

        const options = {
            method: 'GET',
            headers: {
                Authorization: authHeader
            }
        };

        downloadFile(url,
            `order_${date}.xlsx`,
            onStartDownload,
            onFinishDownload,
            options)
    }
}